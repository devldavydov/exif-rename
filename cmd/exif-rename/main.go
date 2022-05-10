package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/devldavydov/exif-rename/pkg/exifrename"
)

func main() {
	imgPath, dryRun := parseArgs()
	fmt.Printf("Images path: %s\nDry run: %t\n\n", imgPath, dryRun)

	exifRenamer := exifrename.CreateExifRenamer(imgPath, dryRun)
	exifRenamer.Run()

	fmt.Println("Press any key...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func parseArgs() (string, bool) {
	imgPath := flag.String("path", "", "images path")
	dryRun := flag.Bool("dry", false, "perform only dry run")
	flag.Parse()
	if *imgPath == "" {
		fmt.Println("ERROR: images path not set")
		os.Exit(1)
	}
	return *imgPath, *dryRun
}
