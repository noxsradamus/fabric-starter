package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	testDataIndex = "Test"
)

const (
	testDataKeyFieldsNumber = 1
	testDataBasicArgumentsNumber = 2
)

type TestData struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Key   string `json:"key"`
	Value string `json:"value"`
}

func CreateTestData() LedgerData {
	return new(TestData)
}

func (data *TestData) FillFromArguments(args []string) error {
	if len(args) < testDataBasicArgumentsNumber {
		return errors.New(fmt.Sprintf("arguments array must contain at least %d items", testDataBasicArgumentsNumber))
	}
	data.ObjectType = testDataIndex
	data.Key = args[0]
	data.Value = args[1]

	return nil
}

func (data *TestData) FillFromCompositeKeyParts(compositeKeyParts []string) error {
	if len(compositeKeyParts) < testDataKeyFieldsNumber {
		return errors.New(fmt.Sprintf("composite key parts array must contain at least %d items", testDataKeyFieldsNumber))
	}

	data.Key = compositeKeyParts[0]

	return nil
}

func (data *TestData) FillFromLedgerValue(ledgerBytes []byte) error {
	if err := json.Unmarshal(ledgerBytes, &data); err != nil {
		return err
	} else {
		return nil
	}
}

func (data *TestData) ToCompositeKey(stub shim.ChaincodeStubInterface) (string, error) {
	compositeKeyParts := []string {
		data.Key,
	}

	return stub.CreateCompositeKey(testDataIndex, compositeKeyParts)
}

func (data *TestData) ToLedgerValue() ([]byte, error) {
	return json.Marshal(data)
}

