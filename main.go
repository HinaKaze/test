package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	path, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	dir := filepath.Dir(path)
	fmt.Println(dir)
	base := filepath.Base(path)
	fmt.Println(base)

	str := "R099   "
	fmt.Println(string(str[0]))

	str = "			R099             CHANGELOG.md    CHANGELOG1.md"
	str = strings.TrimSpace(str)
	fmt.Println(str)
	re := regexp.MustCompile(" +")
	strBytes := re.ReplaceAll([]byte(str), []byte(" "))
	fmt.Println(string(strBytes))
	str = string(strBytes)
	s := strings.Split(str, " ")
	for _, i := range s {
		fmt.Println(i)
	}

	ssss := "中国"
	fmt.Println(len(ssss), len(strings.Split(ssss, "")))
	fmt.Println([]byte(ssss)[6:])
	return
}
