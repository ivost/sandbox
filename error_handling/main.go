package main

// https://blog.golang.org/go1.13-errors

/*
The Unwrap method

Go 1.13 introduces new features to the errors and fmt standard library packages
to simplify working with errors that contain other errors. The most significant of these is a convention rather than a change: an error which contains another may implement an Unwrap method returning the underlying error. If e1.Unwrap() returns e2, then we say that e1 wraps e2, and that you can unwrap e1 to get e2.

Following this convention, we can give the QueryError type above
an Unwrap method that returns its contained error:

func (e *QueryError) Unwrap() error { return e.Err }
The result of unwrapping an error may itself have an Unwrap method;
we call the sequence of errors produced by repeated unwrapping the error chain.

Examining errors with Is and As

The Go 1.13 errors package includes two new functions for examining errors: Is and As.

The errors.Is function compares an error to a value.
*/

import (
	"errors"
	"fmt"
	"log"
)

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Unwrap() error { return e.Err }

var ErrNotFound = errors.New("not found")

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

func (e *QueryError) Error() string { return e.Query + ": " + e.Err.Error() }

func runme(n int) {

	err := foo(n)

	log.Printf("=== foo returns: %v", err)

	if err == ErrNotFound {
		// something wasn't found
		log.Printf("ErrNotFound error %v", err)
	}
	// Similar to:
	//   if err == ErrNotFound { … }
	if errors.Is(err, ErrNotFound) {
		// something wasn't found
		log.Printf("Using errors.Is - ErrNotFound error %v", err)
	}

	if er, ok := err.(*NotFoundError); ok {
		log.Printf("NotFoundError error %v", er)
	}

	// Similar to:
	//   if e, ok := err.(*QueryError); ok { … }
	var e *QueryError
	if errors.As(err, &e) {
		// err is a *QueryError, and e is set to the error's value
		log.Printf("Using errors.Ad - NotFoundError error %v", e)
	}

}

func main() {
	for i := 0; i < 5; i++ {
		runme(i)
	}
}

// defer/recover need named return value
func foo(n int) (err error) {

	// The recover built-in function allows a program to manage behavior of a
	// panicking goroutine. Executing a call to recover inside a deferred
	// function (but not any function called by it) stops the panicking sequence
	// by restoring normal execution and retrieves the error value passed to the
	// call of panic. If recover is called outside the deferred function it will
	// not stop a panicking sequence. In this case, or when the goroutine is not
	// panicking, or if the argument supplied to panic was nil, recover returns
	// nil. Thus the return value from recover reports whether the goroutine is
	// panicking.

	defer func() {
		if perr := recover(); perr != nil {
			// this is the named error return above
			err = perr.(error)
			log.Printf("*** in defer: panic error: %v", err)
		}
	}()

	switch n {
	case 1:
		return &QueryError{
			Query: "select foo from bar",
			Err:   errors.New("query error"),
		}
	case 2:
		return ErrNotFound
	case 3:
		return &NotFoundError{Name: "foo-bar"}
	default:
		// make sure to panic with error not string
		panic(fmt.Errorf("unexpected input %v", n))
	}
	return
}
