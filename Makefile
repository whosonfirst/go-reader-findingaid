cli:
	go build -mod vendor -o bin/read cmd/read/main.go
	go build -mod vendor -o bin/resolverd cmd/resolverd/main.go

lambda:
	@make lambda-resolverd

lambda-resolverd:
	if test -f main; then rm -f main; fi
	if test -f resolverd.zip; then rm -f resolverd.zip; fi
	GOOS=linux go build -mod vendor -o main cmd/resolverd/main.go
	zip resolverd.zip main
	rm -f main
