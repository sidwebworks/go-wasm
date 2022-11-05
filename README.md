## Usage (default)

### Building the binary

```
GOOS=js GOARCH=wasm go build -o ./client/main.wasm
```

## Copying get the `wasm_exec.js` script

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./client
```

## Usage with Tinygo

> Requires tinygo installed

You can use tinygo to get a somewhat okayish output binary size but for some reason it's not working properly for the http example in this repo. I might take a look at it later and push a fix.

## Building the binary

```
tinygo build -target wasm -o ./client/main.wasm
```

## Copying the tinygo `wasm_exec.js` script

```sh
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./client
```

### More tinified builds

You can use the `--no-debug` flag to remove debug symbols in the output binary

```
tinygo build -target wasm -o ./client/main.wasm --no-debug
```

All the above build commands are in the `Makefile` too
