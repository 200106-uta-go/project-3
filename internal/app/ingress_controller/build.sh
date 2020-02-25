#! /bin/bash -xe

cd sdn_Command
go build .
cd ..

cd sdn_Reasource
go build .
cd ..

cd sdn_Proxy
go build .
cd ..

cd ..

cd dashboard
go build .