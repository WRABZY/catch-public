package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	outputFilePath = "game/assets.go"
	packageName    = "game"

	inputDir       = "game/assetsgen/png/"
	inputExtension = ".png"

	UICellPixels = 70
)

func main() {
	_ = os.Remove(outputFilePath)
	output, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}

	entries, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}
	files := make([]string, 0)

	for _, entry := range entries {
		ename := entry.Name()
		if strings.Contains(ename, inputExtension) {
			files = append(files, ename[0:strings.Index(ename, ".")])
		}
	}

	fmt.Fprintln(output, `package `+packageName)
	fmt.Fprintln(output)
	fmt.Fprintln(output, `// This file was generated automatically.`)
	fmt.Fprintln(output, `// Do not modify it.`)
	fmt.Fprintln(output, `// To regenerate, use the command: `)
	fmt.Fprintln(output, `// go run game/assetsgen/codegen.go `)
	fmt.Fprintln(output)
	fmt.Fprintln(output, `import (`)
	fmt.Fprintln(output, `	"image"`)
	fmt.Fprintln(output, `	"image/color"`)
	fmt.Fprintln(output, `)`)
	fmt.Fprintln(output)
	fmt.Fprintln(output, `var (`)
	for _, f := range files {
		fmt.Fprintln(output, `	`+f+` *image.NRGBA`)
	}
	fmt.Fprintln(output, `)`)
	fmt.Fprintln(output)
	fmt.Fprintln(output, `func init() {`)
	for _, f := range files {
		input, err := os.OpenFile(inputDir+f+inputExtension, os.O_RDONLY, 0777)
		if err != nil {
			log.Fatal(err)
		}
		defer input.Close()

		img, _, err := image.Decode(input)
		if err != nil {
			log.Fatal(err)
		}
		width := img.Bounds().Dx()
		height := img.Bounds().Dy()
		if width == 1 && height == 1 {
			fmt.Fprintln(output, `	`+f+` = image.NewNRGBA(image.Rect(0, 0, `+fmt.Sprint(UICellPixels)+`, `+fmt.Sprint(UICellPixels)+`))`)

			c := color.NRGBAModel.Convert(img.At(0, 0)).(color.NRGBA)
			fmt.Fprintln(output, `	for y := 0; y < `+fmt.Sprint(UICellPixels)+`; y++ {`)
			fmt.Fprintln(output, `		for x := 0; x < `+fmt.Sprint(UICellPixels)+`; x++ {`)
			fmt.Fprintln(output, `			`+f+`.Set(x, y, color.NRGBA{`+fmt.Sprint(c.R)+`, `+fmt.Sprint(c.G)+`, `+fmt.Sprint(c.B)+`, `+fmt.Sprint(c.A)+`})`)
			fmt.Fprintln(output, `		}`)
			fmt.Fprintln(output, `	}`)
		} else {
			fmt.Fprintln(output, `	`+f+` = image.NewNRGBA(image.Rect(0, 0, `+fmt.Sprint(width)+`, `+fmt.Sprint(height)+`))`)

			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
					if c.A != 0 {
						fmt.Fprintln(output, `	`+f+`.Set(`+fmt.Sprint(x)+`, `+fmt.Sprint(y)+`, color.NRGBA{`+fmt.Sprint(c.R)+`, `+fmt.Sprint(c.G)+`, `+fmt.Sprint(c.B)+`, `+fmt.Sprint(c.A)+`})`)
					}
				}
			}
		}
		fmt.Fprintln(output)
	}

	fmt.Fprintln(output, `}`)
	output.Close()

	err = exec.Command(`gofmt`, `-w`, outputFilePath).Run()
	if err != nil {
		log.Fatal(err)
	}
}
