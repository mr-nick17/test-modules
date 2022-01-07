package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	"learning/sessions/model"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

const sessionName = "session_data"

func IndexPageHandler() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/index.tmpl"))
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sd := getSessionData(session)
		page := model.LandingPage(sd)

		tmpl.Execute(w, page)
	}
}

func SetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")

		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Options.MaxAge = 0
		sd := getSessionData(session)
		session.Values["session"] = model.SessionData{
			Name:   name,
			Number: sd.Number,
		}
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func GetNameHandler() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/name.tmpl"))
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sd := getSessionData(session)
		page := model.NameData{
			Name: sd.Name,
		}

		tmpl.Execute(w, page)
	}
}

func SetNumberHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		number := r.FormValue("number")

		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Options.MaxAge = 0
		sd := getSessionData(session)
		session.Values["session"] = model.SessionData{Number: number, Name: sd.Name}
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func GetNumberHandler() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/number.tmpl"))
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sd := getSessionData(session)
		page := model.NumberData{
			Number: sd.Number,
		}

		tmpl.Execute(w, page)
	}
}

func getSessionData(session *sessions.Session) model.SessionData {
	var sd model.SessionData
	if val := session.Values["session"]; val != nil {
		sd = val.(model.SessionData)
	}
	return sd
}
