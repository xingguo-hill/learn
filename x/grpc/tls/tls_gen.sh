#!/bin/bash
#根证书 
subj="/C=GB/L=BeiJing/O=Personal/CN=XG"
days=36500
openssl genrsa -out pem/ca.key 2048
openssl req -new -x509 -days $days -subj "$subj" -key pem/ca.key -out pem/ca.crt 
#服务端证书
openssl genrsa -out pem/server.key 2048
openssl req -new -subj "$subj" -key pem/server.key -out pem/server.csr 
openssl x509 -req -sha256 -CA pem/ca.crt -CAkey pem/ca.key -CAcreateserial -days $days -in pem/server.csr -out pem/server.crt -extensions req_ext -extfile conf/sans.conf 

#客户端证书
openssl genrsa -out pem/client.key 2048
openssl req -new -subj "$subj" -key pem/client.key -out pem/client.csr 
openssl x509 -req -sha256 -CA pem/ca.crt -CAkey pem/ca.key -CAcreateserial -days $days -in pem/client.csr -out pem/client.crt -extensions req_ext -extfile conf/san.conf 
rm -f pem/server.csr pem/client.csr