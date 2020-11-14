package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/dto"
)

func encode() (string, error) {

	file, err := os.Open("asset/logo.jpg")
	if err != nil {
		return "", nil
	}
	defer file.Close()

	fileInfo, err := file.Stat() //FileInfo interface
	if err != nil {
		return "", err
	}

	data := make([]byte, fileInfo.Size())
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data), nil
}

func request(info dto.CreateImageRequest) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		"http://localhost:8080/image/detour/create",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

func main() {
	image, err := encode()
	if err != nil {
		log.Fatal(err)
	}

	info := dto.CreateImageRequest{"logo", image}

	body, err := request(info)
	if err != nil {
		log.Fatal(body)
	}

	fmt.Println(string(body))
}
