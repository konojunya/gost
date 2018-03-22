package server

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"github.com/konojunya/gost/auth"
	"github.com/skratchdot/open-golang/open"
)

var (
	closeCh = make(chan bool, 1)
	addr    = "127.0.0.1:6578"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	config := auth.GetOAuthConfig()
	url := config.AuthCodeURL(state)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	config := auth.GetOAuthConfig()
	token, err := config.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
	if err != nil {
		log.Fatal(err)
	}
	if !token.Valid() {
		log.Fatal(fmt.Errorf("invalid token"))
	}

	err = auth.CreateTokenFile(token)
	if err != nil {
		log.Fatal(err)
	}

	closeCh <- true

	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Listen ListenAndServe
func Listen() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/oauth", callbackHandler)

	go func() {
		open.Run("http://" + addr)
		fmt.Println("listen and serve http://" + addr + "\n")
		if err := http.Serve(listener, nil); err != nil {
			log.Fatal(err)
		}
	}()

	if <-closeCh {
		time.Sleep(time.Second * 3)
		fmt.Println("You have been authenticated.")
		fmt.Println("Please enjoy gost!")
		if err := listener.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
