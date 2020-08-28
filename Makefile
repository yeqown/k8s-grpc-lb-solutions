TAG=v1.2.0

build-client:
	docker build -t "yeqown/lb-client.client:${TAG}" . -f client.Dockerfile
	docker push "yeqown/lb-client.client:${TAG}"

build-server:
	docker build -t "yeqown/lb-client.server:${TAG}" . -f server.Dockerfile
	docker push "yeqown/lb-client.server:${TAG}"

build: build-client build-server
	echo "all build and pushed"

gen-proto:
	# generate code by protocol buffer file
	protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/hello.proto
