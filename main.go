package main

import (
	"gofr.dev/pkg/gofr"
)

// defing the structure of the book.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Initialization.
	app := gofr.New()

	// POST method to add a new book.
	app.POST("/book/{title}/{author}", addBook)

	// GET method to retrieve a book by ID.
	app.GET("/book/{id}", getBook)

	// GET method to retrieve all books.
	app.GET("/books", getAllBooks)

	// PUT method to update a book by ID.
	app.PUT("/book/{id}/{title}/{author}", updateBook)

	// DELETE method to delete a book by ID.
	app.DELETE("/book/{id}", deleteBook)

	// Starting the server.
	app.Start()
}

// creating function for adding the book in the database..
func addBook(ctx *gofr.Context) (interface{}, error) {
	title := ctx.PathParam("title")
	author := ctx.PathParam("author")

	// Inserting a book row into the database.
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO books (title, author) VALUES (?, ?)", title, author)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// creating function for retriving the book by id from the database..
func getBook(ctx *gofr.Context) (interface{}, error) {
	bookID := ctx.PathParam("id")

	var book Book
	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM books WHERE id = ?", bookID).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// creating function for retriving all the book in the database..
func getAllBooks(ctx *gofr.Context) (interface{}, error) {
	var books []Book

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//itrating in rows
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil
}

// creating function for modifing the book stored in the database through id..
func updateBook(ctx *gofr.Context) (interface{}, error) {
	bookID := ctx.PathParam("id")
	title := ctx.PathParam("title")
	author := ctx.PathParam("author")

	// Updating the book information in the database..
	_, err := ctx.DB().ExecContext(ctx, "UPDATE books SET title=?, author=? WHERE id=?", title, author, bookID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// creating function for deleting the book in the database through id..
func deleteBook(ctx *gofr.Context) (interface{}, error) {
	bookID := ctx.PathParam("id")

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM books WHERE id=?", bookID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
