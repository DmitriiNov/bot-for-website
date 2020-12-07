FROM golang:1.15

RUN apt-get update && apt-get install -y sqlite

RUN mkdir /app

ADD . /app
WORKDIR /app

COPY . .

RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
