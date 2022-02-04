AppName=grpc-gateway-example-app

.PHONY: proto
proto:
	@echo "===========Generate grpc source codes==========="
	@sh ./script/proto.sh

.PHONY: build
build:
	@echo "===========Build application==========="
	@sh ./script/build.sh

.PHONY: run
run:
	@echo "===========Run application==========="
	@sh -c "./cmd/"${AppName}" --dir=./cmd start --nodaemon"

.PHONY: start
start:
	@echo "===========Run application with daemon==========="
	@sh -c "nohup ./cmd/"${AppName}" --dir=./cmd start --nodaemon &> "${AppName}".log &"

.PHONY: restart
restart:
	@echo "===========Restart application==========="
	@sh -c "nohup ./cmd/"${AppName}" --dir=./cmd restart &> "${AppName}".log &"

.PHONY: status
status:
	@echo "===========Get application status==========="
	@sh -c "./cmd/"${AppName}" --dir=./cmd status"

.PHONY: version
version:
	@echo "===========Get application version==========="
	@sh -c "./cmd/"${AppName}" --dir=./cmd version"

.PHONY: stop
stop:
	@echo "===========Stop application==========="
	@sh -c "./cmd/"${AppName}" --dir=./cmd stop"

.PHONY: clean
clean:
	@echo "===========Clean application==========="
	@rm -f cmd/"${AppName}"
