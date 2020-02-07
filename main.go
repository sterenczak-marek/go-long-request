package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

func main() {
	n := http.NewServeMux()
	n.HandleFunc("/", indexHandler)

	PORT, declared := os.LookupEnv("PORT")
	if !declared {
		PORT = "8000"
	}
	log.Println("Listening for connections on port: ", PORT)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, n))
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()

	var result []map[string]interface{}
	var testData = map[string]interface{}{
		"id": "Test",
		"some_obj": map[string]interface{}{
			"id":   "test",
			"some": 1,
		},
	}

	for i := 0; i < int(math.Pow10(5))*3; i++ {
		result = append(result, testData)
	}

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	time.Sleep(25 * time.Second)

	json.NewEncoder(rw).Encode(result)

	log.Printf("Response time: %s", time.Since(start))
}
