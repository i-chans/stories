APP=stories
APP_VERSION:="0.1"
APP_COMMIT:=$(shell git rev-parse HEAD)
APP_EXECUTABLE="./out/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

deps:
	go mod download

tidy:
	go mod tidy

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE) -ldflags "-X main.version=$(APP_VERSION) -X main.commit=$(APP_COMMIT)" cmd/*.go

build: deps compile

setup: deps migrate

clean:
	rm -rf out/

serve: build
	$(APP_EXECUTABLE) serve

migrate: build
	$(APP_EXECUTABLE) migrate

rollback: build
	$(APP_EXECUTABLE) rollback