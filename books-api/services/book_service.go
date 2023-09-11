package services

import (
	"books-api/cache"
	bookDao "books-api/daos/book"
	"books-api/dtos"
	model "books-api/models"
	e "books-api/utils/errors"
	"fmt"
	json "github.com/json-iterator/go"
	"time"
)

type bookService struct{}

type bookServiceInterface interface {
	GetBook(id string) (dtos.BookDto, e.ApiError)

	InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError)
}

var (
	BookService bookServiceInterface
)

func init() {
	BookService = &bookService{}
}

func (s *bookService) GetBook(id string) (dtos.BookDto, e.ApiError) {

	time.Sleep(15 * time.Second)

	// get from cache
	var cacheDTO dtos.BookDto
	cacheBytes := cache.Get(id)
	if cacheBytes != nil {
		fmt.Println("Found in cache!")
		_ = json.Unmarshal(cacheBytes, &cacheDTO)
		return cacheDTO, nil
	}

	var book model.Book = bookDao.GetById(id)
	var bookDto dtos.BookDto

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, e.NewBadRequestApiError("book not found")
	}
	bookDto.Name = book.Name
	bookDto.Id = book.Id.Hex()

	// save in cache
	bookBytes, _ := json.Marshal(bookDto)
	cache.Set(id, bookBytes)
	fmt.Println("Saved in cache!")

	return bookDto, nil
}

func (s *bookService) InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError) {

	var book model.Book

	book.Name = bookDto.Name

	book = bookDao.Insert(book)

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, e.NewBadRequestApiError("error in insert")
	}
	bookDto.Id = book.Id.Hex()

	return bookDto, nil
}
