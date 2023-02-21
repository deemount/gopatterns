package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Handler interface {
	Handle(interface{}) error
}

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s", b.Title, b.Author)
}

type Deserializer struct {
	next Handler
}

func (h *Deserializer) Handle(data interface{}) error {
	var book Book
	bytes, ok := data.([]byte)
	if !ok {
		return errors.New("Unsupported type!")
	}

	if err := json.Unmarshal(bytes, &book); err != nil {
		return err
	}

	if h.next != nil {
		return h.next.Handle(book)
	}

	return nil
}

type Logger struct {
	next Handler
}

func (h *Logger) Handle(data interface{}) error {
	book, ok := data.(Book)
	if !ok {
		return errors.New("Argument is not a Book instance")
	}

	fmt.Printf("Book received! Titlle: '%s' Author: '%s", book.Title, book.Author)
	if h.next != nil {
		return h.next.Handle(book)
	}

	return nil
}

type BookDb struct {
	store map[string]Book
	next  Handler
}

func (h *BookDb) Handle(data interface{}) error {
	book, ok := data.(Book)
	if !ok {
		return errors.New("Argument is not a Book instance")
	}

	h.store[book.Title] = book

	if h.next != nil {
		return h.next.Handle(book)
	}

	return nil
}

func main() {}
