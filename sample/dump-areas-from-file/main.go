package main

import (
	"fmt"
	"os"

	"github.com/pen/jma-go/area"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <area.json>\n", os.Args[0])
		fmt.Println(`
download area.json by:
    curl https://www.jma.go.jp/bosai/common/const/area.json > area.json
and retry me`)
		os.Exit(2)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("open failed: %s error: %s\n", os.Args[1], err.Error())
		os.Exit(1)
	}

	areas, err := area.Parse(f)
	if err != nil {
		fmt.Printf("parse failed: %s error: %s", os.Args[1], err.Error())
		os.Exit(1)
	}

	for _, area := range areas {
		fmt.Printf("%s\n", area.Name)
	}
}
