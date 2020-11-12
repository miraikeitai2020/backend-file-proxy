package decode

import (
	"encoding/base64"
	"os"
)

func Image(fileName, image string) error {
	data, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(data)
	return nil
}
