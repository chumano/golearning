#!/usr/bin/env bash
set -euo pipefail

# install apt-add-repository
apt update 
apt install software-properties-common -y
apt update


# install go
GO_VERSION=1.21.1
curl -OL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
tar -C /home/vscode -xvf go${GO_VERSION}.linux-amd64.tar.gz

#source ~/.profile

