package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
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

type config struct {
	Port string `json:"port"`
}

type request struct {
	UserId 	string	`json:"user_id"`
	Date 	string	`json:"date"`
}

type response struct {
	Status int			`json:"status"`
	Result interface{}  `json:"result,omitempty"`
	Err    error		`json:"error,omitempty"`
}

type application struct {
    errlog *log.Logger
    infolog *log.Logger
}

func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()
    mux.HandleFunc("/create_event", app.create)
    mux.HandleFunc("/update_event", app.update)
    mux.HandleFunc("/delete_event", app.delete)
    mux.HandleFunc("/events_for_day", app.eventsD)
	mux.HandleFunc("/events_for_week", app.eventsW)
	mux.HandleFunc("/events_for_month", app.eventsM)
	return mux
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errlog.Output(2, trace)
 
	writeJson(w, response{Status: http.StatusServiceUnavailable, Err:err})
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

	var user request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// отправка запроса в бд на создание ивента
	work, err := imitateWork()
	if err != nil {
		app.serverError(w, err)
		return
	}

	writeJson(w, response{Status: http.StatusOK, Result: work})
}

func (app *application) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

	var user request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// отправка запроса в бд на обновление данных ивента
	work, err := imitateWork()
	if err != nil {
		app.serverError(w, err)
		return
	}

	writeJson(w, response{Status: http.StatusOK, Result: work})
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

	var user request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// отправка запроса в бд на удаление данных ивента
	work, err := imitateWork()
	if err != nil {
		app.serverError(w, err)
		return
	}

	writeJson(w, response{Status: http.StatusOK, Result: work})
}

func (app *application) eventsD(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        w.Header().Set("Allow", http.MethodGet)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

	r.ParseForm()

	var user request
	user.UserId = r.Form.Get("user_id")
	user.Date = r.Form.Get("date")
	if user.UserId == "" || user.Date == "" {
		app.clientError(w, http.StatusBadRequest)
	}

	// делается запрос в базу и получаем данные о ивентах на день и ошибку
	work, err := imitateWork() 
	if err != nil {
		app.serverError(w, err)
		return
	}

	writeJson(w, response{Status: http.StatusOK, Result: work})
}

func (app *application) eventsW(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        w.Header().Set("Allow", http.MethodGet)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

	r.ParseForm()

	var user request
	user.UserId = r.Form.Get("user_id")
	user.Date = r.Form.Get("date")
	if user.UserId == "" || user.Date == "" {
		app.clientError(w, http.StatusBadRequest)
	}

	// делается запрос в базу и получаем данные о ивентах на неделю и ошибку
	work, err := imitateWork()
	if err != nil {
		app.serverError(w, err)
		return
	}

	writeJson(w, response{Status: http.StatusOK, Result: work})
}

func (app *application) eventsM(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        w.Header().Set("Allow", http.MethodGet)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

	r.ParseForm()

	var user request
	user.UserId = r.Form.Get("user_id")
	user.Date = r.Form.Get("date")
	if user.UserId == "" || user.Date == "" {
		app.clientError(w, http.StatusBadRequest)
	}

	// делается запрос в базу и получаем данные о ивентах на месяц и ошибку
	work, err := imitateWork()
	if err != nil {
		app.serverError(w, err)
		return
	}

	writeJson(w, response{Status: http.StatusOK, Result: work})
}

func LoggerMiddleware(next http.Handler, log *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Вызов следующего обработчика в цепочке
		next.ServeHTTP(w, r)

		// Логирование информации о запросе
		log.Printf("Received request: %s %s %s", r.Method, r.URL, r.Proto)
	})
}

func writeJson(w http.ResponseWriter, v response) error {
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(v)
} 

func readConfig() (*config, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("read config error:", err)
		return nil, err
	}
	var cfg config

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("parse config error:", err)
		return nil, err
	}

	return &cfg, nil
}

func imitateWork() ([]string, error) {
	return nil, nil
}

func main() {
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	
	cfg, err := readConfig()
	if err != nil {
		errlog.Fatalf("read config error: %v", err)
	}
	app := application{
		errlog: errlog,
		infolog: infolog,
	}
	
	infolog.Println("Запуск веб-сервера на ", cfg.Port)
	srv := &http.Server{
        Addr: cfg.Port,
        ErrorLog: errlog,
        Handler: LoggerMiddleware(app.routes(), app.infolog),
    }
	err = srv.ListenAndServe()
    if err != nil {
        errlog.Fatal(err)
    }
	
}
