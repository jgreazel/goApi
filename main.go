package main

import (
	// "encoding/json"
	"encoding/json"
	"log"

	"github.com/gorilla/mux"

	// "math/rand"
	"net/http"
	// "strconv"
)

// Book struct (model)
// note: structs are like classes, *points to another struct, go is statically typed
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init books var as a slice Book struct
// note: slices are like lists in java, declare them like []<type>
var books []Book

// get all books
// note: api methods utilize net/http
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	// loop through books and find correct id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update existing book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init router
	router := mux.NewRouter()

	//mock data
	//note: &<type> similar to C pointers
	books = append(books, Book{
		ID: "1", Isbn: "52634", Title: "Harry Potter", Author: &Author{
			Firstname: "J.K.", Lastname: "Rowling"}})
	books = append(books, Book{
		ID: "2", Isbn: "98794", Title: "Bible", Author: &Author{
			Firstname: "Jesus", Lastname: "Christ"}})
	books = append(books, Book{
		ID: "3", Isbn: "35400", Title: "Girl Wash Your Face", Author: &Author{
			Firstname: "Dumb", Lastname: "B."}})

	// route handlers / endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
