all: clean install

.PHONY: install
install:
	go install ./protoc-gen-go-http
	@echo "install finished"

clean:
	go clean ./...

