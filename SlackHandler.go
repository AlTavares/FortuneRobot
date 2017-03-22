package main

import (
	"log"
	"net/http"

	"os"
	"strings"

	"github.com/labstack/echo"
)

type SlackHandler struct{}

func (s SlackHandler) Webhook(c echo.Context) error {
	text := strings.TrimSpace(c.FormValue("text"))
	var err error
	response := &Response{
		Type: "in_channel",
		Text: "No fortune for you",
	}
	defer c.JSON(http.StatusOK, response)

	if text == "help" {
		var files []os.FileInfo
		files, err = GetFortuneFiles()
		var quote string
		for _, file := range files {
			quote += file.Name() + "\n"
		}
		response.Text = quote
		return nil
	}
	if text == "" {
		response.Text, err = GetRandomQuote()
	} else {
		response.Text, err = GetQuoteFrom(text)
	}
	if err != nil {
		log.Println(err)
	}
	return err
}

type Response struct {
	Type string `json:"response_type"`
	Text string `json:"text"`
}
