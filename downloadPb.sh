export GOROOT=/usr/local/go
export GOPATH=/home/siwei/GoglandProjects/beegoTest
go get github.com/golang/protobuf/protoc-gen-go
cd src/github.com/golang/protobuf/protoc-gen-go
go build
go install

