package utils

import (
	"io/ioutil"
	"log"
	"os"
)

// GetFile ファイルの中身を取得する
func GetFile(filepath string) string {
	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return convertToString(body)
}

func convertToString(body []byte) string {
	return string(body)
}

// Exists 存在チェック
func Exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
