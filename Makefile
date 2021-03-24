# apps make file
# simple command to run the apps

test:
	@go test ./... -v -cover -race -count=1
run:
	@go run cmd/pundi/main.go
build:
	@go build -o dompet cmd/pundi/main.go

.PHONY: default