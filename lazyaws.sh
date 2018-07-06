#!/usr/bin/env bash


IP[0]=192.168.120.21
IP[1]=192.168.120.224
IP[2]=192.168.120.227
IP[3]=192.168.120.237

KEY=""
USER="ubuntu"

ORG[1]=a
ORG[2]=b
ORG[3]=c

#sudo rm -rf arts artifacts

for ((i=1; i<${#IP[@]}; i++))
do
#  ssh $KEY $USER@${IP[i]} -t "sudo apt-get remove docker docker-compose --purge -y"
  ssh $KEY $USER@${IP[i]} -t "sudo rm -rf /home/$USER/*"
  scp $KEY -r . $USER@${IP[i]}:~/fabric-starter > transfer.log

  ssh $KEY $USER@${IP[i]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m generate-peer -o ${ORG[i]}"
  #scp $KEY -r $USER@${IP[i]}:~/fabric-starter/artifacts/ arts/
  #ssh $KEY $USER@${IP[i]} -t "env | grep IP"
done


echo "########## UP ORDERER ###########"
 #ssh $KEY $USER@${IP[0]} -t "sudo rm -rf /home/support/*"
 scp $KEY -r . $USER@${IP[0]}:~/fabric-starter > transfer.log

 ssh $KEY $USER@${IP[0]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m generate-orderer"
 ssh $KEY $USER@${IP[0]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m up-orderer"

### export IP_ORDERER=172.16.16.166 IP1=172.16.16.217 IP2=172.16.16.216 IP3=172.16.16.52

for ((i=1; i<${#IP[@]}; i++))
do
#  ssh $KEY $USER@${IP[i]} -t "rm -rf ~/fabric-starter/network.sh"
#  scp $KEY -r network.sh $USER@${IP[i]}:~/fabric-starter/network.sh
  echo "########## UP NODES ###########"
  ssh $KEY $USER@${IP[i]} -t "export IP_ORDERER=${IP[0]} IP1=${IP[1]} IP2=${IP[2]} IP3=${IP[3]} && cd ~/fabric-starter && ./network.sh -m up-$i"

done