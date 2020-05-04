
 package commandcontrol

 import (
	 "encoding/json"
	 "fmt"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 pb "github.com/hyperledger/fabric/protos/peer"
 )

 type ComCtrlChainCode struct {}

 type Command struct {
	 Name string `json:"name"`
	 Id   string `json:"id"`
	 Data string `json:"data"`
 }
 type Info struct {
	 Name string `json:"name"`
	 Id   string `json:"id"`
	 Content string `json:"content"`
 }
 type WeaponState struct {
	 Name string `json:"name"`
	 Id   string `json:"id"`
	 State string `json:"state"`
 }
 func (t *ComCtrlChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	 return shim.Success(nil)
 }

 func (t *ComCtrlChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
 	//get the called function name and args
	 funcName, args := stub.GetFunctionAndParameters()
	 switch funcName {
	 case "setCommand":
		 return setCommand(stub, args)
	 case "getCommand":
		 return getCommand(stub, args)
	 case "deleteCommand":
	 	return deleteCommand(stub,args)
	 case "setInfo":
		 return setInfo(stub, args)
	 case "getInfo":
		 return getInfo(stub, args)
	 case "deleteInfo":
	 	return deleteInfo(stub, args)
	 case "setWeaponState":
		 return setWeaponState(stub, args)
	 case "getWeaponState":
		 return getWeaponState(stub, args)
	 case "deleteWeaponState":
		 return deleteWeaponState(stub,args)
	 default:
		 return shim.Error(fmt.Sprintf("unsupported function: %s", funcName))
	 }

 }
 //type Command struct {                 
	// Name string `json:"name"`         
	// Id   string `json:"id"`           
	// Data string `json:"data"`         
 //}
 // the CommandKey is used for data storage in blockchain
 func constructCommandKey(commandName string, commandId string) string {
	 return fmt.Sprintf("command_%s_%s", commandName, commandId)
 }
 /*
 the command can be set, but cannot be update
  */
 func setCommand(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 	//check the number of args
 	if len(args) != 3 {
		return shim.Error("not enough args")
	}

	//get the parameters
	name := args[0]
	id := args[1]
	data := args[2]

	//invalid args
	if name == "" || id == "" || data == "" {
		return shim.Error("invalid null args")
	}

	//check whether command exists
	var commandBytes []byte
	var err error
	commandBytes, err = stub.GetState(constructCommandKey(name,id))
	if err == nil && len(commandBytes) != 0 {
		return shim.Error("command already exists! ")
	}

	//put data into blockchain
	command := &Command{
		Name: name,
		Id:   id,
		Data: data,
	}
	commandBytes, err = json.Marshal(command)
	if err != nil {
		return shim.Error(fmt.Sprintf("json.Marshal(command) error: %s", err))
	}
	err = stub.PutState(constructCommandKey(name,id), commandBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("put command error: %s", err))
	}
	return shim.Success(nil)
 }


 func getCommand(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 	//check the number of args
	 if len(args) != 2 {
		 return shim.Error("not enough args for getCommand request")
	 }

	 //validation of args
	 name := args[0]
	 id := args[1]
	 if name == "" || id == "" {
		 return shim.Error("invalid null args")
	 }

	 //get data from blockchain
	 commandBytes, err := stub.GetState(constructCommandKey(name,id))
	 if err != nil || len(commandBytes) == 0 {
		 return shim.Error("command not found")
	 }
	 return shim.Success(commandBytes)
 }
 func deleteCommand(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 //check the number of args
	 if len(args) != 2 {
		 return shim.Error("not enough args for deleteCommand request")
	 }

	 //validation of args
	 name := args[0]
	 id := args[1]
	 if name == "" || id == "" {
		 return shim.Error("invalid null args")
	 }

	 //get data from blockchain
	 commandBytes, err := stub.GetState(constructCommandKey(name,id))
	 if err != nil || len(commandBytes) == 0 {
		 return shim.Error("command not found")
	 }
	 err = stub.DelState(constructCommandKey(name,id))
	 if err != nil {
	 	return shim.Error(fmt.Sprintf("delete command error: %s",err))
	 }
	 return shim.Success(nil)
 }


 //type Info struct {
	// Name string `json:"name"`
	// Id   string `json:"id"`
	// Content string `json:"content"`
 //}
 func constructInfoKey(infoName string, infoId string) string {
	 return fmt.Sprintf("info_%s_%s", infoName, infoId)
 }

 func setInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 //correct number of args
	 if len(args) != 3 {
		 return shim.Error("not enough args")
	 }

	 //get the parameters
	 name := args[0]
	 id := args[1]
	 content := args[2]

	 //invalid args
	 if name == "" || id == "" || content == "" {
		 return shim.Error("invalid null args for command name or id")
	 }

	 //check whether command exists
	 var infoBytes []byte
	 var err error
	 infoBytes, err = stub.GetState(constructInfoKey(name,id))
	 if err == nil && len(infoBytes) != 0 {
		 return shim.Error("info already exists!")
	 }

	 //put data into blockchain
	 info := &Info{
		 Name:    name,
		 Id:      id,
		 Content: content,
	 }
	 infoBytes, err = json.Marshal(info)
	 if err != nil {
		 return shim.Error(fmt.Sprintf("json.Marshal(info) error: %s", err))
	 }
	 err = stub.PutState(constructInfoKey(name,id), infoBytes)
	 if err != nil {
		 return shim.Error(fmt.Sprintf("put info error: %s", err))
	 }
	 return shim.Success(nil)
 }


 func getInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 //check the number of args
	 if len(args) != 2 {
		 return shim.Error("not enough args for getInfo request")
	 }

	 //validation of args
	 name := args[0]
	 id := args[1]
	 if name == "" || id == "" {
		 return shim.Error("invalid null args for info name or id")
	 }

	 //get data from blockchain
	 infoBytes, err := stub.GetState(constructInfoKey(name,id))
	 if err != nil || len(infoBytes) == 0 {
		 return shim.Error("info not found!")
	 }
	 return shim.Success(infoBytes)
 }
 func deleteInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 //check the number of args
	 if len(args) != 2 {
		 return shim.Error("not enough args for deleteInfo request")
	 }

	 //validation of args
	 name := args[0]
	 id := args[1]
	 if name == "" || id == "" {
		 return shim.Error("invalid null args for info name or id")
	 }

	 //get data from blockchain
	 infoBytes, err := stub.GetState(constructInfoKey(name,id))
	 if err != nil || len(infoBytes) == 0 {
		 return shim.Error("info not found!")
	 }
	 err = stub.DelState(constructInfoKey(name,id))
	 if err != nil {
	 	return shim.Error(fmt.Sprintf("delete info error: %s", err))
	 }
	 return shim.Success(nil)
 }
 //type WeaponState struct {
	// Name string `json:"name"`
	// Id   string `json:"id"`
	// State string `json:"state"`
 //}
 func constructWeaponStateKey(weaponName string, weaponId string) string {
	 return fmt.Sprintf("weaponState_%s_%s", weaponName, weaponId)
 }
 func setWeaponState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 //check the number of args
	 if len(args) != 3 {
		 return shim.Error("not enough args")
	 }

	 //get the parameters
	 name := args[0]
	 id := args[1]
	 state := args[2]

	 //invalid args
	 if name == "" || id == "" || state == "" {
		 return shim.Error("invalid null args")
	 }

	 //put data into blockchain
	 weaponState := &WeaponState{
		 Name:    name,
		 Id:      id,
		 State:   state,
	 }
	 weaponBytes, err := json.Marshal(weaponState)
	 if err != nil {
	 	return shim.Error(fmt.Sprintf("json.Marshal(weaponState) error: %s", err))
	 }
	 err = stub.PutState(constructWeaponStateKey(name,id), weaponBytes)
	 if err != nil {
		 return shim.Error(fmt.Sprintf("put weapon state error: %s", err))
	 }
	 return shim.Success(nil)
 }

 func getWeaponState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 //check the number of args
	 if len(args) != 2 {
		 return shim.Error("not enough args for getWeaponState request")
	 }

	 //validation of args
	 name := args[0]
	 id := args[1]
	 if name == "" || id == "" {
		 return shim.Error("empty variable for weapon name or id")
	 }

	 //get data from blockchain
	 weaponStateBytes, err := stub.GetState(constructWeaponStateKey(name,id))
	 if err != nil || len(weaponStateBytes) == 0 {
		 return shim.Error("weaponState not found")
	 }
	 return shim.Success(weaponStateBytes)
 }

 func deleteWeaponState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 //check the number of args
	 if len(args) != 2 {
		 return shim.Error("not enough args for deleteWeaponState request")
	 }

	 //validation of args
	 name := args[0]
	 id := args[1]
	 if name == "" || id == "" {
		 return shim.Error("empty variable for weapon name or id")
	 }

	 //get data from blockchain
	 weaponStateBytes, err := stub.GetState(constructWeaponStateKey(name,id))
	 if err != nil || len(weaponStateBytes) == 0 {
		 return shim.Error("weaponState not found")
	 }
	 err = stub.DelState(constructWeaponStateKey(name,id))
	 if err != nil {
	 	return shim.Error(fmt.Sprintf("delete weapon state error: %s",err))
	 }
	 return shim.Success(weaponStateBytes)
 }










