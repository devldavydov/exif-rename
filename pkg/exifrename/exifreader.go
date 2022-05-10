package exifrename

import (
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

type ExifReader struct {
	tags map[string]interface{}
}

func (er *ExifReader) LoadExifData(imgFilePath string) error {
	rawExif, err := exif.SearchFileAndExtractExif(imgFilePath)
	if err != nil {
		return err
	}

	ti := exif.NewTagIndex()
	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return err
	}

	visitor := func(ite *exif.IfdTagEntry) (err error) {
		tagName := ite.TagName()
		value, err := ite.Value()
		if err == nil {
			er.tags[tagName] = value
		}
		return nil
	}

	er.tags = make(map[string]interface{})
	_, _, err = exif.Visit(exifcommon.IfdStandardIfdIdentity, im, ti, rawExif, visitor, &exif.ScanOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (er *ExifReader) Tag(tagName string) (interface{}, bool) {
	val, ok := er.tags[tagName]
	return val, ok
}
