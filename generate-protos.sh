protoc -I=protobuf --go_out=server/conversation --go_opt=paths=source_relative --go-grpc_out=server/conversation --go-grpc_opt=paths=source_relative protobuf/chat.proto


