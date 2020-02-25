#! /bin/bash -xe

cd executeCommand
chmod 777 executeCommand
./executeCommand &
cd ..

cd proxy
chmod 777 proxy
./proxy &
cd ..

cd scanner
chmod 777 scanner
./scanner &
cd ..