#!/usr/bin/env bash

#IP[0]=192.168.120.210
#IP[1]=192.168.120.225
#IP[2]=192.168.120.234

IP[0]=172.16.16.217
IP[1]=172.16.16.216
IP[2]=172.16.16.52

USER=support
KEY="-i  ~/.ssh/debian-server"

ORG[0]=a
ORG[1]=b
ORG[2]=c


qry[0]='{"Args":["move","a","b","10"]}'
qry[1]='{"Args":["query","a"]}'
qry[2]='{"Args":["query","b"]}'

#for ((i=0; i<${#IP[@]}; i++))
for ((i=0; i<${#IP[@]}; i++))
do
    echo -e "############\nInvoking chaincode on ${ORG[i]} node.\n############\n"
    for ((j=0; j<${#qry[@]}; j++))
    do
        echo -e "\n############\nQuery: ${qry[j]}\n############\n"
        ssh $KEY $USER@${IP[i]} -t docker exec -e CORE_PEER_ADDRESS=peer0.${ORG[i]}.example.com:7051 -ti cli.${ORG[i]}.example.com /usr/local/bin/peer chaincode invoke --cafile /etc/hyperledger/artifacts/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt --tls -o orderer.example.com:7050 -C common -n reference -c \'${qry[j]}\' | sed 's/\\"/"/g' | sed 's/\\n/\n/g'
        echo "###### Waiting for a transaction....."
        sleep 2
    done
done

for ((i=0; i<${#IP[@]}; i++))
do
    echo -e "\n############\nShowing channel info on ${ORG[i]} node.\n############\n"
    ssh $KEY $USER@${IP[i]} -t docker exec -e CORE_PEER_ADDRESS=peer0.${ORG[i]}.example.com:7051 -ti cli.${ORG[i]}.example.com /usr/local/bin/peer channel -c common getinfo
done

# docker exec -e CORE_PEER_ADDRESS=peer0.b.example.com:7051 -ti cli.b.example.com /usr/local/bin/peer chaincode invoke --cafile /etc/hyperledger/artifacts/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt --tls -o orderer.example.com:7050 -C common -n reference -c \''{"Args":["query","b"]}'\' | sed 's/\\"/"/g' | sed 's/\\n/\n/g'