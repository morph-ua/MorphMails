package main

type (
	result struct {
		ID          string   `json:"id,omitempty"`
		Message     message  `json:"message"`
		RenderedURI string   `json:"rendered_uri,omitempty"`
		Files       []string `json:"files,omitempty"`
	}
	message struct {
		From    string `json:"from,omitempty"`
		To      string `json:"to,omitempty"`
		Text    string `json:"text,omitempty"`
		Subject string `json:"Subject,omitempty"`
	}
	unwrappedDefaults struct {
		Recipients []string `json:"recipients,omitempty"`
		From       string   `json:"from,omitempty"`
		Subject    string   `json:"subject,omitempty"`
		HTML       string   `json:"html,omitempty"`
		Text       string   `json:"text,omitempty"`
	}
)

const (
	fieldID    = "id"
	fieldConID = "connectorID"
)
