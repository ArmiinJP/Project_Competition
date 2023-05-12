package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"net/mail"
)

type GravatarResponse struct {
	Ok          bool   `json:"ok"`
	GravatarUrl string `json:"gravatar_url"`
}

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func HandleGravatarRequest(w http.ResponseWriter, r *http.Request) {
 	request := strings.Split(r.RequestURI, "?email=")
	if len(request) < 2 ||  request[1] == ""{
		responseData, _ := json.Marshal(ErrorResponse{Ok: false, Message: "No email address provided"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, string(responseData))

		return
	}

	if valid(request[1]){
		requestLower := strings.ToLower(request[1])
		hashRequest := GetMD5Hash(requestLower)
		urlRequest := fmt.Sprintf("https://www.gravatar.com/avatar/%s", hashRequest)
		responseData, _ := json.Marshal(GravatarResponse{Ok: true, GravatarUrl: urlRequest})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(responseData))		
	} else {
		responseData, _ := json.Marshal(ErrorResponse{Ok: false, Message: "Invalid email address"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, string(responseData))		
	}
}
func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}
func valid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}
