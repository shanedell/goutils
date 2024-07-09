# goutils

Golang utilities.

- `env`
    - Wrapper functions around `os.Getenv` to simplify retrieval of environment variables
        - Supported types: `string, int, bool, float32, float64, complex64 and complex128`.
    - Wrapper functions around `godotenv.Load`
        - These either return its error or ignore the error when trying to source the env file passed.
- `path`
    - Method to check if passed path exists.
- `slice`
    - Functions to check if list contains a value.
        - Supported types: `string, int, int32, int64, bool, float32, float64, complex64 and complex128`.
