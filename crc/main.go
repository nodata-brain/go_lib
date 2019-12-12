package crc

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Resp string `json:"response_token"`
}

func Crc(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("crc_token")
	m := hmac.New(sha256.New, []byte(os.Getenv("Sec")))
	m.Write([]byte(q))

	w.Header().Set("Content-Type", "application/json")

	resp := Response{Resp: "sha256=" + base64.StdEncoding.EncodeToString(m.Sum(nil))}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
	}

	fmt.Fprintf(w, string(jsonBytes))
}
