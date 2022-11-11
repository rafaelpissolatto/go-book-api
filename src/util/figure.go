package util

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/raphamorim/go-rainbow"
)

func Figure() {
	myFigure := figure.NewFigure("DEVBOOK API", "univers", true)
	fmt.Println(rainbow.Bold(rainbow.Hex("#8E44AD", myFigure.String())))
	fmt.Println("")
}
