package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fileUrl := "https://github.com/iamaldren/xml-to-pojo/archive/master.zip"

	path, project, err := GetPath(fileUrl)
	if err != nil {
		panic(err)
	}

	fileName := strings.Trim(project,"/") + ".zip"
	if err := DownloadFile(fileName, path, fileUrl); err != nil {
		panic(err)
	}

}

func GetPath(fileUrl string) (string, string, error)  {
	var splitter []string = strings.SplitAfter(fileUrl, "/")
	github := splitter[2]
	user := splitter[3]
	project := splitter[4]

	rootFolder := os.Getenv("GOPATH") + "/src/"
	subFolder := github + user + project //+ strings.Trim(splitter[4],"/") + ".zip"

	path := filepath.Join(rootFolder, subFolder)
	err := exists(path)

	return path, project, err
}

func exists(path string) error {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		} else {
			return err
		}
	}

	return err
}

func DownloadFile(fileName string, path string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	err = MoveFile(fileName, path)

	return err
}

func MoveFile(fileName string, path string) error {
	err := os.Rename(fileName, path + "/" + fileName)
	return err
}
