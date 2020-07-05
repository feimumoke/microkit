## golang学习工程


docker pull consul

docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0

-server 代表服务端方式启动
-bootstrap 指定自己为leader 不需要选举

-ui启动一个内置管理web界面
-client 指定客户端可以访问的IP，0.0.0.0代表任意访问

http://localhost:8500/ui/dc1/services
http://localhost:8500/v1/agent/services

curl --request PUT --data @p.json localhost:8500/v1/agent/service/register
curl --request PUT localhost:8500/v1/agent/service/deregister/{ID}

单项验证生成服务端证书
openssl genrsa -des3 -out server.key 2048
openssl req -new -key server.key -out server.csr
cp server.key server.key.org
openssl rsa -in server.key.org -out server.key
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

credentials.NewServerTLSFromFile(basepath+"server.crt", basepath+"server.key")
credentials.NewClientTLSFromFile(basepath+"server.crt", "tiger.com")

--------------------------------
Signature ok
subject=/C=cn/ST=beijing/L=beijing/O=tiger/OU=tiger/CN=tiger.com
Getting Private key
--------------------------------


双向认证
1、 生成ca证书
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 3650 -key ca.key -out ca.pem

Common Name: localhost

2、生成服务端证书
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem

--------------------------
Signature ok
subject=/C=cn/ST=beijing/L=beijing/O=tiger/OU=tiger/CN=localhost
Getting CA Private Key
---------------------------

3、生成客户端证书
openssl ecparam -genkey -name secp384r1 -out client.key
openssl req -new -key client.key -out client.csr
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem

-----------------------------
Signature ok
subject=/C=cn/ST=beijing/L=beijing/O=tiger/OU=tiger/CN=localhost
Getting CA Private Key
-----------------------------