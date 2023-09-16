#!/usr/bin/env bash
echo "install-dev-tools.sh"

#sudo chown -R $(whoami):root ~/go

#echo "export PATH=\$PATH:/home/vscode/go/bin" >> ~/.profile
#echo "export GOPATH=/home/vscode/go" >> ~/.profile

#source ~/.profile

echo $PATH
echo $(whoami)
go version

curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

#install sql client
#go install githugo install github.com/cosmtrek/air@latest.com/kyleconroy/sqlc/cmd/sqlc@latest

go mod tidy