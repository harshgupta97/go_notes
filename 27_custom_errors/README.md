# Custom Errors

It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

A custom error type usually has the suffix “Error”.

Adding this Error method makes argError implement the error interface.

Return our custom error.

errors.As is a more advanced version of errors.Is. It checks that a given error (or any error in its chain) matches a specific error type and converts to a value of that type, returning true. If there’s no match, it returns false.