GO_TEST_ARGS ?= -v --count=1
COVERAGEDIR = coverage

build:
	go build -o ./bin/tasklist .
start:
	./bin/tasklist
dev:
	go tool -modfile=go.tool.mod air
fmt:
	gofmt -w .
test:
	PARALLEL_CONVEY=false make test-lightspeed
test-lightspeed:
	go test $(GO_TEST_ARGS) --run='${run}' -v --count=1 ./...
test-coverage:
	@mkdir -p $(COVERAGEDIR)
	make test GO_TEST_ARGS="--cover -coverpkg=./app/...,./config/...,./global/...,./middleware/...,./modules/...,./pkg/...,./utils/... --coverprofile=$(COVERAGEDIR)/coverage.out.tmp"
	cat $(COVERAGEDIR)/coverage.out.tmp | grep -v "**/*_mock" > $(COVERAGEDIR)/coverage.out
	rm -f $(COVERAGEDIR)/coverage.out.tmp
	go tool cover -html=$(COVERAGEDIR)/coverage.out -o $(COVERAGEDIR)/index.html
	@echo "\033[0;32mCoverage report generated at $(COVERAGEDIR)/index.html.\033[0m"
test-ui:
	go tool -modfile=go.tool.mod goconvey -port=8081 -excludedDirs="bin,coverage,database,docs,tests" ./...
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
	go tool -modfile=go.tool.mod swag fmt main.go
	go tool -modfile=go.tool.mod swag init -g main.go -o ./docs --outputTypes json