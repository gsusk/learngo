package creditcard

import "fmt"

type creditcard struct {
	number string
}

func New(number string) (creditcard, error) {
	if number == "" {
		return creditcard{}, fmt.Errorf("Invalid card number: %q", number)
	}
	return creditcard{number}, nil
}

func (cc *creditcard) Number() string {
	return cc.number
}
