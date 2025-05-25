build-server:
	go build -C bff main.go

run-server:
	go build -C bff main.go
	./bff/main
