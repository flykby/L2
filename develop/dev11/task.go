package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type Event struct {
	ID      int       `json:"event_id"`
	Date    time.Time `json:"date"`
	UserId  int       `json:"user_id"`
	Content string    `json:"content"`
}

type Storage struct {
	Storage []Event
	mx      *sync.Mutex
}

// New создает календарь.
func NewStorage() *Storage {
	return &Storage{
		Storage: make([]Event, 0, 10),
		mx:      &sync.Mutex{},
	}
}

// Добавляет событие в календарь.
func (s *Storage) CreateEvent(event Event) {
	s.mx.Lock()
	s.Storage = append(s.Storage, event)
	s.mx.Unlock()
}

// Удаляет событие из календаря. 
// Возвращает 3 значения: 
// -1 - когда запись не была найдена
//  0 - когда запись была найдена, но пользователь не может ее удалить
//  1 - когда запись была успешно удалена
func (s *Storage) DeleteEvent(eventID int, userId int) int {
	statusCode := -1

	s.mx.Lock()
	for i := 0; i < len(s.Storage); i++ {
		if s.Storage[i].ID == eventID {
			if s.Storage[i].UserId == userId {
				s.Storage = append(s.Storage[:i], s.Storage[i+1:]...)
				statusCode = 1
			} else {
				statusCode = 0
			}
		}
	}

	s.mx.Unlock()

	return statusCode
}

// Обновляет описание события.
// Возвращает 3 значения: 
// -1 - когда запись не была найдена
//  0 - когда запись была найдена, но пользователь не может ее удалить
//  1 - когда запись была успешно удалена
func (s *Storage) UpdateEvent(eventId int, userId int, content string) int {
	statusCode := -1

	s.mx.Lock()

	for i := 0; i < len(s.Storage); i++ {
		if s.Storage[i].ID == eventId {
			if s.Storage[i].UserId == userId {
				s.Storage[i].Content = content
				statusCode = 1
			} else {
				statusCode = 0
			}
		}
	}

	s.mx.Unlock()

	return statusCode
}

// Возвращает слайс событий по дню.
func (s *Storage) ByDay(date time.Time) []Event {
	s.mx.Lock()

	var events []Event

	for _, e := range s.Storage {
		if e.Date.Year() == date.Year() && e.Date.Month() == date.Month() && e.Date.Day() == date.Day() {
			events = append(events, e)
		}
	}

	s.mx.Unlock()

	return events
}

func (s *Storage) ByWeek(date time.Time) []Event {
	s.mx.Lock()

	var events []Event

	for _, e := range s.Storage {
		difference := date.Sub(e.Date)
		if difference < 0 {
			difference = -difference
		}

		if difference <= time.Duration(7*24)*time.Hour {
			events = append(events, e)
		}
	}

	s.mx.Unlock()

	return events
}

// Возвращает слайс событий по месяцу.
func (s *Storage) ByMonth(date time.Time) []Event {
	s.mx.Lock()

	var events []Event

	for _, e := range s.Storage {
		if e.Date.Year() == date.Year() && e.Date.Month() == date.Month() {
			events = append(events, e)
		}
	}

	s.mx.Unlock()

	return events
}

var storage *Storage


// Собирает в структуру полученные из POST запроса
func decodeParams(r *http.Request) (Event, error) {
	// Парсит и форматирует под нужный тип.
	eventID, errID := strconv.Atoi(r.FormValue("event_id"))
	if errID != nil {
		return Event{}, errors.New(fmt.Sprintf("Format is wrong: %s", errID))
	}

	date, errDate := time.Parse("2006-01-02", r.FormValue("date"))
	if errDate != nil {
		return Event{}, errors.New(fmt.Sprintf("Format is wrong: %s", errDate))
	}

	userId, errDate := strconv.Atoi(r.FormValue("user_id"))
	if errDate != nil {
		return Event{}, errors.New(fmt.Sprintf("Format is wrong: %s", errDate))
	}

	content := r.FormValue("content")
	if content == "" {
		return Event{}, errors.New("Format is wrong: context is void")
	}

	return Event{
		ID:      eventID,
		Date:    date,
		UserId: userId,
		Content: content,
	}, nil
}

// Валидация события.
func validateEvent(event Event) bool {
	return event.ID > 0 && event.UserId > 0
}


// Обрабатывает запрос на создание события.
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	// Если запрос не POST.
	if r.Method != http.MethodPost {
		return
	}

	event, err := decodeParams(r)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)

		return
	}

	if validateEvent(event) {
		storage.CreateEvent(event)
		resultResponse(w, "Event was created!", storage.Storage, 200)
	} else {
		errorResponse(w, "Invalid value!", http.StatusBadRequest)
	}

	log.Println("CreateEventHandler")
}

// Обрабатывает запрос на получение событий за на этот день.
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получает результат.
	result := storage.ByDay(date)
	resultResponse(w, "good!", result, http.StatusOK)

	log.Println("events_for_day")
}

// Обрабатывает запрос на получение событий за на эту неделю.
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получает результат.
	result := storage.ByWeek(date)
	resultResponse(w, "good!", result, http.StatusOK)

	log.Println("events_for_week")
}

// Обрабатывает запрос на получение событий за на этот месяц.
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Если это не GET-request.
	if r.Method != http.MethodGet {
		return
	}

	// Парсит дату.
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получает результат.
	result := storage.ByMonth(date)
	resultResponse(w, "good!", result, http.StatusOK)

	log.Println("events_for_month")
}

// Обрабатывает запрос на обновление события.
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	eventId, err := strconv.Atoi(r.FormValue("event_id"))
	if err != nil {
		errorResponse(w, "Invalid value!", http.StatusBadRequest)
	}

	userId, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		errorResponse(w, "Invalid value!", http.StatusBadRequest)
	}

	description := r.FormValue("content")


	statusCode := storage.UpdateEvent(eventId, userId, description)

	if statusCode == 1 {
		resultResponse(w, "Event was updated!", storage.Storage, http.StatusOK)
	} else if statusCode == 0 {
		errorResponse(w, "User can't update this Event!", http.StatusServiceUnavailable)
	} else {
		errorResponse(w, "There is no such Event!", http.StatusBadRequest)
	}

	log.Println("UpdateEventHandler")
}

// Обрабатывает запрос на удаление события.
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	eventId, err := strconv.Atoi(r.FormValue("event_id"))
	userId, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		errorResponse(w, "Invalid value!", http.StatusBadRequest)
	}
	statusCode := storage.DeleteEvent(eventId, userId)

	if statusCode == 1 {
		resultResponse(w, "Event was deleted!", storage.Storage, http.StatusOK)
	} else if statusCode == 0 {
		errorResponse(w, "User can't delete this Event!", http.StatusServiceUnavailable)
	} else {
		errorResponse(w, "There is no such Event!", http.StatusBadRequest)
	}
}



// Отправляет результат в ответ на запрос.
func resultResponse(w http.ResponseWriter, result string, events []Event, status int) {
	resultResponse := struct {
		Result string       `json:"result"`
		Events []Event		`json:"events"`
	}{Result: result, Events: events}

	responseJSON, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

// Отправляет результат-ошибку в ответ на запрос.
func errorResponse(w http.ResponseWriter, errDesc string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	} {Error: errDesc}

	// Сериализация.
	responseJSON, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTTP-response.
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func main() {
	// POST
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)

	// GET
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	// Storage с event-ами
	storage = NewStorage()

	// Запускает сервер.
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err == nil {
		log.Fatal("Server error:", err)
	}
}
