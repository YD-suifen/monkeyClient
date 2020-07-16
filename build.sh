

#!/bin/sh

cd ./
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build

tar -zcvf monkeyClient.tar.gz conf logs monkeyClient