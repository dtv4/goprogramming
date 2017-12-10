package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// set cookie to response
		expiration := time.Now().Add(1 * time.Hour)
		cookie := http.Cookie{Name: "user", Value: "duy", Expires: expiration}
		fmt.Printf("header before %+v\n", w.Header())
		http.SetCookie(w, &cookie)
		fmt.Printf("header after %+v\n", w.Header())
	})

	http.HandleFunc("/getcookie", func(w http.ResponseWriter, r *http.Request) {
		// get cookie from request
		cookie, _ := r.Cookie("user")
		fmt.Printf("cookie: %+v\n", cookie)
		fmt.Fprint(w, cookie)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
