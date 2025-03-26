package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var authors []models.Author
var authorIDCounter int

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	// 🔍 Валидация данных
	if author.Name == "" {
		http.Error(w, "Имя автора не может быть пустым", http.StatusBadRequest)
		return
	}

	// ✅ Если всё в порядке, добавляем автора
	authorIDCounter++
	author.ID = authorIDCounter
	authors = append(authors, author)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}
