cli:
	go build -mod vendor -o bin/read cmd/read/main.go
	go build -mod vendor -o bin/resolverd cmd/resolverd/main.go
