package exifrename

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ExifRenamer struct {
	imgPath    string
	dryRun     bool
	duplicates map[string]int
}

func CreateExifRenamer(imgPath string, dryRun bool) *ExifRenamer {
	return &ExifRenamer{
		imgPath:    imgPath,
		dryRun:     dryRun,
		duplicates: make(map[string]int),
	}
}

func (rn *ExifRenamer) Run() {
	err := filepath.WalkDir(rn.imgPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(path), ".JPG") {
			exifReader := &ExifReader{}
			fErr := exifReader.LoadExifData(path)
			if fErr != nil {
				fmt.Printf("ERROR: EXIF error, path=%s err=%v\n", path, err)
				return nil
			}
			valDateTimeOriginal, ok := exifReader.GetTag("DateTimeOriginal")
			if !ok {
				fmt.Printf("ERROR: EXIF tag DateTimeOriginal not exists, path=%s\n", path)
				return nil
			}
			formatDateTimeOriginal, fErr := rn.formatDate(valDateTimeOriginal.(string))
			if fErr != nil {
				fmt.Printf("ERROR: EXIT DateTimeOriginal parse error, path=%s, err=%v\n", path, fErr)
				return nil
			}
			rn.renameFile(path, filepath.Join(filepath.Dir(path), fmt.Sprintf("%s.jpg", formatDateTimeOriginal)))
		}
		return nil
	})
	if err != nil {
		fmt.Println("ERROR: images path walk error")
	}
}

func (rn *ExifRenamer) renameFile(oldFilePath string, newFilePath string) {
	var suffix string
	if rn.fileExists(newFilePath) {
		val, exists := rn.duplicates[newFilePath]
		if !exists {
			val = 1
		}
		rn.duplicates[newFilePath] = val + 1
		suffix = fmt.Sprintf("_%d.jpg", val+1)
	}
	if suffix != "" {
		pos := strings.LastIndex(newFilePath, ".")
		newFilePath = newFilePath[:pos] + suffix
	}
	fmt.Printf("Move %s -> %s\n", oldFilePath, newFilePath)
	if !rn.dryRun {
		err := os.Rename(oldFilePath, newFilePath)
		if err != nil {
			fmt.Printf("ERROR: move to new path=%s err=%v\n", newFilePath, err)
		}
	}
}

func (rn *ExifRenamer) fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (rn *ExifRenamer) formatDate(strDate string) (string, error) {
	dt, err := time.Parse("2006:01:02 15:04:05", strDate)
	if err != nil {
		return "", err
	}
	return dt.Format("20060102_150405"), nil
}
