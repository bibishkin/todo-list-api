package main

import (
	"net/http"
)

func HandlerTLS() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", def)
	//mux.HandleFunc("/user/login", login)
	mux.HandleFunc("/user/signup", signUp)

	return mux
}

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+addr+r.URL.Path, 307)
	})

	return mux
}

func def(w http.ResponseWriter, r *http.Request) {

	infoLog.Printf("path %v not accepted\n", r.URL.Path)
	http.Error(w, "404 not found", 404)
	return

}

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		infoLog.Printf("method %v not accepted", r.Method)
		http.Error(w, "400 bad request", 400)
		return
	}

	u, err := GetUser(r.Body)
	if err != nil {
		infoLog.Println(err)
		http.Error(w, "400 bad request", 400)
		return
	}

	exist, err := IsUserExist(u)
	if err != nil {
		infoLog.Println(err)
		http.Error(w, "500 internal server error", 500)
		return
	}

	if exist {
		infoLog.Println("the user with a name %v already exists")
		http.Error(w, "409 conflict", 409)
		return
	}

	err = AddUser(u)
	if err != nil {
		infoLog.Println(err)
		http.Error(w, "500 internal server error", 409)
		return
	}

	infoLog.Printf("user with name %v has registered\n", u.Name)
	http.Error(w, "201 created", 201)
	return
}
