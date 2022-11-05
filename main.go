package main

import (
	"fmt"
	"io"
	"net/http"
	"syscall/js"
)

type PromiseHandle = func(resolve js.Value, reject js.Value, args []js.Value)

func main() {
	fmt.Println("Go Webassembly initialized")

	js.Global().Set("request", promisify(func(resolve js.Value, reject js.Value, args []js.Value) {
		fmt.Println("Go recieved function args: ", args)

		url := (args)[0].String()

		res, err := http.Get(url)

		errConstructor := js.Global().Get("Error")

		if err != nil {
			obj := errConstructor.New(err.Error())
			reject.Invoke(obj)
		}

		data, err := io.ReadAll(res.Body)

		defer res.Body.Close()

		if err != nil {
			obj := errConstructor.New(err.Error())
			reject.Invoke(obj)
		}

		resolve.Invoke(string(data))
	}))

	select {}
}

func promisify(fn PromiseHandle) js.Func {
	return js.FuncOf(func(this js.Value, input []js.Value) any {
		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			go fn(resolve, reject, input)

			return nil
		})

		promise := js.Global().Get("Promise")

		return promise.New(handler)
	})
}
