#!/usr/bin/env bash

source ./network.sh

COMMON_POLICY=""
CH_AB_POLICY=""
CH_AC_POLICY=""

chaincode_version="$(od -vAn -N4 -tu4 < /dev/urandom)"

echo "=== Upgrading all chaincodes with $chaincode_version"

for org in ${ORG1} ${ORG2} ${ORG3}
do
    for chaincode_name in ${CHAINCODE_COMMON_NAME} ${CHAINCODE_BILATERAL_NAME}
    do
      installChaincode "${org}" "${chaincode_name}" $chaincode_version
    done
done

# Upgrade chaincodes

upgradeChaincode "${ORG1}" "common" "$chaincode_version" "${CHAINCODE_COMMON_NAME}" "${CHAINCODE_COMMON_INIT}" "$COMMON_POLICY" "${COLLECTION_CONFIG}"
upgradeChaincode "${ORG1}" "common" "$chaincode_version" "${CHAINCODE_BILATERAL_NAME}" "${CHAINCODE_BILATERAL_INIT}" "$COMMON_POLICY" "${COLLECTION_CONFIG}"
upgradeChaincode "${ORG1}" "${ORG1}-${ORG2}" "$chaincode_version" "${CHAINCODE_BILATERAL_NAME}" "${CHAINCODE_BILATERAL_INIT}" "$CH_AB_POLICY" "${COLLECTION_CONFIG}"
upgradeChaincode "${ORG1}" "${ORG1}-${ORG3}" "$chaincode_version" "${CHAINCODE_BILATERAL_NAME}" "${CHAINCODE_BILATERAL_INIT}" "$CH_AC_POLICY" "${COLLECTION_CONFIG}"

echo "=== Testing chaincodes"

docker exec cli.${ORG1}.${DOMAIN} "CORE_PEER_ADDRESS=peer0.${ORG1}.${DOMAIN}:7051 && peer chaincode invoke -n ${CHAINCODE_COMMON_NAME} -c '{\"Args\":[\"invokeChaincode\",\"a-b\",\"relationship\",\"move\",\"a\",\"b\",\"15\"]}' -o orderer.$DOMAIN:7050 -C common --tls --cafile /etc/hyperledger/crypto/orderer/tls/ca.crt"
docker exec cli.${ORG1}.${DOMAIN} "CORE_PEER_ADDRESS=peer0.${ORG1}.${DOMAIN}:7051 && peer chaincode query -n ${CHAINCODE_COMMON_NAME} -c '{\"Args\":[\"getResponse\",\"getResponse\"]}' -o orderer.$DOMAIN:7050 -C common --tls --cafile /etc/hyperledger/crypto/orderer/tls/ca.crt"
