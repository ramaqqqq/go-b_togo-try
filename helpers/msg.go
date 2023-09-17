package helpers

import (
	"encoding/json"
	"net/http"
)

func MsgOk(status int, message string) map[string]interface{} {
	return map[string]interface{}{"code": status, "message": message}
}

func MsgErr(status int, message string, err string) map[string]interface{} {
	return map[string]interface{}{"code": status, "message": message, "error": err}
}

func Response(w http.ResponseWriter, status int, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
