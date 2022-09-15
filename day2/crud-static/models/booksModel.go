package models

type Books struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type BooksDTO struct {
	Id          int
	Title       string
	Description string
	Author      string
}

func AssembBooksDTO(b *BooksDTO) *Books {
	return &Books{
		Title:       b.Title,
		Description: b.Description,
		Author:      b.Author,
	}
}

func ToBooksDTO(book *Books) *BooksDTO {
	return &BooksDTO{
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}
}
