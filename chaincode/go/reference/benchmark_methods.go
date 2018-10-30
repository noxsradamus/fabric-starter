package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"encoding/json"
)

func (cc *BenchmarkChaincode) put(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {
	logger.Info("Benchmark put is running")
	logger.Debug("Benchmark.put")

	data := TestData{}
	if err := data.FillFromArguments(args); err != nil {
		message := fmt.Sprintf("cannot fill a data from arguments: %s", err.Error())
		logger.Error(message)
		return shim.Error(message)
	}

	if bytes, err := json.Marshal(data); err == nil {
		logger.Debug("Data: " + string(bytes))
	}

	if err := UpdateOrInsertIn(stub, &data, collection); err != nil {
		message := fmt.Sprintf("persistence error: %s", err.Error())
		logger.Error(message)
		return pb.Response{Status: 500, Message: message}
	}

	logger.Info("Benchmark.put exited without errors")
	logger.Debug("Success: Benchmark.put")
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) edit(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {
	logger.Info("Benchmark.edit is running")
	logger.Debug("Benchmark.edit")

	var data, dataToUpdate TestData
	if err := data.FillFromCompositeKeyParts(args[:1]); err != nil {
		message := fmt.Sprintf("cannot fill a data from arguments: %s", err.Error())
		logger.Error(message)
		return shim.Error(message)
	}

	dataToUpdate.Key = data.Key
	if err := data.FillFromArguments(args); err != nil {
		message := fmt.Sprintf("cannot fill a data from arguments: %s", err.Error())
		logger.Error(message)
		return shim.Error(message)
	}

	logger.Debug("Data to edit: " + data.Key)

	if !ExistsIn(stub, &dataToUpdate, collection) {
		message := fmt.Sprintf("data with ID %s not found", data.Key)
		logger.Error(message)
		return pb.Response{Status: 404, Message: message}
	}

	if err := LoadFrom(stub, &dataToUpdate, collection); err != nil {
		message := fmt.Sprintf("cannot load existing data: %s", err.Error())
		logger.Error(message)
		return pb.Response{Status: 404, Message: message}
	}

	dataToUpdate.Value = data.Value

	if bytes, err := json.Marshal(dataToUpdate); err == nil {
		logger.Debug("Data: " + string(bytes))
	}

	if err := UpdateOrInsertIn(stub, &dataToUpdate, collection); err != nil {
		message := fmt.Sprintf("persistence error: %s", err.Error())
		logger.Error(message)
		return pb.Response{Status: 500, Message: message}
	}

	logger.Info("Benchmark.edit exited without errors")
	logger.Debug("Success: Benchmark.edit")
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) query(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {
	logger.Info("Benchmark.query is running")
	logger.Debug("Benchmark.query")

	var data TestData
	if err := data.FillFromCompositeKeyParts(args[:1]); err != nil {
		message := fmt.Sprintf("cannot fill a data from arguments: %s", err.Error())
		logger.Error(message)
		return shim.Error(message)
	}
	logger.Debug("Data to query: " + data.Key)

	if !ExistsIn(stub, &data, collection) {
		message := fmt.Sprintf("data with ID %s not found", data.Key)
		logger.Error(message)
		return pb.Response{Status: 404, Message: message}
	}

	if err := LoadFrom(stub, &data, collection); err != nil {
		message := fmt.Sprintf("cannot load existing data: %s", err.Error())
		logger.Error(message)
		return pb.Response{Status: 404, Message: message}
	}

	result, err := json.Marshal(data)
	if err != nil {
		message := fmt.Sprintf("cannot marshaling a data: %s", err.Error())
		logger.Error(message)
		return shim.Error(message)
	}

	ledgerDataLogger.Debug("Result: " + string(result))

	logger.Info("Benchmark.query exited without errors")
	logger.Debug("Success: Benchmark.query")
	return shim.Success(result)
}

func (cc *BenchmarkChaincode) queryAll(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {
	logger.Info("Benchmark.queryAll is running")
	logger.Debug("Benchmark.queryAll")

	result, err := Query(stub, testDataIndex, []string{}, CreateTestData, EmptyFilter, collection)
	if err != nil {
		message := fmt.Sprintf("unable to perform query: %s", err.Error())
		logger.Error(message)
		return shim.Error(message)
	}

	logger.Debug("Result: " + string(result))

	logger.Info("Benchmark.queryAll exited without errors")
	logger.Debug("Success: Benchmark.queryAll")
	return shim.Success(result)
}

func (cc *BenchmarkChaincode) queryCouch(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {

	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) filter(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {
	return shim.Success(nil)
}

func (cc *BenchmarkChaincode) filterCouch(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {
	return shim.Success(nil)
}