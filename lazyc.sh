#!/usr/bin/env bash


IP[0]=192.168.120.226
IP[1]=192.168.120.210
IP[2]=192.168.120.225
IP[3]=192.168.120.234


ORG[1]=a
ORG[2]=b
ORG[3]=c

#rm -rf arts

for ((i=1; i<${#IP[@]}; i++))
do
#  ssh -i ~/debian-server centos@${IP[i]} -t "sudo apt-get remove docker docker-compose --purge -y"
  ssh centos@${IP[i]} -t "sudo rm -rf /home/centos/fabric-starter"
  scp -r . centos@${IP[i]}:~/fabric-starter > transfer.log

  ssh centos@${IP[i]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m generate-peer -o ${ORG[i]}"
  #scp -i ~/debian-server -r centos@${IP[i]}:~/fabric-starter/artifacts/ arts/
  #ssh -i ~/debian-server centos@${IP[i]} -t "env | grep IP"
done

#cp -r arts/artifacts artifacts

echo "########## UP ORDERER ###########"
 ssh centos@${IP[0]} -t "sudo rm -rf /home/centos/fabric-starter"
 scp -r . centos@${IP[0]}:~/fabric-starter > transfer.log

 ssh centos@${IP[0]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m generate-orderer"
 ssh centos@${IP[0]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m up-orderer"

### export IP_ORDERER=172.16.16.166 IP1=172.16.16.217 IP2=172.16.16.216 IP3=172.16.16.52

for ((i=1; i<${#IP[@]}; i++))
do

  echo "########## UP NODES ###########"
  ssh centos@${IP[i]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m up-$i"

done