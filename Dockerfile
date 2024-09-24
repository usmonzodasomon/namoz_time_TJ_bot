FROM golang:1.22.2 AS builder
WORKDIR /home/namazbot
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.19.1
WORKDIR /home/namazbot
COPY --from=builder /home/namazbot .
ENV TZ="Asia/Dushanbe"
RUN chmod +x ./main
CMD ["./main"]