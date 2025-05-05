package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dhruvsaxena1998/rel/internal/parser"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type TranslateRequest struct {
	Expression string `json:"expression"`
}

type TranslateResponse struct {
	JSONLogic interface{} `json:"jsonLogic"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func translateHandler(w http.ResponseWriter, r *http.Request) {
	var req TranslateRequest

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request format"})
		return
	}

	// Create lexer and parser
	lexer := parser.NewLexer(req.Expression)
	p := parser.NewParser(lexer)

	// Parse expression
	expression := p.ParseExpression()
	if expression == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid expression: " + p.Errors()[0]})
		return
	}

	// Transform to JSONLogic
	jsonLogic, err := parser.Transform(expression)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Transform error: " + err.Error()})
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TranslateResponse{JSONLogic: jsonLogic})
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Post("/translate", translateHandler)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
