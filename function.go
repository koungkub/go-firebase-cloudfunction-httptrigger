package function

import (
	"fmt"
	"net/http"
)

func Fck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "eiei")
}
