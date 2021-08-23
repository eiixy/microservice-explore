#!bin/base
echo "第一步：安装 protobuf...\r\n"
cd /tmp/
git clone --depth=1 https://github.com/protocolbuffers/protobuf
cd protobuf
./autogen.sh
./configure
make
sudo make install
protoc --version # 查看 protoc 版本，成功输出版本号，说明安装成功


# 第二步：安装 protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
