package main

import{
	"fmt"
	"net/http"
}

func handlerFunc(w http.ResponseWriter, r *http.Request)
{
fmt.Fprint(w,"<h1>Welcome!</h1>")
}
func main() {
	http.HandleFunc("/",handlerFunc)
	http.ListenAndServer(":3000",nil)
}