package main

import (  
"fmt"
"net/http"

)


func handler(w http.ResponseWriter, r *http.Request) {  
            fmt.Fprintf(w,"%s","Hello World")
}


func main() {  
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8086", nil)
}