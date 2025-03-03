package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilmsg/skydash/handler"
)

func main() {

	indexHandler := &handler.IndexHandler{}
	userHandler := &handler.UserHandler{}

	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	r.HandleFunc("/users/register", userHandler.Register).Methods(http.MethodGet)
	r.HandleFunc("/users/register", userHandler.PostRegister).Methods(http.MethodPost)

	r.HandleFunc("/users/login", userHandler.Login).Methods(http.MethodGet)
	r.HandleFunc("/users/login", userHandler.PostLogin).Methods(http.MethodPost)

	r.HandleFunc("/pages/table", indexHandler.Table).Methods(http.MethodGet)
	r.HandleFunc("/pages/ui/button", indexHandler.UIButton).Methods(http.MethodGet)
	r.HandleFunc("/pages/ui/dropdowns", indexHandler.UIDropdown).Methods(http.MethodGet)
	r.HandleFunc("/pages/ui/typography", indexHandler.UITypography).Methods(http.MethodGet)
	r.HandleFunc("/pages/form/basic", indexHandler.UITypography).Methods(http.MethodGet)
	r.HandleFunc("/pages/chart/chart", indexHandler.Chart).Methods(http.MethodGet)
	r.HandleFunc("/pages/icon/mdi", indexHandler.IconMdi).Methods(http.MethodGet)
	r.HandleFunc("/", indexHandler.Index).Methods(http.MethodGet)

	// r.NotFoundHandler = http.HandlerFunc(indexHandler.Notfound)
	http.ListenAndServe(":3070", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
