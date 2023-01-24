package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var nottaken []string
	for char := 'a'; char <= 'z'; char++ {
		for i := 3; i <= 15; i++ {
			name := repeater(string(char), i)
			resp, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
			handleError(err)

			fmt.Println(name+" : ", resp.StatusCode)

			if resp.StatusCode == 204 {
				nottaken = append(nottaken, name)
			}
		}
	}

	content := strings.Join(nottaken, "\n")

	file, err := os.Create("output.txt")
	handleError(err)

	defer file.Close()

	_, err = file.WriteString(content)
	handleError(err)
}

func repeater(letter string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += letter
	}
	return repeated
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
