build-server:
	env.sh
	go build -C bff ./...

run-server:
	env.sh
	go build -C bff ./...
	./bff/bff
