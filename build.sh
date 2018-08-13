#!/bin/sh
echo Building tensorchen/meizhi:1.0.0

docker build -t tensorchen/meizhi:latest . -f Dockerfile.build

docker create --name meizhi_1 tensorchen/meizhi:latest
docker cp meizhi_1:/go/src/github.com/tensorchen/meizhi/app ./app  
docker rm -f meizhi_1

echo Building tensorchen/meizhi:1.0.0

docker build --no-cache -t tensorchen/meizhi:latest .
rm ./app
