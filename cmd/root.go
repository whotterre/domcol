package cmd

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"regexp"
	"sort"

	"github.com/spf13/cobra"
)

var (
	imgPath  string
	wantsRgb bool
	wantsHex bool
	topN     int
)


var rootCmd = &cobra.Command{
	Use: "domcol",
	Short: "DomCol is a CLI tool that helps you find the dominant colors in an image file.",
	Run: func(cmd *cobra.Command, args []string) {
		runDomCol()
	},
}


func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	rootCmd.PersistentFlags().StringVarP(&imgPath, "imgpath", "i", "", "Path to the image file (required)")
	rootCmd.PersistentFlags().BoolVar(&wantsRgb, "rgba", false, "Returns the dominant colors in RGB form")
	rootCmd.PersistentFlags().BoolVar(&wantsHex, "hex", false, "Returns the dominant colors in hex form")
	rootCmd.PersistentFlags().IntVar(&topN, "top", 5, "Number of top dominant colors to show")
	rootCmd.MarkPersistentFlagRequired("imgpath")
}
func runDomCol() {
	// Verify that the file is an image file
	imgPattern := `\.(jpg|png)$`
	re, err := regexp.Compile(imgPattern)
	if err != nil {
		fmt.Println("Error compiling regex for image file validation:", err)
		return
	}
	isValidImg := re.MatchString(imgPath)

	if !isValidImg {
		fmt.Println("Image is not a valid image file. Try one ending with jpg or png")
	}
	// Read input file from file path
	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println("An error occurred while opening the image file:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("An error occurred while decoding the image file:", err)
	}
	// Make frequency map
	bounds := img.Bounds()
	freqMap := make(map[[3]uint32]int) // Map from [R, G, B, A] to frequency

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixel := [3]uint32{r, g, b}
			freqMap[pixel]++
		}
	}
	// Sort by count
	type colorFreq struct {
		Color [3]uint32
		Count int
	}

	var sortedColors []colorFreq
	for color, count := range freqMap {
		sortedColors = append(sortedColors, colorFreq{color, count})
	}

	sort.Slice(sortedColors, func(i, j int) bool {
		return sortedColors[i].Count > sortedColors[j].Count
	})

	if wantsRgb {
		for i, cf := range sortedColors {
			if i >= topN {
				break
			}
			fmt.Println(convToEightBitRGB(cf.Color), cf.Count)
		}
	}

	if wantsHex {
		for i, cf := range sortedColors {
			if i >= topN {
				break
			}
			fmt.Println(convToHexRGB(cf.Color), cf.Count)
		}
	}

}

func convToEightBitRGB(color [3]uint32) string {
	r := color[0] * 255 / 65535
	g := color[1] * 255 / 65535
	b := color[2] * 255 / 65535

	return fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
}

func convToHexRGB(color [3]uint32) string {
	r := color[0] * 255 / 65535
	g := color[1] * 255 / 65535
	b := color[2] * 255 / 65535

	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
