# pwgen

An example password generation utility written in Go.


## Usage

A simple command line tool can be built and run with:
```bash
$ go build cmd/pwgen/pwgen.go
$ ./pwgen
```

Available command-line flags:
- `-length` (int) length of the generated password (default 12)
- `-nouppercase` exclude uppercase A-Z characters (default false)
- `-nolowercase` exclude lowercase a-z characters (default false)
- `-number` (int) exact number of numeric 0-9 characters (default 2)
- `-special` (int) exact number of special @%!?*^& characters (default 2)

Example: generate a password that is 30 characters in length, with 6 numbers, 3 special characters, and only lowercase letters:
```
$ ./pwgen -length 30 -number 6 -special 3 -nouppercase
```

The root `pwgen` package is a library that can be imported in other programs. The key function is `pwgen.GeneratePassword`:

```go
pwgen.GeneratePassword(length int, uppercase bool, lowercase bool, number int, special int) (password string, err error)
```


## Notes

`math/rand` is seeded with `crypto/rand` to provide an adequate level of entropy to token selection and shuffling.

For simplicity, it is assumed that at least one of either uppercase or lowercase character pools must be enabled. However, if the number of numeric and special characters combined is equal to the requested length of the password, then the password will not contain any letters.

Some tests are provided, but testing the randomness of Go's random number generator is beyond the scope of this utility.
