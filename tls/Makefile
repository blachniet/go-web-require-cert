CA_FILES := ca-key.pem ca.csr ca.pem
SERVER_FILES := server-key.pem server.csr server.pem
ALL_FILES := $(CA_FILES) $(SERVER_FILES) user1.pin user1.key user1.csr user1.pem user1.p12

all: $(ALL_FILES)

$(CA_FILES): ca-csr.json
	cfssl gencert -initca ca-csr.json | cfssljson -bare ca

$(SERVER_FILES): server-csr.json $(CA_FILES)
	cfssl gencert -ca ca.pem -ca-key ca-key.pem server-csr.json | cfssljson -bare server

user1.pin:
	openssl rand 40 -base64 > user1.pin

user1.key: user1.pin
	openssl genrsa -aes256 -passout file:user1.pin -out user1.key 4096

user1.csr: user1.key user1.pin
	openssl req -new -key user1.key -passin file:user1.pin -out user1.csr -subj '/CN=user1' -sha256

user1.pem: user1.csr ca.pem ca-key.pem
	openssl x509 -req -days 3650 -in user1.csr -CA ca.pem -CAkey ca-key.pem  -set_serial 1 -out user1.pem -sha256

user1.p12: user1.key user1.pin user1.pem
	openssl pkcs12 -export -out user1.p12 -inkey user1.key -passin file:user1.pin -in user1.pem -passout pass:password1234

clean:
	rm -f $(ALL_FILES)
