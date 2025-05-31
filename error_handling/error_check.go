package main

import (
    "errors"
    "fmt"
    "os"
)

// Sentinel error
var ErrPermissionDenied = errors.New("permission denied")

// Custom error type
type ValidationError struct {
    Field string
    Msg   string
}

func (v ValidationError) Error() string {
    return fmt.Sprintf("Validation failed on '%s': %s", v.Field, v.Msg)
}

// Function that returns a basic error
func basicError() error {
    return errors.New("something went wrong")
}

// Function that wraps another error
func wrappedError() error {
    err := basicError()
    return fmt.Errorf("wrapped error: %w", err)
}

// Function that returns a sentinel error
func checkPermission(user string) error {
    if user != "admin" {
        return ErrPermissionDenied
    }
    return nil
}

// Function that returns a custom error type
func validate(field string, value string) error {
    if value == "" {
        return ValidationError{Field: field, Msg: "cannot be empty"}
    }
    return nil
}

// Panic example with recover
func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("unexpected failure in riskyOperation")
}

// Function with return-early pattern and error wrapping
func process(user string, data string) error {
    if err := checkPermission(user); err != nil {
        return fmt.Errorf("permission check failed: %w", err)
    }

    if err := validate("data", data); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }

    return nil
}

// Simulates a multi-step operation and returns the first error
func simpleValidationSteps() error {
    if err := validate("email", ""); err != nil {
        return err
    }

    if err := validate("password", ""); err != nil {
        return err
    }

    return nil
}

func main() {
    fmt.Println("\n--- Basic Error ---")
    if err := basicError(); err != nil {
        fmt.Println("Basic Error:", err)
    }

    fmt.Println("\n--- Wrapped Error ---")
    err := wrappedError()
    fmt.Println("Wrapped Error:", err)

    unwrapped := errors.Unwrap(err)
    fmt.Println("Unwrapped:", unwrapped)

    fmt.Println("\n--- Sentinel Error with errors.Is ---")
    if err := checkPermission("guest"); errors.Is(err, ErrPermissionDenied) {
        fmt.Println("Detected permission denied")
    }

    fmt.Println("\n--- Custom Error Type with errors.As ---")
    err = validate("username", "")
    var ve ValidationError
    if errors.As(err, &ve) {
        fmt.Println("Caught validation error on:", ve.Field)
    }

    fmt.Println("\n--- Panic and Recover ---")
    riskyOperation()

    fmt.Println("\n--- Return Early Pattern with Wrapping ---")
    err = process("guest", "")
    fmt.Println("Process Error:", err)

    fmt.Println("\n--- Simulated Multi-Step Validation ---")
    if err := simpleValidationSteps(); err != nil {
        fmt.Println("First validation error:", err)
    }

    fmt.Println("\n--- Program Completed ---")
    os.Exit(0)
}
