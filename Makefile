proto:
	protoc --go_out=. --go-grpc_out=. ./internal/orb/proto/orb.proto

serve:
	go run . serve

