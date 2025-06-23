package server_utils

import (
	"net/http"
	"strconv"
	"strings"

	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	session "github.com/supertokens/supertokens-golang/recipe/session"
	usermetadata "github.com/supertokens/supertokens-golang/recipe/usermetadata"
)

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

func GetSessionUserID(r *http.Request) (string, error) {
	sessionContainer := session.GetSessionFromRequestContext(r.Context())
	if sessionContainer == nil {
		return "", error_utils.AuthError{Message: "No Session Found"}
	}
	userID := sessionContainer.GetUserID()
	if userID == "" {
		return "", error_utils.AuthError{Message: "No User ID Found"}
	}
	return userID, nil
}

func GetSessionUserEmail(userid string) (string, error) {
	metadata, err := usermetadata.GetUserMetadata(userid)
	if err != nil {
		return "", err
	}

	email, ok := metadata["email"]
	if !ok || email == "" {
		return "", error_utils.AuthError{Message: "No Email Found"}
	}
	return email.(string), nil
}
