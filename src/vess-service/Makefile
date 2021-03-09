build:
	protoc -I. --go_out=plugins=micro:./  ./proto/vess/vess.proto
buildDocker:
	docker build -t vess-service .
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vess-service