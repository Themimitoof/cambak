package files

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/themimitoof/cambak/config"
)

type FileCategory int8

const (
	FC_PICTURE FileCategory = iota
	FC_RAW
	FC_MOVIE
)

type File struct {
	Path     string
	File     fs.FileInfo
	Category FileCategory
}

func (file File) ExtractFile(conf config.Configuration, dest string) error {
	errMsg := "Unable to copy file '%s' to his destination folder. Err: %s\n"
	fileDest := fmt.Sprintf("%s/%s", dest, file.File.Name())

	if conf.Extract.CleanAfterCopy {
		// Move the file
		if err := os.Rename(file.Path, fileDest); err != nil {
			color.Red(errMsg, file.Path, err)
			return err
		}
	} else {
		// Copy the file instead of moving it
		fl, err := ioutil.ReadFile(file.Path)
		if err != nil {
			color.Red(errMsg, file.Path, err)
			return err
		}

		if err = ioutil.WriteFile(fileDest, fl, 0755); err != nil {
			color.Red(errMsg, file.Path, err)
			return err
		}
	}

	return nil
}

func (file File) PrepareFileDestinationFolder(conf config.Configuration) (string, error) {
	destinationPath := file.GenerateDestinationPath(conf)

	if err := os.MkdirAll(destinationPath, 0755); err != nil {
		color.Red(
			"Unable to create the destination folder '%s' for file '%s'. Err: %s\n",
			destinationPath,
			file.File.Name(),
			err,
		)

		return "", err
	}
	return destinationPath, nil
}
