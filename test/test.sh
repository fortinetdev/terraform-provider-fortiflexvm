#!/bin/sh
read -p "Please input FLEXVM_ACCESS_USERNAME:" host
read -p "Please input FLEXVM_ACCESS_PASSWORD:" password

p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
export "FLEXVM_ACCESS_USERNAME"=$username
export "FLEXVM_ACCESS_PASSWORD"=$password

echo $FLEXVM_ACCESS_USERNAME
echo $FLEXVM_ACCESS_PASSWORD

make -C ../  testacc
