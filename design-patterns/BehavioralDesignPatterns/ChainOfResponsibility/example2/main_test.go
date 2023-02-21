package main

import (
	"encoding/json"
	"testing"
)

func TestCORPattern(t *testing.T) {
	book := Book{"goPatterns", "Salvatore Gonda"}
	bytes, _ := json.Marshal(book)
	bookDb := &BookDb{
		make(map[string]Book),
		nil,
	}
	logger := &Logger{
		next: bookDb,
	}
	chain := &Deserializer{
		next: logger,
	}

	t.Run(`It will be unmurshalled on deserializer handler, then it will be logged
		  on the logger handler and at the end, it will be persisted by a third handler`,
		func(t *testing.T) {

			err := chain.Handle(bytes)
			if err != nil {
				t.Fatalf("Something went wrong: %s", err.Error())
			}

			storedBook := bookDb.store[book.Title]
			if storedBook != book {
				t.Errorf("Expected %T but got %T", book, storedBook)
			}
		})

}
