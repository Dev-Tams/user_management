package main

import "fmt"

type LoginError struct {
	Reason string
}

func (l LoginError) Error() string {
	return fmt.Sprintf("Login failed: %s", l.Reason)
}
