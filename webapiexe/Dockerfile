FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app/webapi
RUN go build -o webapi .
CMD ["/app/webapi/webapi"]