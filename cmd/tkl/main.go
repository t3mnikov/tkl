package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/t3mnikov/tkl/internal/layout"
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "tkl <en|ru> <text>\n")
	}

	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	inputLayout := flag.Arg(0)
	text := strings.Join(flag.Args()[1:], " ") // будет работать "tkl en цудсщьу ghbdtn"

	parsedLayout, err := layout.ParseLayout(inputLayout)
	if err != nil {
		fmt.Println("Ошибка при разборе раскладки клавиатуры:", err)
		os.Exit(1)
	}

	output, err := layout.Convert(parsedLayout, text)
	if err != nil {
		fmt.Println("Ошибка при конвертации текста:", err)
		os.Exit(1)
	}

	fmt.Println(output)
}
