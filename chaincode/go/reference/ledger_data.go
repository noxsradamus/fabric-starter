package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"errors"
	"encoding/json"
)

var ledgerDataLogger = shim.NewLogger("LedgerData")

type LedgerData interface {
	FillFromArguments(args []string) error

	FillFromCompositeKeyParts(compositeKeyParts []string) error

	FillFromLedgerValue(ledgerValue []byte) error

	ToCompositeKey(stub shim.ChaincodeStubInterface) (string, error)

	ToLedgerValue() ([]byte, error)
}

func ExistsIn(stub shim.ChaincodeStubInterface, data LedgerData) bool {
	compositeKey, err := data.ToCompositeKey(stub)
	if err != nil {
		return false
	}

	if data, err := stub.GetState(compositeKey); err != nil || data == nil {
		return false
	}

	return true
}

func LoadFrom(stub shim.ChaincodeStubInterface, data LedgerData) error {
	compositeKey, err := data.ToCompositeKey(stub)
	if err != nil {
		return err
	}

	bytes, err := stub.GetState(compositeKey)
	if err != nil {
		return err
	}

	return data.FillFromLedgerValue(bytes)
}

func UpdateOrInsertIn(stub shim.ChaincodeStubInterface, data LedgerData) error {
	compositeKey, err := data.ToCompositeKey(stub)
	if err != nil {
		return err
	}

	value, err := data.ToLedgerValue()
	if err != nil {
		return err
	}

	if err = stub.PutState(compositeKey, value); err != nil {
		return err
	}

	return nil
}

func UpdateOrInsertInCollection(stub shim.ChaincodeStubInterface, data LedgerData, collection string) error {
	compositeKey, err := data.ToCompositeKey(stub)
	if err != nil {
		return err
	}

	value, err := data.ToLedgerValue()
	if err != nil {
		return err
	}

	if err = stub.PutPrivateData(collection, compositeKey, value); err != nil {
		return err
	}

	return nil
}

type FactoryMethod func() LedgerData

type FilterFunction func(data LedgerData) bool

func EmptyFilter(data LedgerData) bool {
	return true
}

func Query(stub shim.ChaincodeStubInterface, index string, partialKey []string,
	createEntry FactoryMethod, filterEntry FilterFunction) ([]byte, error) {

	ledgerDataLogger.Info(fmt.Sprintf("Query(%s) is running", index))
	ledgerDataLogger.Debug("Query " + index)

	it, err := stub.GetStateByPartialCompositeKey(index, partialKey)
	if err != nil {
		message := fmt.Sprintf("unable to get state by partial composite key %s: %s", index, err.Error())
		ledgerDataLogger.Error(message)
		return nil, errors.New(message)
	}
	defer it.Close()

	entries := []LedgerData{}
	for it.HasNext() {
		response, err := it.Next()
		if err != nil {
			message := fmt.Sprintf("unable to get an element next to a query iterator: %s", err.Error())
			ledgerDataLogger.Error(message)
			return nil, errors.New(message)
		}

		ledgerDataLogger.Debug(fmt.Sprintf("Response: {%s, %s}", response.Key, string(response.Value)))

		entry := createEntry()

		if err := entry.FillFromLedgerValue(response.Value); err != nil {
			message := fmt.Sprintf("cannot fill entry value from response value: %s", err.Error())
			ledgerDataLogger.Error(message)
			return nil, errors.New(message)
		}

		_, compositeKeyParts, err := stub.SplitCompositeKey(response.Key)
		if err != nil {
			message := fmt.Sprintf("cannot split response key into composite key parts slice: %s", err.Error())
			ledgerDataLogger.Error(message)
			return nil, errors.New(message)
		}

		if err := entry.FillFromCompositeKeyParts(compositeKeyParts); err != nil {
			message := fmt.Sprintf("cannot fill entry key from composite key parts: %s", err.Error())
			ledgerDataLogger.Error(message)
			return nil, errors.New(message)
		}

		if bytes, err := json.Marshal(entry); err == nil {
			ledgerDataLogger.Debug("Entry: " + string(bytes))
		}

		if filterEntry(entry) {
			entries = append(entries, entry)
		}
	}

	result, err := json.Marshal(entries)
	if err != nil {
		return nil, err
	}
	ledgerDataLogger.Debug("Result: " + string(result))

	ledgerDataLogger.Info(fmt.Sprintf("Query(%s) exited without errors", index))
	ledgerDataLogger.Debug("Success: Query " + index)
	return result, nil
}

// TODO: refactor code to remove duplicates
func QueryPrivate(stub shim.ChaincodeStubInterface, index string, partialKey []string,
	createEntry FactoryMethod, filterEntry FilterFunction, collections []string) ([]byte, error) {

	ledgerDataLogger.Info(fmt.Sprintf("QueryPrivate(%s) is running", index))
	ledgerDataLogger.Debug("QueryPrivate " + index)

	entries := []LedgerData{}
	for _, collection := range collections {
		it, err := stub.GetPrivateDataByPartialCompositeKey(collection, index, partialKey)
		if err != nil {
			message := fmt.Sprintf("unable to get state by partial composite key %s: %s", index, err.Error())
			ledgerDataLogger.Error(message)
			return nil, errors.New(message)
		}

		for it.HasNext() {
			response, err := it.Next()
			if err != nil {
				message := fmt.Sprintf("unable to get an element next to a query iterator: %s", err.Error())
				ledgerDataLogger.Error(message)
				return nil, errors.New(message)
			}

			ledgerDataLogger.Debug(fmt.Sprintf("Response: {%s, %s}", response.Key, string(response.Value)))

			entry := createEntry()

			if err := entry.FillFromLedgerValue(response.Value); err != nil {
				message := fmt.Sprintf("cannot fill entry value from response value: %s", err.Error())
				ledgerDataLogger.Error(message)
				return nil, errors.New(message)
			}

			_, compositeKeyParts, err := stub.SplitCompositeKey(response.Key)
			if err != nil {
				message := fmt.Sprintf("cannot split response key into composite key parts slice: %s", err.Error())
				ledgerDataLogger.Error(message)
				return nil, errors.New(message)
			}

			if err := entry.FillFromCompositeKeyParts(compositeKeyParts); err != nil {
				message := fmt.Sprintf("cannot fill entry key from composite key parts: %s", err.Error())
				ledgerDataLogger.Error(message)
				return nil, errors.New(message)
			}

			if bytes, err := json.Marshal(entry); err == nil {
				ledgerDataLogger.Debug("Entry: " + string(bytes))
			}

			if filterEntry(entry) {
				entries = append(entries, entry)
			}
		}

		it.Close()
	}

	result, err := json.Marshal(entries)
	if err != nil {
		return nil, err
	}
	ledgerDataLogger.Debug("Result: " + string(result))

	ledgerDataLogger.Info(fmt.Sprintf("QueryPrivate(%s) exited without errors", index))
	ledgerDataLogger.Debug("Success: QueryPrivate " + index)
	return result, nil
}