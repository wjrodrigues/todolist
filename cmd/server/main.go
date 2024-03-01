package main

import (
	"fmt"
	"log"
	"net/http"
	"todolist/configs"
)

func main() {
	values, err := configs.LoadEnv(".", ".env")

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi!")
	})

	fmt.Printf("Start on %s \n", values.WebServerPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", values.WebServerPort), nil))
}
