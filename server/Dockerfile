FROM golang

MAINTAINER "tensorchen@gmail.com"

WORKDIR /Meizhi/

COPY . /Meizhi/

RUN go build  -o Meizhi *.go

EXPOSE 8000

ENTRYPOINT  ./Meizhi &
