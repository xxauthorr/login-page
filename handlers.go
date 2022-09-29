package main

import (
	"fmt"
	"net/http"
)

func auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		if session.Values["autheticated"] == false || session.Values["authenticated"] == nil {
			val := Credentials{ErrMsg: "You Must Login !"}
			tpl.ExecuteTemplate(w, "loginPage.html", val)
			return
		} else {

			HandlerFunc.ServeHTTP(w, r)
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Check weather the user is already loggedin
	if session.Values["authenticated"] == true {

		fmt.Println("Home Page")
		//redirect to "/home"
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		// First clear cache
		w.Header().Set("Cache-Control", "no-store")

		//shows the login template
		val := Credentials{Heading: "Login Page"}
		fmt.Println("Login Page")
		tpl.ExecuteTemplate(w, "loginPage.html", val)
	}
}

func LoginCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("Parse error! (Login checker)")
		fmt.Fprintln(w, "Parse error :", err)
		return
	}
	session, _ := store.Get(r, "session")
	//user send values are asigned to variables
	Email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	if Email == "test@gmail.com" && password == "12345" {
		// Assigning values to the session variables authenticated and emailId
		session.Values["authenticated"] = true
		session.Values["emailId"] = Email
		// Saving the session values
		session.Save(r, w)
		//Show the homepage
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else if Email == "test@gmail.com" {
		fmt.Println("Wrong Password")
		// if password and username is wrong, we assign false to authenticated
		session.Values["authenticated"] = false
		// saving the session variables
		session.Save(r, w)
		//show the login page with the errMsg
		val := Credentials{Heading: "Login page", ErrMsg: "Invalid password"}
		tpl.ExecuteTemplate(w, "loginPage.html", val)
	} else {
		fmt.Println("Invalid Login")

		// Clear cache
		w.Header().Set("Cache-Control", "no-store")

		session.Values["authenticated"] = false
		session.Save(r, w)
		val := Credentials{Heading: "Login page", ErrMsg: "Invalid Login"}
		tpl.ExecuteTemplate(w, "loginPage.html", val)
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	// if session.Values["authenticated"] == true {
	fmt.Println("LoggedIn")
	email := session.Values["emailId"].(string)

	val := Credentials{Heading: "Home Page", Email: email}
	// Clear the cache
	w.Header().Set("Cache-Control", "no-store")

	tpl.ExecuteTemplate(w, "homePage.html", val)
	// } else {
	// 	// Clear the cache
	// 	w.Header().Set("Cache-Control", "no-store")

	// 	//redirect to login page
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// }
}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Log Out")

	// Clearing the cache memory
	w.Header().Set("Cache-Control", "no-store")

	// Getting the session
	session, _ := store.Get(r, "session")

	// Destroying the sessions by assigning the maxAge value to -1
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1

	// Saving the session
	session.Save(r, w)

	// Redirect to the login page
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func NoPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "errorPage.html", nil)
}
