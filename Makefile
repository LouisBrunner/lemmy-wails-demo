NPM = npm --prefix frontend

all: lint
.PHONY: all

generate:
# handle issues when generating from scratch
	mkdir -p frontend/dist
	touch frontend/dist/.tmp
	wails generate module
.PHONY: generate

dev: generate
	wails dev
.PHONY: dev

build: generate
	wails build
.PHONY: build

build-all: generate
	wails build -platform darwin/amd64,darwin/arm64,windows/amd64,linux/amd64
.PHONY: build-all

install:
	go mod download
	$(NPM) i
	make doctor
.PHONY: install

doctor:
	wails doctor
.PHONY: doctor

lint-ts: generate
	$(NPM) run format
	$(NPM) run lint
	$(NPM) run types
.PHONY: lint-ts

lint-go: generate
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
.PHONY: lint-go

lint: lint-go lint-ts
.PHONY: lint

format-fix-ts:
	$(NPM) run format-fix
.PHONY: format-fix-ts

test: generate
.PHONY: test

clean:
	rm -rf build/bin frontend/dist frontend/wailsjs
.PHONY: clean
