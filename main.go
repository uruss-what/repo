package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name  string
	Age   uint16
	Money float64
}

/*
	func (u User) getAllInfo() string {
		return fmt.Sprintf("user name is: %s. He is %d  and he has "+
			"%d $", u.Name, u.Age, u.Money)
	}
*/
func (u User) setNewName(newName string) {
	u.Name = newName
}
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -3000}
	bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is a contacts page!")
}

func HandleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":2023", nil)
}

func main() {

	HandleRequest()

}
