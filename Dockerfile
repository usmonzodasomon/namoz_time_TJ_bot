FROM golang:latest

WORKDIR /home/namazbot

COPY . .

ENV TZ="Asia/Dushanbe"

RUN go build -o main .

CMD ["/home/namazbot/main"]