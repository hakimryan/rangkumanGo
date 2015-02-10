package main 

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Web server will be run...")
	http.Handle("/", http.FileServer(http.Dir("www")))
	err := http.ListenAndServe(":55", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}