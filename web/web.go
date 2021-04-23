// nodemon --exec go run main.go --signal SIGTERM

package web

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Page struct {
	Title string
}

type User struct {
	Userask string
	Fname   string
	Lname   string
}

type Status struct {
	Title  string
	Status string
}

func reqHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/status" {
		s := Status{"Everything OK", "200"}
		statusPage, _ := template.ParseFiles("web/status.html")
		statusPage.Execute(w, s)
		statusPage.Execute(w, s)
	} else if r.URL.Path == "/" {
		page := Page{"Page " + time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")}
		switch r.Method {
		case "GET":
			fmt.Println("GET Received")
			getPage, _ := template.ParseFiles("web/test.html")

			getPage.Execute(w, page)

		case "POST":
			fmt.Println("POST Received")
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			u := User{}
			u.Userask = r.FormValue("userask")
			u.Fname = r.FormValue("fname")
			u.Lname = r.FormValue("lname")

			postPage, _ := template.ParseFiles("web/test.html")
			err := postPage.Execute(w, u)
			fmt.Println(err)
		default:
			fmt.Fprintf(w, "Unable to get result.")
		}
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

func TestWeb() {
	fmt.Println("Listening on http://localhost:8080")
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":8080", nil)
}
