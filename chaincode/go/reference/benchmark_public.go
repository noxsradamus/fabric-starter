package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (cc *BenchmarkChaincode) put(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) edit(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) queryAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) queryCouch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) filter(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) filterCouch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}