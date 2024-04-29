package main

import "os"

func main() {
	defer println("deferred")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	callDeferred()

	println("not deferred")

	os.Exit(0)

	// http.ListenAndServe(":8080", nil)
}

func callDeferred() {
	defer println("close defer") // 2

	println("callDeferred") // 1

}
