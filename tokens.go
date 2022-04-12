package pwgen

const (
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Mixedcase = Uppercase + Lowercase
	Numbers   = "0123456789"
	Symbols   = "@%!?*^&"

	NoneCase = -1
	Lower    = 1
	Upper    = 2
	Mixed    = 3
)

type Case int

var (
	Cases = []string{
		Lowercase,
		Uppercase,
		Mixedcase,
	}
)
