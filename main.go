package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	fileUrl := "https://github.com/iamaldren/xml-to-pojo/archive/master.zip"

	if err := DownloadFile("xml-to-pojo.zip", fileUrl); err != nil {
		panic(err)
	}

}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
