#!/usr/bin/env bash
FOLDER=$1
if [ -z $FOLDER ]; then
FOLDER=test
fi

for file in $(find . -wholename "*/network/fabric/$FOLDER/*Admin*/msp/keystore/*")
do
  echo "renaming $file to $(echo $file | sed s#/[a-z0-9]*_sk#/key.pem#)"
  mv $file `echo $file | sed s#/[a-z0-9]*_sk#/key.pem#` -f
done