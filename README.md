# Require and Verify Client Certificates in a Go Web App

This project demonstrates a Go web application that requires clients to present
a certificate in order to access it. The client certificate must be signed by a
predefined certificate authority (CA).

## Usage

Create the certificate authority, server certificate and client certificate.

```sh
pushd tls
make
popd
```

Start the Go server.

```sh
go run main.go
```

Make a request with `curl`.

```sh
curl https://localhost:8081/ --cacert tls/ca.pem \
    --key tls/user1.key --cert tls/user1.pem
```

The output should look like this:

```plain
Hello, user1!
```

In order to access the web app with your browser, you need to install the
certificate into either your browser or system's key store. With Chrome on
macOS, run the following.

```sh
security import tls/user1.p12
open https://localhost:8081/
```

