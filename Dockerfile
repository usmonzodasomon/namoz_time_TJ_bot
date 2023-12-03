FROM alpine

WORKDIR /build

COPY main .

CMD [". /main"]