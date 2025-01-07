package controllers

import (
	"Backend-go/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetEmailsHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")
	query := "SELECT * FROM \"email_records\" WHERE match_all_raw_ignore_case('%" + search + "%')"

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	result, err := models.SearchEmails(query, page, limit)
	if err != nil {
		http.Error(w, "Error buscando correos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetEmailByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	email, err := models.SearchEmailByMessageID(id)
	if err != nil {
		http.Error(w, "Error obteniendo correo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(email)
}
