package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := r.RemoteAddr
		method := r.Method
		url := r.URL.Path

		wr := &responseWriter{ResponseWriter: w, statusCode: 200}

		next.ServeHTTP(wr, r)

		log.Printf("[%s] %s %s %d %v",
			ip,
			method,
			url,
			wr.statusCode,
			time.Since(start),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {
	fs := http.FileServer(http.Dir("./public/"))
	http.Handle("/", loggingMiddleware(fs))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			log.Println("Login attempt:", r.FormValue("email"))
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		http.ServeFile(w, r, "./public/index.html")
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			log.Println("Signup attempt:", r.FormValue("name"))
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.ServeFile(w, r, "./public/signup.html")
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			log.Println("Reset password:", r.FormValue("email"))
			http.Redirect(w, r, "/reset-code", http.StatusSeeOther)
			return
		}
		http.ServeFile(w, r, "./public/reset.html")
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Главный экран (заглушка)</h1><p>Здесь будет контент после входа.</p>"))
	})

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
