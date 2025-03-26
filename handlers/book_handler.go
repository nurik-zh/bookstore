package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bookstore/models"
	"github.com/gorilla/mux"
)

var books []models.Book
var bookIDCounter int

func GetBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Получаем параметр category_id
	categoryID, _ := strconv.Atoi(query.Get("category_id"))

	// Фильтрация книг по категории
	var filteredBooks []models.Book
	for _, book := range books {
		if categoryID > 0 && book.CategoryID != categoryID {
			continue
		}
		filteredBooks = append(filteredBooks, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredBooks)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	// 🔍 Валидация данных
	if book.Title == "" {
		http.Error(w, "Название книги не может быть пустым", http.StatusBadRequest)
		return
	}
	if book.AuthorID <= 0 {
		http.Error(w, "ID автора должен быть больше 0", http.StatusBadRequest)
		return
	}
	if book.CategoryID <= 0 {
		http.Error(w, "ID категории должен быть больше 0", http.StatusBadRequest)
		return
	}
	if book.Price <= 0 {
		http.Error(w, "Цена книги должна быть больше 0", http.StatusBadRequest)
		return
	}

	// ✅ Если всё в порядке, добавляем книгу
	bookIDCounter++
	book.ID = bookIDCounter
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			var updatedBook models.Book
			err := json.NewDecoder(r.Body).Decode(&updatedBook)
			if err != nil {
				http.Error(w, `{"error": "Ошибка в JSON"}`, http.StatusBadRequest)
				return
			}

			// 🔍 Валидация данных
			if updatedBook.Title == "" {
				http.Error(w, `{"error": "Название книги не может быть пустым"}`, http.StatusBadRequest)
				return
			}
			if updatedBook.AuthorID <= 0 {
				http.Error(w, `{"error": "ID автора должен быть больше 0"}`, http.StatusBadRequest)
				return
			}
			if updatedBook.CategoryID <= 0 {
				http.Error(w, `{"error": "ID категории должен быть больше 0"}`, http.StatusBadRequest)
				return
			}
			if updatedBook.Price <= 0 {
				http.Error(w, `{"error": "Цена книги должна быть больше 0"}`, http.StatusBadRequest)
				return
			}

			// ✅ Обновляем книгу (ID оставляем прежним)
			updatedBook.ID = id
			books[i] = updatedBook

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // ✅ Указали статус 200
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Книга не найдена"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
