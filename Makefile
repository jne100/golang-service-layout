.PHONY: build

fmt:
	gofmt -l -s -w api/ cmd/ internal/

generate-proto:
	$(MAKE) generate-proto-go
	$(MAKE) generate-proto-python

generate-proto-go:
	@PATH=$(PATH):$(GOPATH)/bin protoc -I=api/ api/inventory.proto \
		--go_out=api/ \
		--go-grpc_out=api/ \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative

generate-proto-python:
	@PATH=$(PATH):$(shell python3 -m site --user-base)/bin python3 -m grpc_tools.protoc -I=api/ api/inventory.proto \
		--python_out=api/ \
		--grpc_python_out=api/ \
		--mypy_out=api/
ifeq ($(shell uname -s), Darwin)
	@sed -i "" -e "s/import inventory_pb2 as inventory__pb2/from . import inventory_pb2 as inventory__pb2/" api/inventory_pb2_grpc.py # for mac
else
	@sed -i -e "s/import inventory_pb2 as inventory__pb2/from . import inventory_pb2 as inventory__pb2/" api/inventory_pb2_grpc.py # for linux
endif

generate-mocks:
	mockgen -source=internal/controller/controller.go -destination=internal/controller/mocks/mocks.go -package=mocks Controller
	mockgen -source=internal/handler/argsvalidator/argsvalidator.go -destination=internal/handler/argsvalidator/mocks/mocks.go -package=mocks ArgsValidator

build:
	mkdir -p bin/
	CGO_ENABLED=1 go build -o bin/inventory ./cmd/service/
	@go test -run=none ./... >/dev/null

test:
	go test ./...

run-client-demo:
	go run ./cmd/clientdemo

clean:
	go clean ./...
	rm -rf bin/

images:
	docker build -f deployments/Dockerfile -t inventory:latest .

examine:
	docker run --rm -it --entrypoint /bin/sh inventory:latest

up:
	docker run -d \
		--name inventory-prod \
		--restart unless-stopped \
		-e CONFIG_PATH=/app/configs/inventory.prod.toml \
		-p 50051:50051 \
		inventory:latest

down:
	docker stop inventory-prod
	docker rm inventory-prod

deploy:
	$(MAKE) images
	$(MAKE) down
	$(MAKE) up

restart:
	docker restart inventory-prod

ps:
	@docker ps --filter "name=inventory-prod"

ssh:
	@docker exec -it inventory-prod /bin/sh

logs:
	@docker logs -f inventory-prod
