package snippet

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateSecretKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	encodedKey := base64.URLEncoding.EncodeToString(key)
	return encodedKey, nil
}

func GenSecretmain() {
	s, err := generateSecretKey(32)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)


}