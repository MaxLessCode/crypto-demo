FROM golang:1.24 AS build
WORKDIR /app

RUN apt-get update && apt-get install -y nodejs npm curl git build-essential

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN npx --yes @tailwindcss/cli@latest -i ./assets/css/input.css -o ./assets/css/output.css

RUN /go/bin/templ generate

RUN go build -o main ./main.go

FROM ubuntu:24.04
WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/main .
COPY --from=build /app/assets ./assets
COPY --from=build /app/.env* . 

EXPOSE 8090

CMD ["./main"]