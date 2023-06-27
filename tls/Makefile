CA_FILES := ca-key.pem ca.csr ca.pem
SERVER_FILES := server-key.pem server.csr server.pem
ALL_FILES := $(CA_FILES) $(SERVER_FILES) user1.key user1.csr user1.pem user1.p12

.PHONY: all
all: $(ALL_FILES)

$(CA_FILES): ca-csr.json
	cfssl gencert -initca ca-csr.json | cfssljson -bare ca

$(SERVER_FILES): server-csr.json $(CA_FILES)
	cfssl gencert -ca ca.pem -ca-key ca-key.pem server-csr.json | cfssljson -bare server

user1.key:
	openssl genrsa -out $@ 2048

user1.csr: user1.key user1.cnf
	openssl req -new -key user1.key -out $@ -sha256 -config user1.cnf

user1.pem: user1.csr ca.pem ca-key.pem user1.cnf
	openssl x509 -req -days 3650 -in user1.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out $@ -sha256 -copy_extensions copyall

user1.p12: user1.key user1.pem
	openssl pkcs12 -export -out $@ -inkey user1.key -in user1.pem -passout pass:password1234

.PHONY: clean
clean:
	rm -f $(ALL_FILES)
