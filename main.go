package main

import (
	"bufio"
	"flag"
	"fmt"
	"regexp"
)

func main(){
	fmt.Println("DomCol\n. DomCol is a CLI tool that helps you") 
	fmt.Println("find the dominant colors in an image file.")			
	// Get flags
	imgPath := flag.String("imgpath", "", "Path to the image file")
	rgba := flag.Bool("rgba", false, "Returns the dominant colors in RGB form")
	hex := flag.Bool("hex", false, "Returns the dominant colors in hex form")

	
	
}
