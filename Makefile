build-store:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o ports-store store/main.go

build-client:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o ports-client client/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/protobuf/v1/port.proto
