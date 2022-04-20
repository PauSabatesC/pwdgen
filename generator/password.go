package generator

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const (
	MAX_ASCII_VALUE       = 127
	MAX_PASSWORD_LENGTH   = 512
	PASSWORD_DEFAULT_SIZE = 12
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Password struct {
	minAsciiValue   int
	maxAsciiValue   int
	startAsciiValue int
	length          int
}

func (pwd *Password) GeneratePassword() (string, error) {
	if pwd.length > MAX_PASSWORD_LENGTH {
		return "", errors.New("Password length should be < 512 characters.")
	}

	var passwd strings.Builder
	for i := 0; i < pwd.length; i++ {
		randChar, err := pwd.getRandomAsciiCode()
		if err != nil {
			panic(err)
		}
		passwd.WriteString(string(byte(pwd.startAsciiValue) + byte(randChar)))
	}

	return passwd.String(), nil
}

func (pwd *Password) getRandomAsciiCode() (int, error) {
	ok := false
	var randValue int

	for !ok {
		randValue = rand.Intn(pwd.maxAsciiValue-pwd.minAsciiValue) + pwd.minAsciiValue
		if randValue+pwd.startAsciiValue >= MAX_ASCII_VALUE {
			return -1, errors.New("ASCII values are not correctly setup as the generated number is bigger than maximum ascii value.")
		}
		ok = checkIsInWhitelist(randValue + pwd.startAsciiValue)
	}

	return randValue, nil
}

func checkIsInWhitelist(asciiValue int) bool {
	odd_ascii_values := []int{int('\''), int(','), int('.'), int('/'), int(':'), int('<'), int('>'), int('['), int(']'), int('\\'), int('_'), int('`'), int('{'), int('}'), int('|'), int('~'), int(';'), int('('), int(')')}
	for _, char := range odd_ascii_values {
		if asciiValue == char {
			return false
		}
	}
	return true
}
