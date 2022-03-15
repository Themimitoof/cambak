package files

import (
	"fmt"
	"strings"

	"github.com/themimitoof/cambak/config"
)

// Returns a string with the destination path of the given file.
//
// List of available verbs:
//   - %y: Year
//   - %m: Month
//   - %d: Day
//   - %n: Camera name
//   - %t: Media type (Pictures, RAW, Movies)
func (file File) GenerateDestinationPath(conf config.Configuration) string {
	destPath := conf.Extract.DestinationPath
	format := conf.Extract.DestinationFormat

	// TODO: Implement this part to use the created date.
	createDate := file.File.ModTime()

	format = strings.ReplaceAll(format, "%y", fmt.Sprintf("%d", createDate.Year()))
	format = strings.ReplaceAll(format, "%m", fmt.Sprintf("%02d", createDate.Month()))
	format = strings.ReplaceAll(format, "%d", fmt.Sprintf("%02d", createDate.Day()))
	format = strings.ReplaceAll(format, "%n", conf.Extract.CameraName)

	switch file.Category {
	case FC_PICTURE:
		format = strings.ReplaceAll(format, "%t", "Pictures")
	case FC_RAW:
		format = strings.ReplaceAll(format, "%t", "RAW")
	case FC_MOVIE:
		format = strings.ReplaceAll(format, "%t", "Movies")
	}

	if destPath[len(destPath)-1] != '/' {
		destPath += "/"
	}

	destPath += format

	return destPath
}
