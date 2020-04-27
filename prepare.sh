#!/bin/bash
#generate crypto & certificate file
cryptogen generate --config=./crypto-config.yaml
#generate system config genesis block
configtxgen -profile OrdererGenesis -outputBlock ./config/genesis.block
# config InfoChannel with infochannel as id
configtxgen -profile InfoChannel -outputCreateChannelTx ./config/InfoChannel/infochannel.tx -channelID infochannel

# infochannel config
configtxgen -profile InfoChannel -outputAnchorPeersUpdate ./config/InfoChannel/OrgCCMSPanchors.tx -channelID infochannel -asOrg OrgCCMSP

configtxgen -profile InfoChannel -outputAnchorPeersUpdate ./config/InfoChannel/OrgInfo1MSPanchors.tx -channelID infochannel -asOrg OrgInfo1MSP

configtxgen -profile InfoChannel -outputAnchorPeersUpdate ./config/InfoChannel/OrgInfo2MSPanchors.tx -channelID infochannel -asOrg OrgInfo2MSP
