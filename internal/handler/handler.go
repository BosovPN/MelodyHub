package handler

import (
	"encoding/json"
	"melodyhub/internal/service"
	"melodyhub/pkg/models"
	"net/http"
	"strconv"

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

func respondJSON(w http.ResponseWriter, status int, message string, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := make(map[string]interface{})

	response["message"] = message

	if data != nil {
		response["data"] = data
	}

	if err != nil {
		response["error"] = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}

// GetLibrary handles GET requests to retrieve all songs
func (h *Handler) GetLibrary(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err != nil {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err != nil {
			limit = l
		}
	}

	songs, err := h.service.GetSongs(filter, page, limit)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Failed to retrieve songs", nil, err)
		return
	}

	respondJSON(w, http.StatusOK, "Songs retrieved successfully", songs, nil)
}

// GetSong handles GET requests to retrieve a specific song by ID
func (h *Handler) GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	song, err := h.service.GetSong(idStr)
	if err != nil {
		respondJSON(w, http.StatusNotFound, "Song not found", nil, err)
		return
	}

	respondJSON(w, http.StatusOK, "Song retrieved successfully", song, nil)
}

// AddSong handles POST requests to add a new song
func (h *Handler) AddSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, "Bad requrest", nil, err)
		return
	}

	err = h.service.AddSong(&song)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Internal servier error", nil, err)
		return
	}

	respondJSON(w, http.StatusCreated, "Song added successfully", nil, nil)
}

// UpdateSong handles PUT requests to update a song by ID
func (h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	var song models.Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, "Bad requrest", nil, err)
		return
	}

	err = h.service.UpdateSong(idStr, &song)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Internal servier error", nil, err)
		return
	}

	respondJSON(w, http.StatusOK, "Song updated successfully", nil, nil)
}

// DeleteSong handles DELETE requests to delete a song by ID
func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	err := h.service.DeleteSong(idStr)
	if err != nil {
		respondJSON(w, http.StatusNotFound, "Not found", nil, err)
		return
	}

	respondJSON(w, http.StatusOK, "Song deleted successfully", nil, nil)
}
