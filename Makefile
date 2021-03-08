#build docker image
build:
	docker build -t consignment-service .
#run as rpc
run:
	docker run -p 50051:50051 consignment-service
#run as microservice
runMicro:
	docker run -p 50051:50051 \
         -e MICRO_SERVER_ADDRESS=:50051 \
         -e MICRO_REGISTRY=mdns \
         consignment-service
