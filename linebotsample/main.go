package main

import (
	"net/http"

	"linebotsample"
)

func init() {
	http.Handle("/", linebot.NewServer())
}
