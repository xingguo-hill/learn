#!/bin/bash
#根证书 
subj="/C=GB/L=BeiJing/O=Personal/CN=XG"
days=36500
pemdir=pem
confdir=conf
openssl genrsa -out $pemdir/ca.key 2048
openssl req -new -x509 -days $days -subj "$subj" -key $pemdir/ca.key -out $pemdir/ca.crt 
#服务端证书
openssl genrsa -out $pemdir/server.key 2048
#echo "openssl req -new -subj $subj -key $pemdir/server.key -out $pemdir/server.csr" 
openssl req -new -subj "$subj" -key $pemdir/server.key -out $pemdir/server.csr 
openssl x509 -req -sha256 -CA $pemdir/ca.crt -CAkey $pemdir/ca.key -CAcreateserial -days $days -in $pemdir/server.csr -out $pemdir/server.crt -extensions req_ext -extfile $confdir/san.conf 

#客户端证书
openssl genrsa -out $pemdir/client.key 2048
openssl req -new -subj "$subj" -key $pemdir/client.key -out $pemdir/client.csr 
openssl x509 -req -sha256 -CA $pemdir/ca.crt -CAkey $pemdir/ca.key -CAcreateserial -days $days -in $pemdir/client.csr -out $pemdir/client.crt -extensions req_ext -extfile $confdir/san.conf 
rm -f $pemdir/server.csr $pemdir/client.csr   