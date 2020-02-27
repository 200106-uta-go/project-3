#! /bin/bash -xe

cd scanner
go build .
cd ..

cd proxy
go build .
cd ..