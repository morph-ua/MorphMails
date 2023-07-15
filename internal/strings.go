package main

import (
	"net/http"
)

var messageTemplate = `From: %s
To: %s
Subject: %s

%s`

var internalServerErrorMessage = &HttpError{
	Status:  http.StatusInternalServerError,
	Message: "Internal server error",
}

var notFoundMessage = &HttpError{
	Status:  http.StatusNotFound,
	Message: "The user with the specified ID was not found in the API database. Please register with the bot",
}

var badRequestMessage = &HttpError{
	Status:  http.StatusBadRequest,
	Message: "Bad request. Couldn't parse an ID",
}
