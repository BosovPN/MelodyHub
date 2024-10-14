package handler

import (
	"encoding/json"
	"melodyhub/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func RegisterRoutes(r *mux.Router, svc *service.Service) {
	h := NewHandler(svc)

	r.HandleFunc("/library", h.GetLibrary).Methods("GET")
	r.HandleFunc("/library/{id}", h.GetSong).Methods("GET")
	r.HandleFunc("/library", h.AddSong).Methods("POST")
	r.HandleFunc("/library/{id}", h.UpdateSong).Methods("PUT")
	r.HandleFunc("/library/{id}", h.DeleteSong).Methods("DELETE")
}

// GetLibrary handles GET requests to retrieve all songs
func (h *Handler) GetLibrary(w http.ResponseWriter, r *http.Request) {
	// Stub: Respond with an empty JSON array
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]string{}) // Placeholder response
}

// GetSong handles GET requests to retrieve a specific song by ID
func (h *Handler) GetSong(w http.ResponseWriter, r *http.Request) {
	// Stub: Respond with a placeholder song
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"id":          "1",
		"group":       "Sample Group",
		"song":        "Sample Song",
		"releaseDate": "2024-10-14",
		"text":        "Sample lyrics...",
		"link":        "http://example.com/sample-song",
	})
}

// AddSong handles POST requests to add a new song
func (h *Handler) AddSong(w http.ResponseWriter, r *http.Request) {
	// Stub: Respond with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Song added successfully",
	})
}

// UpdateSong handles PUT requests to update a song by ID
func (h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	// Stub: Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Song updated successfully",
	})
}

// DeleteSong handles DELETE requests to delete a song by ID
func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	// Stub: Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Song deleted successfully",
	})
}
