package main

import (
	"github.com/alexedwards/scs/v2"
	"go-web/pkg/config"
	"go-web/pkg/handlers"
	"go-web/pkg/render"
	"log"
	"net/http"
	"time"
)

const PORT = ":4000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this true when app in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
