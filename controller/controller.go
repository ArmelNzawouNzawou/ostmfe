package controller

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin"
	"ostmfe/controller/collection"
	"ostmfe/controller/event"
	"ostmfe/controller/history"
	"ostmfe/controller/home"
	"ostmfe/controller/people"
	"ostmfe/controller/place"
	"ostmfe/controller/user"
	"ostmfe/controller/visit"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", homeHanler(env))
	mux.Mount("/home", home.Home(env))
	mux.Mount("/visit", visit.Home(env))
	mux.Mount("/history", history.Home(env))
	mux.Mount("/collection", collection.Home(env))
	mux.Mount("/place", place.Home(env))
	mux.Mount("/people", people.Home(env))
	mux.Mount("/admin_user", admin.Home(env))
	mux.Mount("/user", user.Home(env))
	mux.Mount("/event", event.Home(env))

	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "index.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
