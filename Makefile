lint:
	$(info ************ RUN LINTER ************)
	golangci-lint run --timeout 5m -c .golangci.yaml

test:
	go test -race ./...