protoc --proto_path ../../../ -I=./proto --go_out=plugins=grpc:./proto proto/versionserver.proto
mv proto/github.com/brotherlogic/versionserver/proto/* ./proto
