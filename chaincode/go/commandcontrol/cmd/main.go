
 package main

 import (
	 "fmt"
	 "github.com/hyperledger/fabric/ccnetwork/chaincode/go/commandcontrol"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
 )
 
 func main() {
	 err := shim.Start(new(commandcontrol.ComCtrlChainCode))
	 if err != nil {
		 fmt.Printf("Error starting CommandControlCC: %s", err)
	 }
 }
