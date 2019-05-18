package services

import (
	"fastvault/app/token"
	"fmt"
	"strings"
)

type SecretService interface {
	CreateSecret(value []byte) (string, error)
}

type SecretServiceImpl struct{}

func NewSecretService() SecretService {
	return &SecretServiceImpl{}
}

func (ss *SecretServiceImpl) CreateSecret(value []byte) (string, error) {
	tokens, err := token.GetTokens(value)
	if err != nil {
		return "", err
	}

	output := make([]string, 0)

	for _, v := range tokens {
		output = append(output, fmt.Sprintf("%s", v))
	}

	return strings.Join(output, "-"), nil
}
