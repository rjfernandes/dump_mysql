info:
	@echo "Usage: make [target]"

build:
	@echo "Building..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backup_db.mrf -a -ldflags '-extldflags "-static"' ./cmd/main.go

b_local:
	@echo "Building..."
	@go build -o backup_db.mrf ./cmd/main.go