package rockside

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rocksideio/rockside-sdk-go"
)

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

	rocksideAPIclient, err := rockside.NewClientFromAPIKey(os.Getenv("APIKEY"), rockside.Network(networks[0]), os.Getenv("BASE_URL"))
	if err != nil {
		replyError(w, http.StatusInternalServerError, "Error initialize")
		return
	}

	response, err := rocksideAPIclient.Transaction.Show(r.URL.Path[1:])
	if err != nil {
		replyError(w, http.StatusBadRequest, err.Error())
		return
	}

	b, _ := json.Marshal(response)
	fmt.Fprint(w, string(b))
}

func replyError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	errResp, _ := json.Marshal(ErrorResponse{Error: message})
	fmt.Fprint(w, string(errResp))
}
