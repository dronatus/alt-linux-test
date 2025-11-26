package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"simple-quiz/data"
)

var (
	tmpl *template.Template
)

// Init инициализирует обработчики (загружает шаблоны)
func Init() error {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		return fmt.Errorf("failed to load templates: %v", err)
	}
	return nil
}

// QuizHandler отображает страницу с тестом
func QuizHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	questions := data.GetQuestions()

	// Логируем для отладки
	fmt.Printf("Rendering quiz with %d questions\n", len(questions))
	for i, q := range questions {
		fmt.Printf("Question %d: ID=%d, Text=%s, Answers=%d\n", i, q.ID, q.Text, len(q.Answers))
	}

	if err := tmpl.ExecuteTemplate(w, "quiz.html", questions); err != nil {
		fmt.Printf("Template error: %v\n", err)
		http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
		return
	}
}

// ResultHandler обрабатывает результаты теста
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	character := calculateResult(r)
	if err := tmpl.ExecuteTemplate(w, "result.html", character); err != nil {
		http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
		return
	}
}

// calculateResult вычисляет результат на основе ответов
func calculateResult(r *http.Request) data.Character {
	questions := data.GetQuestions()
	characters := data.GetCharacters()

	characterCount := make(map[string]int)

	// Считаем голоса для каждого персонажа
	for _, q := range questions {
		value := r.FormValue(fmt.Sprintf("q%d", q.ID))
		if value != "" {
			characterCount[value]++
		}
	}

	// Находим персонажа с максимальным количеством голосов
	var maxCount int
	var resultCharacter string

	for char, count := range characterCount {
		if count > maxCount {
			maxCount = count
			resultCharacter = char
		}
	}

	// Если нет ответов или ничья - случайный персонаж
	if resultCharacter == "" || maxCount == 0 {
		return data.GetRandomCharacter()
	}

	return characters[resultCharacter]
}
