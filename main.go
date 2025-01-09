package main

import (
	
	"io"
	"os"
	"strings"
	
	"github.com/gofiber/fiber/v2"
)
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func BookShelf() []string {
	file,err:=os.Open("books.txt")
	checkError(err)
	defer file.Close()
	arr,err:=io.ReadAll(file)
	checkError(err)
	var Books []string
	lines := strings.Split(string(arr), "\n")
	
	Books=append(Books, lines...)
	return Books
}
func main() {
	app:=fiber.New()
	app.Get("/",func(c * fiber.Ctx)error{
		return c.SendFile("./index.html")
	})
	app.Get("/search",func(c *fiber.Ctx)error{
		query:=c.Query("search"," ")
		if query == "" {
			return c.SendString("Please enter the name of the book to search for it")
		}
		var result []string
		Books:=BookShelf()
		for _,j:= range Books {
			if strings.Contains(strings.ToLower(j), strings.ToLower(query)){
				result=append(result, j)
			}
		}

		if len(result)==0 {
			return c.SendString("Sorry can't find the book you mentioned ("+query+")")
		}

		html:="<ul>"
		for _,j:=range result{
			html+="<li>"+j+"</li>"
		}
		html+="</ul>"
		if query != ""{

			return c.SendString(html)
		}
		return c.SendString("Please enter the name of the book to search for it")
	})

	app.Listen(":8000")
	
}