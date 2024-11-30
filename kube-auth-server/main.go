package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TempStruct struct {
}

func main() {

	http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var tr TempStruct

		err := decoder.Decode(&tr)
		if err != nil {
			log.Printf("[Error]  \n")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code": http.StatusBadRequest,
			})
			return
		}
		//TODO......
	})

	log.Println(http.ListenAndServe(":3000", nil))
}
