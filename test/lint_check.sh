p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../flexvm
make -C ../  build

