package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func TrimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func IdTrimer(s string) int {
	pathSegments := strings.Split(s, "/")

	appIndex := -1
	for i, segment := range pathSegments {
		if segment == "app" && i+1 < len(pathSegments) {
			appIndex = i + 1
			break
		}
	}

	if appIndex == -1 || appIndex >= len(pathSegments) {
		fmt.Println("Invalid URL structure")
		return 0
	}

	var result strings.Builder

	for _, char := range s {
		if char >= '0' && char <= '9' {
			result.WriteRune(char)
		}
	}

	s = result.String()

	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting ID to integer:", err)
		return 0
	}

	return id
}

func UrlBuilder(gameId int) string {
	urlExmpl := `https://store.steampowered.com/app/`
	url := urlExmpl + strconv.Itoa(gameId)
	return url
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
