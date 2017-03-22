package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"os"
)

func readFortuneFile(fortuneFile string) ([]string, error) {
	content, err := ioutil.ReadFile(fortuneFile)
	var fortunes []string = nil
	if err == nil {
		fortunes = strings.Split(string(content), "\n%\n")
	}
	return fortunes, err
}

func GetQuoteFrom(fortuneFile string) (string, error) {
	file := fortunePath + fortuneFile
	fortunes, err := readFortuneFile(file)
	if err == nil {
		i := getRandomInt(len(fortunes))
		quote := fortunes[i]
		return quote, nil
	}
	return "", err
}

func GetRandomQuote() (string, error) {
	files, err := GetFortuneFiles()
	if err != nil {
		return "", err
	}
	i := getRandomInt(len(files))
	return GetQuoteFrom(files[i].Name())
}

func getRandomInt(limit int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int() % limit
}

func GetFortuneFiles() ([]os.FileInfo, error) {
	return ioutil.ReadDir(fortunePath)
}

const fortunePath = "./fortunes/"
