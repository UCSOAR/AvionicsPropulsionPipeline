package function

import (
	"fmt"
	"net/http"
)

// HelloWorld is an HTTP Cloud Function.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
