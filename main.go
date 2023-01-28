package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A slice of users to paginate
var users = []User{
	{ID: 1, Name: "User 1"},
	{ID: 2, Name: "User 2"},
	{ID: 3, Name: "User 3"},
	{ID: 4, Name: "User 4"},
	{ID: 5, Name: "User 5"},
	{ID: 6, Name: "User 6"},
	{ID: 7, Name: "User 7"},
	{ID: 8, Name: "User 8"},
	{ID: 9, Name: "User 9"},
	{ID: 10, Name: "User 10"},
	{ID: 11, Name: "User 11"},
	{ID: 12, Name: "User 12"},
	{ID: 13, Name: "User 13"},
	{ID: 14, Name: "User 14"},
	{ID: 15, Name: "User 15"},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		pageSize := r.URL.Query().Get("page_size")

		// Default values for page and page size
		if page == "" {
			page = "1"
		}
		if pageSize == "" {
			pageSize = "10"
		}

		// Parse the page and page size
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			http.Error(w, "Invalid page size", http.StatusBadRequest)
			return
		}

		// Calculate the start and end indices for the current page
		start := (pageInt - 1) * pageSizeInt
		end := start + pageSizeInt
		if end > len(users) {
			end = len(users)
		}

		// Create a slice of the users for the current page
		pageUsers := users[start:end]

		// Calculate the total number of pages
		totalPages := len(users) / pageSizeInt
		if len(users)%pageSizeInt != 0 {
			totalPages++
		}

		// Set the headers and write the response
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total-Pages", strconv.Itoa(totalPages))
		w.Header().Set("X-Total-Items", strconv.Itoa(len(users)))
		fmt.Fprintf(w, `{"page": %d, "page_size": %d, "data": %v}`, pageInt, pageSizeInt, pageUsers)
	})

	http.ListenAndServe(":8080", nil)
}
