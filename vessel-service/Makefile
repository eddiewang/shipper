.PHONY: proto
build: proto docker

proto:
	protoc -I. --micro_out=. --go_out=. proto/vessel/vessel.proto
docker:
	docker build -t vessel-service .

run: 
	docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns vessel-service