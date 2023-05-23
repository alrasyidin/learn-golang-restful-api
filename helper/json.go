package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result any) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	HandleIfPanicError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	fmt.Printf("response json: %+v", response)
	err := encoder.Encode(response)
	HandleIfPanicError(err)
}
