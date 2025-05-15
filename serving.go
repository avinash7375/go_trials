package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}

func main() {
    http.HandleFunc("/users", handleUsers)
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        json.NewEncoder(w).Encode(users)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}