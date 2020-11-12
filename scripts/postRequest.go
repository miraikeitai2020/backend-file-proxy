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

func encode() string {

	file, _ := os.Open("asset/logo.jpg")
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

func request(info dto.CreateImageRequest) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		"http://localhost:8080/image/create",
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
	image := encode()

	info := dto.CreateImageRequest{"logo", image}

	body, err := request(info)
	if err != nil {
		log.Fatal(body)
		return
	}

	fmt.Println(string(body))
}
