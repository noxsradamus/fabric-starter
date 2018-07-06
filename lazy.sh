#!/usr/bin/env bash


IP[0]=172.16.16.166
IP[1]=172.16.16.39
IP[2]=172.16.16.216
IP[3]=172.16.16.131


ORG[1]=a
ORG[2]=b
ORG[3]=c

#sudo rm -rf arts artifacts

for ((i=1; i<${#IP[@]}; i++))
do
#  ssh -i ~/debian-server support@${IP[i]} -t "sudo apt-get remove docker docker-compose --purge -y"
  ssh -i ~/.ssh/debian-server support@${IP[i]} -t "sudo rm -rf /home/support/*"
  scp -i  ~/.ssh/debian-server -r . support@${IP[i]}:~/fabric-starter > transfer.log

  ssh -i  ~/.ssh/debian-server support@${IP[i]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m generate-peer -o ${ORG[i]}"
  #scp -i ~/debian-server -r support@${IP[i]}:~/fabric-starter/artifacts/ arts/
  #ssh -i ~/debian-server support@${IP[i]} -t "env | grep IP"
done


echo "########## UP ORDERER ###########"
 #ssh -i  ~/.ssh/debian-server support@${IP[0]} -t "sudo rm -rf /home/support/*"
 #scp -i  ~/.ssh/debian-server -r . support@${IP[0]}:~/fabric-starter > transfer.log

 ssh -i ~/.ssh/debian-server support@${IP[0]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m generate-orderer"
 ssh -i ~/.ssh/debian-server support@${IP[0]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m up-orderer"

### export IP_ORDERER=172.16.16.166 IP1=172.16.16.217 IP2=172.16.16.216 IP3=172.16.16.52

for ((i=1; i<${#IP[@]}; i++))
do
  echo "########## UP NODES ###########"
  ssh -i ~/.ssh/debian-server support@${IP[i]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m up-$i"

done