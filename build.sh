#/bin/bash

go mod tidy

go-bindata -pkg bindata -ignore bindata.go -o resources/bindata/bindata.go  resources/...
go build -o build/bin/envconfig  cmd/envconfig/envconfig.go
