package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type MathRequest struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type MathResponse struct {
	Result float64 `json:"result"`
	Status string  `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
		Status:  "success",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
		return
	}

	var req MathRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON"})
		return
	}

	result := req.Num1 + req.Num2
	response := MathResponse{
		Result: result,
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
		return
	}

	var req MathRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON"})
		return
	}

	result := req.Num1 - req.Num2
	response := MathResponse{
		Result: result,
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/add", addHandler)
	http.HandleFunc("/api/subtract", subtractHandler)
	
	fmt.Println("Server starting on port 8080...")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /          - Calculator UI")
	fmt.Println("  GET  /api/hello - Hello API")
	fmt.Println("  POST /api/add   - Addition service")
	fmt.Println("  POST /api/subtract - Subtraction service")
	log.Fatal(http.ListenAndServe(":8080", nil))
}