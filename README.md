```sh
pushd tls
make
popd

security import tls/user1.p12

go run main.go

curl https://localhost:8081/ --cacert tls/ca.pem \
    --key tls/user1.key --cert tls/user1.pem --pass $(cat tls/user1.pin)
```
