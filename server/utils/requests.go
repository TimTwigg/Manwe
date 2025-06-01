package server_utils

import (
	"net/http"
	"strconv"
	"strings"
)

func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Max-Age", "86400")
}

func GetDetailLevel(r *http.Request) (int, error) {
	detail_level := r.URL.Query().Get("detail_level")
	var detail int = 1
	if detail_level != "" {
		d, err := strconv.Atoi(detail_level)
		if err != nil {
			return 0, err
		}
		detail = d
	}
	return detail, nil
}

func ErrorStatus(err error) int {
	if err != nil {
		if strings.HasPrefix(err.Error(), "ParseError") {
			return http.StatusInternalServerError
		} else {
			return http.StatusNotFound
		}
	}
	return http.StatusOK
}
