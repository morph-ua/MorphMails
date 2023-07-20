package main

type (
	Result struct {
		ID          string   `json:"id,omitempty"`
		Message     Message  `json:"message"`
		RenderedURI string   `json:"rendered_uri,omitempty"`
		Files       []string `json:"files,omitempty"`
	}
	Error struct {
		Status  int    `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	}
	FileUploader struct {
		Uploaded bool   `json:"uploaded,omitempty"`
		Status   int    `json:"status,omitempty"`
		Message  string `json:"message,omitempty"`
	}
	Message struct {
		From string `json:"from,omitempty"`
		To   string `json:"to,omitempty"`
		Text string `json:"text,omitempty"`
	}
)
