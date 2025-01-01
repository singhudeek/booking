package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/singhudeek/booking/internal/config"
	"github.com/singhudeek/booking/internal/handlers"
	"github.com/singhudeek/booking/internal/renders"

	"github.com/alexedwards/scs/v2"
)

const portnumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true if in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	renders.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting application on port %s\n", portnumber)

	// _ = http.ListenAndServe(portnumber, nil)

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
