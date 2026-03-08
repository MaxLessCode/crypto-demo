FROM tailwindlabs/tailwindcss:latest AS tailwind
WORKDIR /app
COPY . .
RUN tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css

FROM golang:1.24-alpine AS build
WORKDIR /app

RUN apk add --no-cache build-base git
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=tailwind /app/assets/css/output.css ./assets/css/output.css

RUN templ generate
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./main.go

FROM alpine:3.20.2
WORKDIR /app
RUN apk add --no-cache ca-certificates
ENV GO_ENV=production

COPY --from=build /app/main .
COPY --from=build /app/assets ./assets

EXPOSE 8090
CMD ["./main"]