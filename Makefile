build:
	go build -o server.bin server.go

run: build
	./server.bin

watch:
	reflex -s -r '\.go$$' make run