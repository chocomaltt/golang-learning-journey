package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CreateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	writeJSON(w, http.StatusOK, map[string]string{"id": id})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserReq
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		log.Println(err.Error())
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid body request"})
		return
	}

	if req.Name == "" || req.Email == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "name and email is required"})
		return
	}
	writeJSON(w, http.StatusOK, req)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func recoverMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal"})
			}
		}()
		next.ServeHTTP(w,r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("POST /users", createUser)

	handler := recoverMW(logging(mux))
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
