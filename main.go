package main

import (
	"flag"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"ostmfe/config"
	"ostmfe/controller"
	"time"
)

var sessionManager *scs.SessionManager

func Environment() *config.Env {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	env := &config.Env{
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
		InfoLog:  log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		Path:     "./view/html/",
		Session:  sessionManager,
	}
	return env
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: Environment().ErrorLog,
		Handler:  controller.Controllers(Environment()),
	}

	Environment().InfoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	error := srv.ListenAndServe()
	Environment().ErrorLog.Fatal(error)

}
