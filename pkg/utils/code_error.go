package utils

import "net/http"

func CodeError(error string) int {
	switch error {
	case "product not found":
		return http.StatusNotFound
	case "duplicate product":
		return http.StatusConflict
	case "product is exist":
		return http.StatusConflict
	case "duplicate user":
		return http.StatusConflict
	case "user not found":
		return http.StatusNotFound
	default:
		return http.StatusBadRequest
	}
}
