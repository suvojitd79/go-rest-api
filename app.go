package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
}

func generateData(n int) (*[]Employee, error) {

	emp := []Employee{}

	for i := 0; i < n; i++ {

		emp = append(emp, Employee{ID: int64(i + 1), Name: generateName(), Designation: generateName()[:5]})

	}

	return &emp, nil

}

func generateName() string {

	data := `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

	f := getRandom(5, 10)

	n := make([]byte, f)

	for x := 0; x < f; x++ {

		n[x] = byte(data[getRandom(0, len(data))])
	}

	n[f/2] = ' '

	return string(n)
}

func getRandom(x, y int) int {
	return x + rand.Intn(y-x)
}

var E []Employee

func init() {

	emp, e := generateData(50)

	if e != nil {
		fmt.Println(e)
		return
	}

	E = *emp

	fmt.Printf("running server..... %d", E[0].ID)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/{id:[0-9]{1,}}", getEmployeeById).Methods("GET")
	router.HandleFunc("/employees", getAllEmployee).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println(E)

}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))

}

func getEmployeeById(w http.ResponseWriter, r *http.Request) {

	map_ := mux.Vars(r)
	emp_id, err := strconv.ParseInt(map_["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(make(map[string]string))
		return
	}

	w.Header().Set("content-type", "application/json")

	for _, val := range E {
		if val.ID == emp_id {
			json.NewEncoder(w).Encode(val)
			return
		}
	}

	json.NewEncoder(w).Encode(make(map[string]string))

}

func getAllEmployee(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(E)

}
