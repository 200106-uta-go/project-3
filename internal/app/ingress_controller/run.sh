#! /bin/bash -xe

cd proxy
chmod 777 proxy
./proxy
cd ..

cd scanner
chmod 777 scanner
./scanner &
cd ..