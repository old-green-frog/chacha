package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//User structure
type User struct {
	Name     string
	Password string
	Status   uint8
}

func postPage(w http.ResponseWriter, r *http.Request) {

	user := User{}

	user.Name = r.FormValue("login")
	user.Password = r.FormValue("password")
	user.Status = 1

	//					   add your user:password@/base
	db, err := sql.Open("mysql", "gotest:gotest@/start")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Installing data
	sel, err := db.Query(fmt.Sprintf("SELECT `logn`, `passwd` FROM Users WHERE `logn` LIKE '%s' AND `passwd` LIKE '%s';", user.Name, user.Password))
	if err != nil {
		user.Status = 0
		panic(err)
	}

	defer sel.Close()

	tmpl, _ := template.ParseFiles("../../templates/root.html")
	tmpl.Execute(w, user)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../../templates/root.html")
	tmpl.Execute(w, User{"user", "0000", 3})
}

func main() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/postform", postPage)
	http.ListenAndServe(":8000", nil)
}
