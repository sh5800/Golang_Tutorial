go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

%USERPROFILE%\go\bin

protoc-gen-go --version
protoc-gen-go-grpc --version

protoc --go_out=. --go-grpc_out=. proto/greet.proto
