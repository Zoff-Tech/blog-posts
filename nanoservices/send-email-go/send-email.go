
package main

import (
    "encoding/json"
    "net/http"
)

type EmailRequest struct {
    To      string `json:"to"`
    Message string `json:"message"`
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
    var req EmailRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil || req.To == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "status": "email sent",
        "to":     req.To,
    })
}

func main() {
    http.HandleFunc("/send-email", notifyHandler)
    http.ListenAndServe(":8080", nil)
}
