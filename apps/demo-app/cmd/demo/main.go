package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "hello from demo-app")
	})

	fmt.Println("demo-app listening on :8081")
	_ = http.ListenAndServe(":8081", nil)
}
