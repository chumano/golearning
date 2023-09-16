# golearning

## settings

Init module
`go mod init golearning`
``

Install modules in project
`go mod tidy`

Run command
`go run cmd/main.go`

Fix error : gcc: error: unrecognized command-line option '-m64' on MacOs M2
```
sudo dpkg --add-architecture amd64
sudo apt-get update
sudo apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross
export CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc
export CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=x86_64-linux-gnu-gcc

sudo apt install -y musl-tools musl-dev
sudo apt-get install -y build-essential
yes | sudo apt install gcc-x86-64-linux-gnu

sudo chown -R $(whoami):root ~/go
```
## project structure
/web
/api
/configs
/build
/scripts
/deployments
/docs
/test
