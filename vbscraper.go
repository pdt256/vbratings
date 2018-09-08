package vbscraper

import (
	"log"
)

type NotFoundError struct{ Err error }

func (e *NotFoundError) Error() string { return "not found" }

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
