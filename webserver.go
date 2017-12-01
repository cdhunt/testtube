package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var name = os.Getenv("COMPUTERNAME")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}

	path := os.Getenv("HEREIAM_DIR")
	if path == "" {
		path = "C:\\Users\\"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n%s", html.EscapeString(r.URL.Path), time.Now().Format("Jan 2 15:04:05"))
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n%s", name, time.Now().Format("Jan 2 15:04:05"))
	})

	http.HandleFunc("/echovar", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		text := "No key provided"
		if r.Form["var"] != nil {
			text = os.Getenv(r.Form["var"][0])
		}
		fmt.Fprintf(w, "%s\n%s", text, time.Now().Format("Jan 2 15:04:05"))
	})

	http.HandleFunc("/writetest", func(w http.ResponseWriter, r *http.Request) {

		f, err := ioutil.TempFile(path, "writetest")
		defer f.Close()
		if err != nil {
			fmt.Fprintf(w, "%s\n%s", err, time.Now().Format("Jan 2 15:04:05"))
			return
		}

		f.WriteString(time.Now().Format("Jan 2 15:04:05"))
		if err != nil {
			fmt.Fprintf(w, "%s\n%s", err, time.Now().Format("Jan 2 15:04:05"))
			return
		}

		f.Sync()

		fmt.Fprintf(w, "%s\n%s", f.Name(), time.Now().Format("Jan 2 15:04:05"))

	})

	http.HandleFunc("/volume", func(w http.ResponseWriter, r *http.Request) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			fmt.Fprintf(w, "%s\n", f.Name())
		}
	})

	fmt.Println("Serving content on port", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
