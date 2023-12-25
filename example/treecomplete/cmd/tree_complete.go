package main

import (
	"fmt"
	"net/http"
	"tree-complete/example/treecomplete"
)

func main() {
	http.HandleFunc("/tree-complete", treecomplete.Handler)
	err := http.ListenAndServe(":3131", nil)
	if err != nil {
		fmt.Printf("err server start %s", err.Error())
		return
	}

	fmt.Println("started")

}
