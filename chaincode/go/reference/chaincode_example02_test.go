package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestPutGet(t *testing.T) {
	cc := new(SimpleChaincode)
	stub := shim.NewMockStub("example", cc)

	stub.MockInit("1", [][]byte{[]byte("init")})

	putArgs := []string{"put", "key", "value"}
	response := stub.MockInvoke("example", toByteArray(putArgs))
	if response.Status >= 400 {
		fmt.Printf("Error on put: %s", response.Message)
		t.FailNow()
	}

	getArgs := []string{"get", "key"}
	response = stub.MockInvoke("example", toByteArray(getArgs))
	if response.Status >= 400 {
		fmt.Printf("Error on get: %s", response.Message)
		t.FailNow()
	}

	if string(response.Payload) != "value" {
		fmt.Printf("Unexpected response: %s", string(response.Payload))
		t.FailNow()
	}
}
