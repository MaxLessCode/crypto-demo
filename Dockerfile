FROM golang:1.24-alpine AS build
WORKDIR /app

RUN apk add --no-cache build-base git curl

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
    && chmod +x tailwindcss-linux-x64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ./tailwindcss-linux-x64 -i ./assets/css/input.css -o ./assets/css/output.css
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