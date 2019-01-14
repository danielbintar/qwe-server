bin:
	go build -o qwe-server app/http/main.go

run: bin
	./qwe-server
