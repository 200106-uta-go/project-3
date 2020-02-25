#! /bin/bash -xe

cd executeCommand
go build .
cd ..

cd scanner
go build .
cd ..

cd proxy
go build .
cd ..