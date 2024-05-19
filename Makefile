lint:
	$(info ************ RUN LINTER ************)
	golangci-lint run --timeout 5m -c .golangci.yaml

test:
	go test -race ./...


build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/filegrouper.exe cmd/filegrouper/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build - o bin/filegrouper cmd/filegrouper/main.go

build-macamd:
	GOOS=darwin GOARCH=amd64 go build -o bin/filegrouper cmd/filegrouper/main.go

build-macarm:
	GOOS=darwin GOARCH=arm64 go build -o bin/filegrouper cmd/filegrouper/main.go


