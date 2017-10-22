export GOROOT=`pwd`/src/github.com/go
echo $GOROOT
export GOPATH=`pwd`
echo $GOPATH
cd src/siwei.com/TCPClient
go build main.go
./main
