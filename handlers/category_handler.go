package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var categories []models.Category
var categoryIDCounter int

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	// 🔍 Валидация данных
	if category.Name == "" {
		http.Error(w, "Название категории не может быть пустым", http.StatusBadRequest)
		return
	}

	// ✅ Если всё в порядке, добавляем категорию
	categoryIDCounter++
	category.ID = categoryIDCounter
	categories = append(categories, category)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
