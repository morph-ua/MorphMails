package main

import (
	"github.com/lib/pq"
	"time"
)

type Account struct {
	ID            string         `json:"id" gorm:"not null"`
	Emails        pq.StringArray `json:"emails" gorm:"index;type:text[]"`
	Forward       bool           `json:"forward" gorm:"default:true"`
	Paid          bool           `json:"paid" gorm:"default:false;not null"`
	Clients       pq.StringArray `json:"clients" gorm:"not null;type:text[];default:'{telegram}'"`
	TimesReceived int            `json:"timesReceived" gorm:"default:0;not null"`
}

type Letter struct {
	ID        string    `json:"id" gorm:"not null"`
	Html      string    `json:"html"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Client struct {
	ID     string `json:"id" gorm:"not null"`
	Name   string `json:"name" gorm:"not null"`
	URL    string `json:"url" gorm:"not null"`
	Secret string `json:"secret" gorm:"not null"`
}

type ClientDTO struct {
	ID     string `json:"id" query:"id" form:"id"`
	Name   string `json:"name" query:"name" form:"name"`
	URL    string `json:"url" query:"url" form:"url"`
	Secret string `json:"secret" query:"secret" form:"secret"`
}

type FinalResult struct {
	Message     string   `json:"message"`
	RenderedURI string   `json:"renderedURI"`
	ID          string   `json:"id"`
	Files       []string `json:"files"`
}

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CDNResponse struct {
	Uploaded bool   `json:"uploaded"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
}
