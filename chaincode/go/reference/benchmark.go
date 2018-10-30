
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("BenchmarkChaincode")

// BenchmarkChaincode example simple Chaincode implementation
type BenchmarkChaincode struct {
}

const (
	collectionName = "collection"
)

func (cc *BenchmarkChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debug("Init")

	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debug("Invoke")

	function, args := stub.GetFunctionAndParameters()
	if function == "put" {
		return cc.put(stub, args, "")
	} else if function == "edit" {
		return cc.edit(stub, args, "")
	} else if function == "query" {
		return cc.query(stub, args, "")
	} else if function == "queryAll" {
		return cc.queryAll(stub, args, "")
	} else if function == "queryCouch" {
		return cc.queryCouch(stub, args, "")
	} else if function == "filter" {
		return cc.filter(stub, args, "")
	} else if function == "filterCouch" {
		return cc.filterCouch(stub, args, "")
	} else if function == "putPrivate" {
		return cc.put(stub, args, collectionName)
	} else if function == "editPrivate" {
		return cc.edit(stub, args, collectionName)
	} else if function == "queryPrivate" {
		return cc.query(stub, args, collectionName)
	} else if function == "queryAllPrivate" {
		return cc.queryAll(stub, args, collectionName)
	} else if function == "queryCouchPrivate" {
		return cc.queryCouch(stub, args, collectionName)
	} else if function == "filterPrivate" {
		return cc.filter(stub, args, collectionName)
	} else if function == "filterCouchPrivate" {
		return cc.filterCouch(stub, args, collectionName)
	}

	return pb.Response{Status:403, Message:"Invalid invoke function name."}
}

func main() {
	err := shim.Start(new(BenchmarkChaincode))
	if err != nil {
		logger.Error(err.Error())
	}
}
