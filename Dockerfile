FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY app .
EXPOSE 8000
CMD ["./app"]  
