package main

import (
	"fmt"
	"io"
	"net/http"
	"syscall/js"
)

type PromiseHandle = func(resolve js.Value, reject js.Value)

func main() {
	fmt.Println("Yo yo yo")

	js.Global().Set("request", promisify(func(resolve js.Value, reject js.Value) {

		res, err := http.Get("https://jsonplaceholder.typicode.com/todos")

		if err != nil {
			reject.Invoke("Shit happend")
		}

		data, err := io.ReadAll(res.Body)

		if err != nil {
			reject.Invoke("Shit happend")
		}

		resolve.Invoke(string(data))
	}))

	select {}
}

func promisify(fn PromiseHandle) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go fn(resolve, reject)

			return nil
		})

		promise := js.Global().Get("Promise")

		return promise.New(handler)
	})
}
