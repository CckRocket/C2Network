/*
 * 文件名称 : chaincode.go 
 * 文件描述: 实现存储Helloworld字符串的智能合约
 */

 package helloworld

 import (
	 "fmt"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 pb "github.com/hyperledger/fabric/protos/peer"
 )
 
 // SimpleChaincode implements a simple chaincode to manage an asset
 type SimpleChaincode struct {
 }
 
 // 链码实例化时，调用Init函数初始化数据
 // 链码升级时，也会调用此函数重置或迁移数据
 func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	 return shim.Success(nil);
 }
 
 // 调用Invoke函数进行资产交易
 // 每笔交易通过get或set操作Init函数创建的key和value
 // 通过set可以创建新的key和value
 func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	 // 获取交易提案中的函数和参数
	 fn, args := stub.GetFunctionAndParameters()
 
	 var result string
	 var err error
	 switch fn {
	 case "set": 
	   result, err = set(stub, args)
	   if err != nil {
	   	return shim.Error(fmt.Sprintf("SetMsg Error: %s", err))
	   }

	 case "get":
		result, err = get(stub, args)
		if err != nil {
		 return shim.Error(fmt.Sprintf("GetMsg Error: %s", err))
	 	}
	 default:
		return shim.Error(fmt.Sprintf("unsupported function: %s", fn))
	 }
	 // Return the result as success payload
	 return shim.Success([]byte(result))
 }
 
 // 保存key和value到账本上
 // 如果key存在，覆盖原有的value
 func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	 if len(args) != 2 {
		 return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	 }
 
	 err := stub.PutState(args[0], []byte(args[1]))
	 if err != nil {
		 return "", fmt.Errorf("Failed to set asset: %s", args[0])
	 }
	 return args[1], nil
 }
 
 // 获取key对应的value
 func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	 if len(args) != 1 {
		 return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	 }
 
	 value, err := stub.GetState(args[0])
	 if err != nil {
		 return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	 }
	 if value == nil {
		 return "", fmt.Errorf("Asset not found: %s", args[0])
	 }
	 return string(value), nil
 }