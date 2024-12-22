package internal

import (
	"fmt"
	"github.com/vaskipa/webcalculate/iternal/calculate"
	"net/http"
	"strconv"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		return http.HandleFunc()
	}

	fmt.Fprintf(w, "rpc_duration_milliseconds_count "+strconv.Itoa(requestCount))
}
func main() {
	http.HandleFunc("/api/v1/calculate", CalculateHandler)

	http.ListenAndServe(":8080", nil)
}
