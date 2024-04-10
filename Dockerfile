FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build KeDuBack

EXPOSE 8080

CMD ["./KeDuBack"]
