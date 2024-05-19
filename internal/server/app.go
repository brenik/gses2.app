package server

import (
	"encoding/json"
	"fmt"
	"github.com/brenik/gses2.app/internal/service"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

func getRateHandler(w http.ResponseWriter, r *http.Request) {
	rate, err := service.GetRate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postSubscriberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := service.NewRepository(db)

	exist, err := repo.IsExistEmail(email)
	if err != nil {
		log.Fatal(err)
	}

	if exist {
		w.WriteHeader(http.StatusConflict) // 409 Conflict
		fmt.Fprintf(w, "Email %s already exists\n", email)
	} else {
		err = repo.AddSubscriber(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK) // 200 OK
		fmt.Fprintf(w, "Your email %s was saved\n", email)
	}

}

func Run(port string) {

	r := mux.NewRouter()
	//For example, you can use CLI: curl http://localhost:8000/api/rate
	r.HandleFunc("/api/rate", getRateHandler).Methods("GET")

	//For example, you can use CLI: curl -X POST -d "email=test@example.com" http://localhost:8000/api/subscribe
	r.HandleFunc("/api/subscribe", postSubscriberHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(port, r))
}
