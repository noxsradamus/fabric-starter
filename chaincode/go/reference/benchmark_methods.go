package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strings"
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

	result, err := Query(stub, testDataIndex, []string{}, CreateTestData, EmptyFilter, []string{collection})
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
	logger.Info("Benchmark.queryCouch is running")
	logger.Debug("Benchmark.queryCouch")

	queryString := args[0]
	result, err := getQueryResultForQueryString(stub, queryString, collection)
	if err != nil {
		return shim.Error(err.Error())
	}

	ledgerDataLogger.Debug("Result: " + string(result))

	logger.Info("Benchmark.queryCouch exited without errors")
	logger.Debug("Success: Benchmark.queryCouch")
	return shim.Success(result)
}

func (cc *BenchmarkChaincode) filter(stub shim.ChaincodeStubInterface, args []string, collection string) pb.Response {
	logger.Info("Benchmark.filter is running")
	logger.Debug("Benchmark.filter")

	valueFilter := func(data LedgerData) bool {
		testData, ok := data.(*TestData)
		if ok && strings.HasPrefix(testData.Value, "filter") {
			return true
		}

		return false
	}

	result, err := Query(stub, testDataIndex, []string{}, CreateTestData, valueFilter, []string{collection})
	if err != nil {
		message := fmt.Sprintf("unable to perform query: %s", err.Error())
		logger.Error(message)
		return shim.Error(message)
	}

	logger.Debug("Result: " + string(result))

	logger.Info("Benchmark.filter exited without errors")
	logger.Debug("Success: Benchmark.filter")
	return shim.Success(result)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string, collection string) ([]byte, error) {
	logger.Debug("getQueryResultForQueryString(" + queryString + ") is running")

	var it shim.StateQueryIteratorInterface
	var err error

	if collection != ""{
		logger.Debug("GetPrivateDataQueryResult")
		it, err = stub.GetPrivateDataQueryResult(collection, queryString)
	} else {
		logger.Debug("GetQueryResult")
		it, err = stub.GetQueryResult(queryString)
	}

	if err != nil {
		return nil, err
	}
	defer it.Close()

	entries, err := queryImpl(it, CreateTestData, stub, EmptyFilter)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	result, err := json.Marshal(entries)
	if err != nil {
		return nil, err
	}
	logger.Debug("Result: " + string(result))

	logger.Info(fmt.Sprintf("getQueryResultForQueryString(%s) exited without errors", queryString))
	logger.Debug("Success: getQueryResultForQueryString " + queryString)
	return result, nil
}