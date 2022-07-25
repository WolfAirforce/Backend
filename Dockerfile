FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/bin/api ./cmd/api

RUN chmod +x /app/bin/api

EXPOSE 8080

CMD [ "/app/bin/api" ]
