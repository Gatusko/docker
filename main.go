package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func EncodePlayReadyKID(hexKID string) (string, error) {
	b, err := hex.DecodeString(hexKID)
	if err != nil {
		return "", err
	}

	if len(b) != 16 {
		return "", errors.New("invalid KID byte size")
	}

	// Reverse the byte swapping done during decoding
	b[0], b[3] = b[3], b[0]
	b[1], b[2] = b[2], b[1]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	// bytes 8–15 stay the same

	return base64.StdEncoding.EncodeToString(b), nil
}
func main() {
	value, _ := EncodePlayReadyKID("0000000068f26d2071f0568fb0410c0b")
	fmt.Println("Hello World 2", value)
	test := "0000000062169ab559ee495df190d9ec"
	if len(test) != 32 {
		fmt.Println("test string length is:", len(test))
		return
	}
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
	http.ListenAndServe(":8080", nil)
}

type myServer struct {
}

func (m *myServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Print("server")
}
