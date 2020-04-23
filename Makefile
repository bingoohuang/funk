test:
	go test -v

bench:
	go test -benchmem -bench .

lint:
	gofmt -s -w .&&go mod tidy&&go fmt ./...&&revive .&&goimports -w .&&golangci-lint run --enable-all&&go install -ldflags="-s -w" ./...