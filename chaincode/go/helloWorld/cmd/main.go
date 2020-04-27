
 package main

 import (
	 "fmt"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 "github.com/hyperledger/fabric/ccnetwork/chaincode/go/helloWorld"
 )
 
 func main() {
	 err := shim.Start(new(helloworld.SimpleChaincode))
	 if err != nil {
		 fmt.Printf("Error starting Simple chaincode: %s", err)
	 }
 }
