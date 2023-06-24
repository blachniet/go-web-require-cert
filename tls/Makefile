PIN_LEN := 40
SERVER_FILES := server-key.pem server.csr server.pem
ALL_FILES := ca.pin ca.key ca-key.pem ca.pem user1.pin user1.key user1.csr user1.pem user1.p12 $(SERVER_FILES)

all: $(ALL_FILES)

ca.pin:
	openssl rand 40 -base64 > ca.pin

ca.key: ca.pin
	openssl genrsa -aes256 -passout file:ca.pin -out ca.key 4096

ca-key.pem: ca.key ca.pin
	openssl rsa -passin file:ca.pin -in ca.key -out ca-key.pem

ca.pem: ca.key ca.pin
	openssl req -new -x509 -days 3650 -key ca.key -passin file:ca.pin -out ca.pem -subj '/CN=example.com' -sha256

user1.pin:
	openssl rand 40 -base64 > user1.pin

user1.key: user1.pin
	openssl genrsa -aes256 -passout file:user1.pin -out user1.key 4096

user1.csr: user1.key user1.pin
	openssl req -new -key user1.key -passin file:user1.pin -out user1.csr -subj '/CN=user1' -sha256

user1.pem: user1.csr ca.pem ca.key ca.pin
	openssl x509 -req -days 3650 -in user1.csr -CA ca.pem -CAkey ca.key -passin file:ca.pin -set_serial 1 -out user1.pem -sha256

user1.p12: user1.key user1.pin user1.pem
	openssl pkcs12 -export -out user1.p12 -inkey user1.key -passin file:user1.pin -in user1.pem -passout pass:password1234

$(SERVER_FILES): server-csr.json ca.pem ca-key.pem
	cfssl gencert -ca ca.pem -ca-key ca-key.pem server-csr.json | cfssljson -bare server

clean:
	rm -f $(ALL_FILES)
