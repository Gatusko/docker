package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("mis argumentos", os.Args)
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	fmt.Println("hello world")
	//	h := myServer{}
	http.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		// vamos acrear un archivo random
		myRandom := rand.Intn(1000)
		fileName := fmt.Sprintf("/data/%d.txt", myRandom)
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		f.WriteString("Hello World")
	})
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Print("server")
	})
	http.HandleFunc("GET /read", func(w http.ResponseWriter, r *http.Request) {
		result, err := os.ReadFile("/data/read.txt")
		if err != nil {
			log.Print("ERROR", err)
			return
		}
		log.Print("server", result)
	})
	http.ListenAndServe(":8080", nil)
}

type myServer struct {
}

func (m *myServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Print("server")
}
