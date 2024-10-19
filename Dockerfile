FROM golang:alpine

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/kippy cmd/main.go

FROM gcr.io/distroless/base-debian10

COPY --from=0 /app/bin/kippy /

ENTRYPOINT [ "/kippy" ]