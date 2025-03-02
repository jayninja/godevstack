#!/bin/bash
APP="goprom"

cd source
rm go.sum go.mod
go mod init ${APP}
go mod tidy

cd ..
docker build -t ${APP} .
