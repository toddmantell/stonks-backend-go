package stonksbackend

import (
	"fmt"
	"net/http"
	s "strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	// router.HandleFunc("/stonks", StonksHandler)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	value := GrahamFormula(2.1, 10.0, 2.3)
	// This is just to make this work, will need to figure out how to output all info
	fmt.Printf("%s", s.Itoa(int(value.highConservativeGrahamFormulaNumber)))
}
