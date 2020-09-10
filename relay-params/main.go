package rockside

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rocksideio/rockside-sdk-go"
)

type ParamsRequest struct {
	Address string `json:"address"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Execute(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	networks, ok := r.URL.Query()["network"]

	if !ok || len(networks[0]) < 1 {
		replyError(w, http.StatusInternalServerError, "Url Param 'network' is missing")
		return
	}

	network := rockside.Network(networks[0])
	rocksideAPIclient, err := rockside.NewClientFromAPIKey(os.Getenv("APIKEY"), network, os.Getenv("BASE_URL"))
	if err != nil {
		replyError(w, http.StatusInternalServerError, "Error initialize")
		return
	}

	// Read the request body.
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		replyError(w, http.StatusInternalServerError, "Error reading request")
		return
	}

	// Parse the request body to get the topic name.
	paramsRequest := ParamsRequest{}
	if err := json.Unmarshal(data, &paramsRequest); err != nil {
		replyError(w, http.StatusInternalServerError, "Error parsing request")
		return
	}

	params, err := rocksideAPIclient.Relay.GetParams(paramsRequest.Address)
	if err != nil {
		replyError(w, http.StatusBadRequest, err.Error())
		return
	}

	b, _ := json.Marshal(params)
	fmt.Fprint(w, string(b))
}

func replyError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	errResp, _ := json.Marshal(ErrorResponse{Error: message})
	fmt.Fprint(w, string(errResp))
}
