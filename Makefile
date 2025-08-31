GO_TEST_ARGS ?= -v --count=1

build:
	go build -o ./bin/tasklist .
start:
	./bin/tasklist
dev:
	go tool -modfile=go.tool.mod air
format:
	gofmt -w .
test:
	PARALLEL_CONVEY=false make test-lightspeed
test-lightspeed:
	go test $(GO_TEST_ARGS) --run='${run}' -v --count=1 ./tests/...
test-coverage:
	@mkdir -p ./coverage
	make test GO_TEST_ARGS="--cover -coverpkg=./cmd/...,./core/...,./plugins/...,./utils/... --coverprofile=./coverage/coverage.out"
	go tool cover -html=./coverage/coverage.out -o ./coverage/index.html
	@echo "\033[0;32mCoverage report generated at ./coverage/index.html.\033[0m"
benchmark:
	go test -bench=. $(GO_BENCH_ARGS) -benchmem -tags=benchmark ./tests/benchmarks/... 
lint:
	golangci-lint run ./...
lint-fix:
	golangci-lint run --fix ./...
install:
	go mod tidy
	@echo "\033[0;32mGo modules installed successfully.\033[0m"
	mkdir -p tools && mv ./go.tool.mod ./tools/go.mod && cd ./tools && go mod tidy && cd .. && mv ./tools/go.mod ./go.tool.mod && mv ./tools/go.sum ./go.tool.sum && rm -rf ./tools
	@echo "\033[0;32mGo tools installed successfully.\033[0m"
	go tool -modfile=go.tool.mod lefthook install
	@echo "\033[0;32mLefthook configured successfully.\033[0m"
	@which npm > /dev/null && \
		npm install -g @commitlint/config-conventional@17.6.5 @commitlint/cli@17.6.5 && \
		echo "\033[0;32mCommitlint installed successfully.\033[0m" || \
		echo "\033[0;31mNode is not installed. Please install Node.js to use commitlint.\033[0m"
sandbox:
	docker compose -f ./infrastructure/docker-compose.yml up
teardown:
	docker compose -f ./infrastructure/docker-compose.yml down
mocks:
	go tool -modfile=go.tool.mod mockery --config ./tests/mocks/mockery.yml
swagger:
	go tool -modfile=go.tool.mod swag init -g main.go -o ./docs --outputTypes json