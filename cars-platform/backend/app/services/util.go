package services

const Token = "123456"

func CheckToken(token string) bool {
	if token == Token {
		return true
	}
	return false
}
