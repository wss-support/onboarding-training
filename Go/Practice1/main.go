package main

import (
	"fmt"
	"github.com/dsnet/compress/brotli"
	"github.com/gookit/color"
	"github.com/mholt/archiver/v3"
)

func main() {
	fmt.Println("Hello World!")
}

func DoNothing1() {
	_ = archiver.Unarchive("", "")
}

func DoNothing2(c color.Color256, text string) string {
	return c.Sprintf(text)
}

func DoNothing3() {
	_ = brotli.ReaderConfig{}
}
