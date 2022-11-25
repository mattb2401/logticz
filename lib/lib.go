package lib

import (
	"github.com/matthewhartstonge/argon2"
)


//HashArgonPassword handles argon2 password
func HashArgonPassword(password string) (string, error) {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

//CheckArgonPassword validates argon2 password
func CheckArgonPassword(password, hash string) bool {
	if ok, err := argon2.VerifyEncoded([]byte(password), []byte(hash)); !ok {
		if err != nil {
			return false
		}
		return false
	}
	return true
}
