VERSION=0.0.6
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

install.script:
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)$(VERSION)/$(FILE_ARCH)/$(FILE_COMMAND) '$(HOME)/bin/$(FILE_COMMAND)'

install:
	make build && make install.script

try:
	~/bin/asar version

test:
	go test -v -cover ./...

test.cov:
	go test -v -race -covermode=atomic -coverprofile=coverage.out ./...

test.ci: test.cov cov.func

cov.htm:
	go tool cover -html=coverage.out

cov.func:
	go tool cover -func=coverage.out

try.env:
	export ASAR_PROJECT_BASE=project_base_test && export ASAR_PORT=80 && go run main.go version

try.file:
	go run main.go --config ./.env version

ast:
	go run main.go ast

ast.c:
	go run main.go --config ./.env ast

apib:
	yarn --cwd ./mock mock

asaa:
	go run main.go asaa

rasa:
	go run main.go rasa
