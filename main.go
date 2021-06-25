package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("this is a %s and this is %s.\n", yellow("warning"), red("error"))

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	fmt.Printf("this %s rocks!\n", info("package"))

}
