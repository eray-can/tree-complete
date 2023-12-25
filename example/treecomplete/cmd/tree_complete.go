package main

import (
	"fmt"
	"github.com/eray-can/tree-complete/example/treecomplete"
	"net/http"
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
