# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY . /app/

RUN echo "Hello World"
RUN ls

RUN go mod download

RUN GOOS=linux go build -o reminders-api cmd/main/main.go

EXPOSE 8452

CMD ["./reminders-api"]

