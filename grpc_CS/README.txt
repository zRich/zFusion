1. 安装配置go语言开发环境

2. 配置protobuf
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+

PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip

unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local.

export PATH="$PATH:$HOME/.local/bin"

3. 下载编译插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"

4. 运行server端
cd grpc_CS/server
go run server.go

5. 运行client端
cd grpc_CS/client
go run client.go