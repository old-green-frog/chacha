package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//					   add your user:password@/base
	db, _ = sql.Open("mysql", "mysql:mysql@/start")
)

//User structure
type User struct {
	Name     string
	Password string
	Status   uint8
}

func postPage(w http.ResponseWriter, r *http.Request) {

	user := User{}

	user.Name = strings.Trim(r.FormValue("login"), " ")
	user.Password = strings.Trim(r.FormValue("password"), " ")
	user.Status = 1

	//Installing data
	err := db.QueryRow(fmt.Sprintf("SELECT login, password FROM Users WHERE login='%s' AND password='%s'", user.Name, user.Password)).Scan()
	if err == sql.ErrNoRows {
		user.Status = 0
	}

	tmpl, _ := template.ParseFiles("../../templates/root.html")
	tmpl.Execute(w, user)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../../templates/root.html")
	tmpl.Execute(w, User{"user", "0000", 3})
}

func main() {

	defer db.Close()

	http.HandleFunc("/", rootPage)
	http.HandleFunc("/postform", postPage)
	http.ListenAndServe(":8080", nil)
}
