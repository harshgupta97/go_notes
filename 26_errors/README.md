# Errors

In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.

See the documentation of the errors package and this blog post for additional details.

By convention, errors are the last return value and have type error, a built-in interface.

errors.New constructs a basic error value with the given error message.

A nil value in the error position indicates that there was no error.

A sentinel error is a predeclared variable that is used to signify a specific error condition.

We can wrap errors with higher-level errors to add context. The simplest way to do this is with the %w verb in fmt.Errorf. Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.

It’s common to use an inline error check in the if line.

errors.Is checks that a given error (or any error in its chain) matches a specific error value. This is especially useful with wrapped or nested errors, allowing you to identify specific error types or sentinel errors in a chain of errors.