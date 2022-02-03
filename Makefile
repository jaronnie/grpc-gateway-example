.PHONY: proto
proto:
	@echo "===========> Generate grpc source codes"
	@sh ./script/proto.sh

.PHONY: build
build:
	@echo "===========> Build binary"
	@go mod tidy && go build -o cmd/core cmd/main.go

.PHONY: run
run:
	@echo "===========> run binary"
	@go mod tidy && go build -o cmd/core cmd/main.go
	@sh -c "cd cmd; ./core start -d"

.PHONY: clean
clean:
	@echo "===========> clean binary"
	@rm -f cmd/core
