OUT=./client/main.wasm

build:
	GOOS=js GOARCH=wasm go build -o $(OUT)

build-lean:
	tinygo build -target wasm -o $(OUT)

build-tiny:
	tinygo build -target wasm -o $(OUT) --no-debug