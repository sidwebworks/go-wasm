build:
	GOOS=js GOARCH=wasm go build -o ./client/main.wasm
run:
	go run main.go