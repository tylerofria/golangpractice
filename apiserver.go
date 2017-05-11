package main

//Updatingfiletesting
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//
	"github.com/gorilla/mux"
	"github.com/skratchdot/open-golang/open"
)

//ApiInformation Comment
type ApiInformation struct {
	Name string `json:"name"`
	Food string `json:"food"`
	Job  string `json:"job"`
	//Completed bool      `json:"completed"`
	//Due       time.Time `json:"due"`
}

var m = map[string]ApiInformation{
	"Tyler": ApiInformation{"Tyler", "Pasta", "Programmer"},
	"Cindy": ApiInformation{"Cindy", "Sushi", "CFO"},
	"Rob":   ApiInformation{"Rob", "Hamburger", "CEO"},
	"David": ApiInformation{"David", "Pizza", "Accounting"},
}

type Rodos []ApiInformation

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/api/example", SimpleApi)
	router.HandleFunc("/api/{name}/json", ComplexApi)
	router.HandleFunc("/api/json", AllApi)

	//the variable that is inputed by the user.
	var nameInput string
	//get name from user
	fmt.Println("Please enter the name of the person you want information about, Rob, David, Tyler, or Cindy. To see all data, just press enter with no name.")
	fmt.Scanf("%s", &nameInput)

	//Check input for no string, so they shows all data
	if nameInput == "" {
		strPath := "http://localhost:8080//api/json"
		open.Run(strPath)
	}

	//var strPath to run the url
	strPath := "http://localhost:8080/api/" + nameInput + "/json"

	//Open the server
	open.Run(strPath)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func SimpleApi(w http.ResponseWriter, r *http.Request) {
	rodos := Rodos{
		ApiInformation{Name: "Host meetup"},

		ApiInformation{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(rodos)
}

func ComplexApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	fmt.Fprintln(w, "Typed Person:", name)
	mOutput, _ := json.Marshal(m[name])
	fmt.Fprintln(w, string(mOutput))
}

func AllApi(w http.ResponseWriter, r *http.Request) {
	mOutput, _ := json.Marshal(m)
	fmt.Fprintln(w, string(mOutput))
}

//map back end
// /api/{name}/json
// input name and get back information
