package main

import (

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

func main() {

	// Step1 讀取設定檔，以創建總客戶端
	clientSDK,err1 :=fabsdk.New(config.FromFile("config.yaml"))

	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(clientSDK)

	// Step2 使用總客戶端創建系統資源管理客戶端
	resourceManagerContext := clientSDK.Context(fabsdk.WithUser("Admin"),fabsdk.WithOrg("Org1"))

	resourceClient ,errRe := resmgmt.New(resourceManagerContext)
	if errRe != nil{
		fmt.Println(errRe)
	}
	fmt.Println(resourceClient)


	// Step3 使用總客戶端創造MSP客戶端
	// caClient
	mspClient,err2 := msp.New(clientSDK.Context(),msp.WithOrg("Org1"))
	if err2 !=nil {
		fmt.Println(err2)
	}
	println("1")
	println(mspClient)


	// 取得Ca 管理員
	adminIdentity ,err3 :=mspClient.GetSigningIdentity("admin")
	if err3 != nil {
		fmt.Println(err3)
	}
	println("2")
	println(adminIdentity)


	// Step4 使用總客戶端創建channel客戶端

	channelProvider := clientSDK.ChannelContext("mychannel",
		fabsdk.WithUser("Admin"),
		fabsdk.WithOrg("Org1"))
	channelClient, _ := channel.New(channelProvider)


	// Step5 使用channel客戶端，調用合約，取得資料
	var args [][]byte
	args = append(args, []byte("abc"))

	request := channel.Request{
		ChaincodeID: "chaincode_basic_tutorial_lbh",
		Fcn:         "getSampleAsset",
		Args:        args,
	}

	response, err5 := channelClient.Query(request,)
	if err5 != nil {
		println(err5)
	}
	println(string(response.Payload))

}

