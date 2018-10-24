package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (cc *BenchmarkChaincode) putPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) editPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) queryPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) queryAllPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) queryCouchPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) filterPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) filterCouchPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}