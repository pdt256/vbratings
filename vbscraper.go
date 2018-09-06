package vbscraper

type NotFoundError struct{ Err error }

func (e *NotFoundError) Error() string { return "not found" }
