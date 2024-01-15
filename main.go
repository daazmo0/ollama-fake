package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Version struct {
    Version string `json:"version"`
}

func main() {
    http.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        version := Version{Version: "0.1.20"}
        json.NewEncoder(w).Encode(version)
    })

    fmt.Println("Server is listening on port 11434")
    http.ListenAndServe(":11434", nil)
}
