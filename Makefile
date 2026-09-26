.PHONY: all build test lint clean install deb iso release

VERSION ?= 0.1.0-dev
ARCH ?= amd64

all: build

build:
	./build/build.sh

test:
	cd cli && go test ./... -v -race
	bats tests/

lint:
	cd cli && golangci-lint run
	shellcheck build/*.sh scripts/*.sh packages/*.sh

clean:
	rm -rf build/arivon build/arivon-arm64 build/arivon-installer build/deb
	rm -rf images/output/

install: build
	sudo cp build/arivon /usr/bin/arivon
	sudo cp build/arivon-installer /usr/bin/arivon-installer
	sudo chmod 755 /usr/bin/arivon /usr/bin/arivon-installer

deb:
	./packages/build-deb.sh $(VERSION) $(ARCH)

iso:
	cd images && sudo mkosi build --profile minimal

release:
	./build/release.sh $(VERSION)
