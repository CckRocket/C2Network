#!/bin/bash
# remove files in config
cd config/
rm -rf genesis.block

cd CommandChannel
rm -rf *
cd ..

cd InfoChannel
rm -rf *
cd ..

cd SharedChannel
rm -rf *
cd ..

cd ..
# remove file in crypto-config
cd crypto-config
rm -rf *.block
rm -rf orderer*
rm -rf peer*
cd ..

