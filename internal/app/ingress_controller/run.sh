#! /bin/bash -xe

cd sdn_Command
chmod 777 sdn_Command
./sdn_Command &
cd ..

cd sdn_Proxy
chmod 777 sdn_Proxy
./sdn_Proxy &
cd ..

cd sdn_Reasource
chmod 777 sdn_Reasource
./sdn_Reasource &
cd ..

cd ..

cd dashboard
chmod 777 dashboard
./dashboard &