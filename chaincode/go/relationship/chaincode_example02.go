package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("SimpleChaincode")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debug("Init")

	_, args := stub.GetFunctionAndParameters()
	var a, b string    // Entities
	var aVal, bVal int // Asset holdings
	var err error

	if len(args) != 4 {
		return pb.Response{Status:403, Message:"Incorrect number of arguments. Expecting 4"}
	}

	// Initialize the chaincode
	a = args[0]
	aVal, err = strconv.Atoi(args[1])
	if err != nil {
		return pb.Response{Status:403, Message:"Expecting integer value for asset holding"}
	}
	b = args[2]
	bVal, err = strconv.Atoi(args[3])
	if err != nil {
		return pb.Response{Status:403, Message:"Expecting integer value for asset holding"}
	}
	logger.Debugf("aVal, bVal = %d", aVal, bVal)

	// Write the state to the ledger
	err = stub.PutState(a, []byte(strconv.Itoa(aVal)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(b, []byte(strconv.Itoa(bVal)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debug("Invoke")

	function, args := stub.GetFunctionAndParameters()
	if function == "put" {
		return t.put(stub, args)
	} else if function == "get" {
		return t.get(stub, args)
	} else if function == "invokeChaincode" {
		return t.invokeChaincode(stub, args)
	}

	return pb.Response{Status:403, Message:"Invalid invoke function name."}
}

// put (key, value) pair in the ledger
func (t *SimpleChaincode) put(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	key := args[0]
	value := args[1]

	if err := stub.PutState(key, []byte(value)); err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// get value by key from the ledger
func (t *SimpleChaincode) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	key := args[0]

	// Get the state from the ledger
	valueBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}

	if valueBytes == nil {
		return pb.Response{Status:404, Message:"Entity not found"}
	}

	return shim.Success(valueBytes)
}

func toByteArray(args []string) [][]byte {
	res := [][]byte{}
	for _, arg := range args {
		res = append(res, []byte(arg))
	}

	return res
}

func (t *SimpleChaincode) invokeChaincode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	channelName := args[0]
	chaincodeName := args[1]
	chaincodeArgs := args[2:]

	response := stub.InvokeChaincode(chaincodeName, toByteArray(chaincodeArgs), channelName)
	logger.Debug(fmt.Sprintf("Response: status %d, message \"%s\", payload {%s}",
		response.Status, response.Message, string(response.Payload)))

	return response
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Error(err.Error())
	}
}
