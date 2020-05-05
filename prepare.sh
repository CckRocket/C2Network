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

#commandchannel config
configtxgen -profile CommandChannel -outputAnchorPeersUpdate ./config/CommandChannel/OrgCCMSPanchors.tx -channelID commandchannel -asOrg OrgCCMSP

configtxgen -profile CommandChannel -outputAnchorPeersUpdate ./config/CommandChannel/OrgSC1MSPanchors.tx -channelID commandchannel -asOrg OrgSC1MSP

configtxgen -profile CommandChannel -outputAnchorPeersUpdate ./config/CommandChannel/OrgSC2MSPanchors.tx -channelID commandchannel -asOrg OrgSC2MSP

#sharedchannel config
configtxgen -profile SharedChannel -outputAnchorPeersUpdate ./config/SharedChannel/OrgCCMSPanchors.tx -channelID sharedchannel -asOrg OrgCCMSP

configtxgen -profile SharedChannel -outputAnchorPeersUpdate ./config/SharedChannel/OrgSC1MSPanchors.tx -channelID sharedchannel -asOrg OrgSC1MSP

configtxgen -profile SharedChannel -outputAnchorPeersUpdate ./config/SharedChannel/OrgSC2MSPanchors.tx -channelID sharedchannel -asOrg OrgSC2MSP

configtxgen -profile SharedChannel -outputAnchorPeersUpdate ./config/SharedChannel/OrgInfo1MSPanchors.tx -channelID sharedchannel -asOrg OrgInfo1MSP

configtxgen -profile SharedChannel -outputAnchorPeersUpdate ./config/SharedChannel/OrgInfo2MSPanchors.tx -channelID sharedchannel -asOrg OrgInfo2MSP
