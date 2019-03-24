package main

import (
	"bufio"
	"os"
	"strings"
)

// mainパッケージが複数存在することになってしまうが、ビルドしたバイナリを置いておく
// go build -o export_env export_env.go
func main() {
	f, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), "=")
		if err := os.Setenv(str[0], str[1]); err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
