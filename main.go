package main

import (
	"bufio"
	"flag"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("DomCol\n. DomCol is a CLI tool that helps you")
	fmt.Println("find the dominant colors in an image file.")
	// Get flags
	imgPath := flag.String("imgpath", "", "Path to the image file")
	rgba := flag.Bool("rgba", false, "Returns the dominant colors in RGB form")
	hex := flag.Bool("hex", false, "Returns the dominant colors in hex form")

	// Verify that the file is an image file
	imgPattern := `^[a-zA-Z].*\.(jpg|png)$`
	re, err := regexp.Compile(imgPattern)
	if err != nil {
		fmt.Println("Error compiling regex for image file validation:", err)
		return
	}
	isValidImg := re.MatchString(*imgPath)
	if !isValidImg {
		fmt.Println("Image is not a valid image file. Try one ending with jpg or png")
	}
	// Read input file from file path
}
