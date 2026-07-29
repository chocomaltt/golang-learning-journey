package main

import (
	"log"
	"net/http"
	"sync"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID string `json:"id"`
	ProductName string `json:"productName"`
	Amount int16 `json:"amount"`
}

type ProductStore struct{
	mu sync.RWMutex
	products map[string]Product
}

func NewProductStore() *ProductStore{
	return &ProductStore{
		products: make(map[string]Product),
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (s *ProductStore) addProduct(w http.ResponseWriter, r *http.Request) {
	var req Product
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.ProductName == "" || req.Amount < 0 {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error":"invalid body request"})
		return
	}
	
	id := uuid.New()
	req.ID = id.String()

	s.mu.Lock()
	defer s.mu.Unlock()

	s.products[req.ID] = req

	writeJSON(w, http.StatusCreated, req)
}

func (s *ProductStore) getById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	s.mu.RLock()
	defer s.mu.RUnlock()
	returns, exists := s.products[id]

	if !exists {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
		return
	}
	writeJSON(w, http.StatusOK, returns)
}

func (s *ProductStore) getAllProduct(w http.ResponseWriter, r *http.Request){
	s.mu.RLock()
	defer s.mu.RUnlock()

	if exists := s.products; len(exists) == 0{
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "data is empty"})
		return
	}

	for _,v := range s.products{
		writeJSON(w, http.StatusOK, v)
	}
}

func (s *ProductStore) deleteById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	s.mu.Lock()
	defer s.mu.Unlock()

	if _,exists := s.products[id]; !exists{
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "product not found"})
		return
	}
	delete(s.products, id)
	writeJSON(w, http.StatusOK, map[string]string{"success": "product has been deleted"})
}

func logging(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		next.ServeHTTP(w,r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func recoverMW(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func(){
			if rec := recover(); rec != nil{
				writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal Server Error"})
			}
		}()
		next.ServeHTTP(w,r)
	})
}

func main(){
	mux := http.NewServeMux()
	store := NewProductStore()
	mux.HandleFunc("GET /product", store.getAllProduct)
	mux.HandleFunc("GET /product/{id}", store.getById)
	mux.HandleFunc("POST /product/add", store.addProduct)
	mux.HandleFunc("POST /product/delete/{id}", store.deleteById)

	handler := recoverMW(logging(mux))
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}