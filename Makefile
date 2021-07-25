VERSION=0.0.3
PATH_BUILD=build/
FILE_COMMAND=asar
FILE_ARCH=darwin_amd64

clean:
	@rm -rf ./build

build: clean
	@$(GOPATH)/bin/goxc \
		-bc="darwin,amd64" \
		-pv=$(VERSION) \
		-d=$(PATH_BUILD) \
		-build-ldflags "-X main.VERSION=$(VERSION)"

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)$(VERSION)/$(FILE_ARCH)/$(FILE_COMMAND) '$(HOME)/bin/$(FILE_COMMAND)'

try:
	~/bin/asar version

test:
	go test -cover ./...

test-cov:
	go test -race -covermode=atomic -coverprofile=covprofile ./...

test-ci: test-cov cov-func

cov-htm:
	go tool cover -html=covprofile

cov-func:
	go tool cover -func=covprofile

try-env:
	export ASAR_PROJECT_BASE=project_base_test && export ASAR_PORT=80 && go run main.go version

try-file:
	go run main.go --config ./.env version

ast:
	go run main.go ast

ast-c:
	go run main.go --config ./.env ast

apib:
	yarn --cwd ./mock mock

asaa:
	go run main.go asaa
