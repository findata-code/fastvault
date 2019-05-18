package utils

import "encoding/base64"

func EncodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func DecodeBase64(b []byte) ([]byte, error){
	if out, err := base64.StdEncoding.DecodeString(string(b)); err != nil {
		return nil, err
	}else{
		return out, nil
	}
}