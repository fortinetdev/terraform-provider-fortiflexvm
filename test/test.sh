#!/bin/sh
read -p "Please input FORTIFLEX_ACCESS_USERNAME:" host
read -p "Please input FORTIFLEX_ACCESS_PASSWORD:" password

p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
export "FORTIFLEX_ACCESS_USERNAME"=$username
export "FORTIFLEX_ACCESS_PASSWORD"=$password

echo $FORTIFLEX_ACCESS_USERNAME
echo $FORTIFLEX_ACCESS_PASSWORD

make -C ../  testacc
