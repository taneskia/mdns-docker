package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	for {
		log.Default().Println("Pulling from " + "http://pop-os.local")

		response, err := http.Get("http://pop-os.local")

		if err != nil {
			log.Default().Fatalln("Could not execute http request.")
			continue
		}

		log.Default().Println("Successfully executed http request.")

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Default().Fatalln("Could not read http response body.")
			continue
		}

		log.Default().Println(string(responseData))

		log.Default().Print("Successfully read http response.\n\n")

		time.Sleep(5 * time.Second)
	}
}
