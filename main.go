package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type data struct {
	Id       string
	FullName string
	Age      string
}

type coba struct {
	Nama string
}

var Datas []data
var Hasil = []coba{}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome To My Restfull Api")
}

func showSingleData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, data := range Datas {
		if data.Id == key {
			json.NewEncoder(w).Encode(data)
		}
	}
}

func newData(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newData data
	json.Unmarshal(reqBody, &newData)
	Datas = append(Datas, newData)
	json.NewEncoder(w).Encode(newData)
}
func newSql(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newSql coba

	if r.Method == "POST" {
		json.Unmarshal(reqBody, &newSql)
		Hasil = append(Hasil, newSql)
		db.Exec("insert into user (nama) values (?)", newSql.Nama)
		json.NewEncoder(w).Encode(newSql)
	}

}

func showSql(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Hasil)
}
func showData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Datas)
}

func main() {
	r := mux.NewRouter()
	// Hasil = []coba{}
	Datas = []data{{Id: "1", FullName: "Ummul Qoyimah", Age: "21"}, {Id: "2", FullName: "Dewi Novita Sari", Age: "20"}}
	r.HandleFunc("/", user)
	r.HandleFunc("/show", showData)
	http.Handle("/", r)
	r.HandleFunc("/show/{id}", showSingleData)
	r.HandleFunc("/sql", showSql)
	r.HandleFunc("/addData", newData).Methods("POST")
	r.HandleFunc("/addSql", newSql)

	sqlQuery()

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on: http://localhost:8080")
	log.Println("Server started on: http://localhost:8080")
}

func connect() *sql.DB {
	db, _ := sql.Open("mysql", "root:Ilovereza123@/coba")
	return db
}

func sqlQuery() {
	db := connect()
	defer db.Close()

	rows, err := db.Query("select nama from user")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var each = coba{}
		rows.Scan(&each.Nama)
		Hasil = append(Hasil, each)
	}
	// fmt.Print(result)

	for _, each := range Hasil {
		fmt.Println(each.Nama)
	}

}
