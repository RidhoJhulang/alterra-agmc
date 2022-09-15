package database

import (
	"crud-static/models"
	"errors"
)

func GetBooks() (interface{}, error) {
	modelBooks := []models.Books{
		{
			Id:          201,
			Title:       "Harry Potter | Philosopher's Stone (1997)",
			Description: "Published in 1997",
			Author:      "J. K. Rowling",
		},
		{
			Id:          202,
			Title:       "Harry Potter | Chamber of Secrets (1998)",
			Description: "Published in 1998",
			Author:      "J. K. Rowling",
		},
		{
			Id:          203,
			Title:       "Harry Potter | Prisoner of Azkaban (1999)",
			Description: "Published in 1999",
			Author:      "J. K. Rowling",
		},
		{
			Id:          204,
			Title:       "Harry Potter | Goblet of Fire (2000)",
			Description: "Published in 2000",
			Author:      "J. K. Rowling",
		},
		{
			Id:          205,
			Title:       "Harry Potter | Order of the Phoenix (2003)",
			Description: "Published in 2003",
			Author:      "J. K. Rowling",
		},
		{
			Id:          206,
			Title:       "Harry Potter | Half-Blood Prince (2005)",
			Description: "Published in 2005",
			Author:      "J. K. Rowling",
		},
		{
			Id:          207,
			Title:       "Harry Potter | Deathly Hallows (2007)",
			Description: "Published in 2007",
			Author:      "J. K. Rowling",
		},
	}
	return modelBooks, nil
}

func GetBookById(id int) (interface{}, error) {

	books, _ := GetBooks()

	var findBook models.Books

	allBooks := books.([]models.Books)

	alreadyBook := false

	//find book by Id
	for _, book := range allBooks {
		if book.Id == id {
			findBook = book
			alreadyBook = true
		}
	}

	if !(alreadyBook) {
		return nil, errors.New("data not found")
	}

	return findBook, nil
}

func UpdateBookById(id int, book *models.Books) (interface{}, error) {
	//get books from func GetBooks
	books, _ := GetBooks()

	var findBook models.Books
	//convert interface to struct book
	allBooks := books.([]models.Books)

	//detect is book already
	alreadyBook := false

	//search book by Id
	for _, book := range allBooks {
		if book.Id == id {
			alreadyBook = true
		}
	}

	if !(alreadyBook) {
		return nil, errors.New("data not found")
	} else {
		findBook = models.Books{
			Id:          id,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
		}
	}

	return findBook, nil
}
