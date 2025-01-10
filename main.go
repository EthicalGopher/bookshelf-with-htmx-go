package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}


func BookShelf() []string {
	file, err := os.Open("books.txt")
	checkError(err)
	defer file.Close()

	arr, err := io.ReadAll(file)
	checkError(err)

	lines := strings.Split(string(arr), "\n")
	var Books []string

	for _, line := range lines {
		if line != "" {
			Books = append(Books, line)
		}
	}

	return Books
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})


	app.Get("/search", func(c *fiber.Ctx) error {
		query := c.Query("search", "")
		var result []string
		Books := BookShelf()

		for _, book := range Books {
			if strings.Contains(strings.ToLower(book), strings.ToLower(query)) {
				result = append(result, book)
			}
		}

		if len(result) == 0 {
			return c.SendString("Sorry, can't find the book you mentioned (" + query + ")")
		}

		html := "<ul>"
		for i := 0; i < len(result) && i < 5; i++ {
			html += "<li>" + result[i] + "</li>"
		}
		html += "</ul>"
		if query!="" && query!=" "{
			return c.SendString(html)
		}
		return c.SendString("Please enter the query")
	})



	app.Listen(":8000")

}
