
package main
 
import (
    "fmt"
    "log"
    "net/http"
)
 
func main() {
 	http.Handle("/", http.FileServer(http.Dir("./web")))
 	port := ":80"
    fmt.Println("Server is running on port" + port)
 	log.Fatal(http.ListenAndServe(port, nil))
}