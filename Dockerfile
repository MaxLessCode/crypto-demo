# Build-Stage
FROM golang:1.24-alpine AS build
WORKDIR /app

# Install build dependencies (gcc, musl-dev, make pour CGO et toolchains)
RUN apk add --no-cache build-base

# Cache des modules Go : copier d'abord les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Installer templ pour la génération des templates
RUN go install github.com/a-h/templ/cmd/templ@latest

# Installer Node pour Tailwind
RUN apk add --no-cache nodejs npm

# Copier le code source (node_modules exclu via .dockerignore)
COPY . .

# Installer les dépendances Node et générer output.css
RUN npm ci 2>/dev/null || npm install

# Générer output.css et les templates
RUN npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css
RUN templ generate

# Build de l'application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Deploy-Stage
FROM alpine:3.20.2
WORKDIR /app

RUN apk add --no-cache ca-certificates

ENV GO_ENV=production

# Binaire et assets dans l'image finale
COPY --from=build /app/main .
COPY --from=build /app/assets ./assets

EXPOSE 8090

CMD ["./main"]
