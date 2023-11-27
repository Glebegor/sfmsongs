run:
	go run cmd/main.go
install-depends:
	go mod download
build-win:
	go build cmd/main.go
