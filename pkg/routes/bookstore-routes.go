package routes

import (
	"github.com/Joel1425/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")               // create a new book
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")                   // Get all the books
	router.HandleFunc("/book/{bookid}", controllers.GetBookByID).Methods("GET")       // Get a book by id
	router.HandleFunc("/book/{bookid}", controllers.PutBookByID).Methods("PUT")       // Put a book by id
	router.HandleFunc("/book/{bookid}", controllers.DeleteBookByID).Methods("DELETE") // Delete a book by id
}
