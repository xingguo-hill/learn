#!/bin/bash
#根证书
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=gobook/CN=github.com" \
    -key ca.key -out ca.crt
#服务端证书
openssl genrsa -out server.key 2048
openssl req -new \
    -subj "/C=GB/L=China/O=server/CN=server.io" \
    -key server.key \
    -out server.csr
openssl x509 -req -sha256 \
    -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
    -in server.csr \
    -out server.crt

#客户端证书
openssl genrsa -out client.key 2048
openssl req -new \
    -subj "/C=GB/L=China/O=client/CN=client.io" \
    -key client.key \
    -out client.csr
openssl x509 -req -sha256 \
    -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
    -in client.csr \
    -out client.crt
rm -f server.csr client.csr