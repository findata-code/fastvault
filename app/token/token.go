package token

import (
	"fastvault/app/utils"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GetTokens(input []byte) ([][]byte, error) {
	rand.Seed(time.Now().UnixNano())

	if len(input) > 64 {
		input = utils.ToSha512(input)
	}

	fullToken, err := utils.CFBEncrypt(GenerateKey(), input)
	if err != nil {
		return nil, err
	}

	output := make([][]byte, 0)

	numberOfPiece := rand.Int()%4 + 1
	for i := 0; i < numberOfPiece-1; i++ {
		cutAt := rand.Int() % len(fullToken)
		output = append(output, fullToken[:cutAt])
		fullToken = fullToken[cutAt:]
	}
	output = append(output, fullToken)

	for i, v := range output {
		output[i] = []byte(strings.Trim(fmt.Sprintf("%s", utils.EncodeBase64(v)), "="))
	}

	return output, nil
}

func GenerateKey() []byte {
	rand.Seed(time.Now().UnixNano())

	return []byte(fmt.Sprintf("%d%d%d%d",
		rand.Int()%8999+1000,
		rand.Int()%8999+1000,
		rand.Int()%8999+1000,
		rand.Int()%8999+1000))
}
