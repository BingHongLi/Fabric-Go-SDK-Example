config.yaml之內容，暫不公開

## 環境建置篇

```sh

mkdir hyperledger-go-experiment
cd hyperledger-go-experiment

git clone https://github.com/BingHongLi/fabric-samples.git

# 設定檔內容請看下方，需更改處會用粗體標示
vim config.yaml

cd fabric-samples/basic-network
sh start.sh
docker-compose up -d cli

# 引入新合約
cd ../chaincode
git clone https://github.com/BingHongLi/chaincode_basic_tutorial_lbh.git

# 連入Container中，安裝激活調用合約

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n chaincode_basic_tutorial_lbh -v 1.0 -p github.com/chaincode_basic_tutorial_lbh

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n chaincode_basic_tutorial_lbh -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["setSampleAsset","abc","123"]}' -C mychannel

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getSampleAsset","abc"]}' -C mychannel

```

## 安裝套件

```sh
go get -u github.com/hyperledger/fabric-sdk-go
go get -u github.com/cloudflare/cfssl
go get -u github.com/golang/mock
go get -u github.com/golang/protobuf
go get -u github.com/mitchellh/mapstructure
go get -u github.com/pkg/errors
go get -u github.com/spf13/cast
go get -u github.com/spf13/viper
go get -u github.com/stretchr/testify
go get -u golang.org/x/crypto/ocsp
go get -u golang.org/x/crypto/sha3
go get -u golang.org/x/net/context
go get -u google.golang.org/grpc
```
.golang.org/grpc
