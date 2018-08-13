FROM golang:latest
WORKDIR /go/src/github.com/tensorchen/Meizhi/
COPY server/*.go .
RUN go build -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/tensorchen/Meizhi/app .
EXPOSE 8000
CMD ["./app"]  
