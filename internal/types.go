package main

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"time"
)

type (
	Receiver struct {
		ID     string `json:"id,omitempty"`
		Client Client `json:"client"`
	}
	Account struct {
		ID            uuid.UUID                     `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
		Receivers     datatypes.JSONSlice[Receiver] `json:"receivers" gorm:"index;not null;type:json"`
		Emails        pq.StringArray                `json:"emails" gorm:"index;type:text[]"`
		Forward       bool                          `json:"forward" gorm:"default:true"`
		Paid          bool                          `json:"paid" gorm:"default:false;not null"`
		TimesReceived int                           `json:"timesReceived" gorm:"default:0;not null"`
	}
	Letter struct {
		ID        string    `json:"id" gorm:"not null"`
		Html      string    `json:"html"`
		From      string    `json:"from"`
		To        string    `json:"to"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
	}
	Client struct {
		ID     string `json:"id" query:"id" form:"id" gorm:"not null"`
		Name   string `json:"name" query:"name" form:"name" gorm:"not null"`
		URL    string `json:"url" query:"url" form:"url" gorm:"not null"`
		Secret string `json:"secret" query:"secret" form:"secret" gorm:"not null"`
	}
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
