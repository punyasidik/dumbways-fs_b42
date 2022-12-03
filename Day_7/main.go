package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	//route untuk menginisialisai folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/addblog", addBlog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")

	fmt.Println("Server berjalan pada port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/add-blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// object golang
	data := map[string]interface{}{
		"Title":   "Pasar Coding Dari Dumbways",
		"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan.",
		"Id":      id,
	}

	tmpt.Execute(w, data)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/blog.html")

	// if condition
	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}
