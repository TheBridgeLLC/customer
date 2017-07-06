PROTOC=protoc

all:
	$(PROTOC) --go_out=plugins=grpc:. customer.proto
