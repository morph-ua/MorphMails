package main

type (
	FinalResult struct {
		ID          string   `json:"id"`
		Message     string   `json:"message"`
		RenderedURI string   `json:"renderedURI"`
		Files       []string `json:"files"`
	}
	HttpError struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	CDNResponse struct {
		Uploaded bool   `json:"uploaded"`
		Status   int    `json:"status"`
		Message  string `json:"message"`
	}
)
