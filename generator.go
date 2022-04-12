package pwgen

import (
	"fmt"

	"github.com/mbwk/pwgen/random"
)

// SelectCase returns a character pool given the selected cases
func SelectCase(uppercase bool, lowercase bool) (string, error) {
	cases := NoneCase
	if lowercase {
		cases += Lower
	}
	if uppercase {
		cases += Upper
	}

	if cases == NoneCase {
		return "", fmt.Errorf("neither uppercase nor lowercase were enabled, need at least one")
	}
	return Cases[cases], nil
}

// GeneratePassword pseudo-randomly generates a password given certain parameters
func GeneratePassword(length int, uppercase bool, lowercase bool, number int, special int) (string, error) {
	if length < 6 {
		return "", fmt.Errorf("password length too short")
	}
	if (number + special) > length {
		return "", fmt.Errorf("number of numeric and special characters requested exceeds requested password length")
	}

	alphabetPool, err := SelectCase(uppercase, lowercase)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, length)
	for index := length - 1; index > -1; index-- {
		var character byte
		if special > 0 {
			character = random.RandomSelection(Symbols)
			special--
		} else if number > 0 {
			character = random.RandomSelection(Numbers)
			number--
		} else {
			character = random.RandomSelection(alphabetPool)
		}
		buffer[index] = character
	}

	random.Shuffle(buffer)

	return string(buffer), nil
}
