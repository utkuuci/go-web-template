package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-----LOG-----")
		fmt.Println("Host: " + r.Host)
		fmt.Println("Method: " + r.Method)
		fmt.Println("Remote-Address: " + r.RemoteAddr)
		fmt.Println("Connection: " + r.Header.Get("Connection"))
		fmt.Println("User-Agent: " + r.Header.Get("User-Agent"))
		fmt.Println("X-User-Agent: " + r.Header.Get("X-User-Agent"))
		fmt.Println("X-Device-User-Agent: " + r.Header.Get("X-Device-User-Agent"))
		fmt.Println("Proto: " + r.Proto)
		fmt.Println("Cookie: " + r.Header.Get("Cookie"))
		fmt.Println("Requested the path: " + r.URL.Path)
		fmt.Println("-----LOG-----")

		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
