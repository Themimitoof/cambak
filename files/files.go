package files

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/themimitoof/cambak/config"
)

type Files struct {
	Pictures     []File
	RAW          []File
	Movies       []File
	TotalFiles   uint
	SkippedFiles uint
}

func CollectFiles(conf config.Configuration) Files {
	files := Files{}
	source := conf.Extract.SourcePath

	if conf.Extract.ExtractPictures {
		pictures, ignoredPictures := ListFiles(conf, source, picturesExtensions, FC_PICTURE)
		files.Pictures = append(files.Pictures, pictures...)
		files.SkippedFiles += ignoredPictures
	}

	if conf.Extract.ExtractRaws {
		raws, ignoredRaws := ListFiles(conf, source, rawExtensions, FC_RAW)
		files.RAW = append(files.RAW, raws...)
		files.SkippedFiles += ignoredRaws
	}

	if conf.Extract.ExtractMovies {
		movies, ignoredMovies := ListFiles(conf, source, moviesExtensions, FC_MOVIE)
		files.Movies = append(files.Movies, movies...)
		files.SkippedFiles += ignoredMovies
	}

	files.TotalFiles = uint(len(files.Pictures) + len(files.RAW) + len(files.Movies))
	return files
}

func ListFiles(conf config.Configuration, sourcePath string, extensions []string, category FileCategory) ([]File, uint) {
	var files []File
	fileSystem := os.DirFS(sourcePath)
	var skippedFiles uint = 0

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		// Skip the current folder
		if path == "." {
			return nil
		}

		if err != nil {
			color.Red("Unable to manage %s (err: %s)", path, err)
			return nil
		}

		path = fmt.Sprintf("%s/%s", sourcePath, path)

		if d.Type().IsRegular() {
			fileInfo, _ := d.Info()
			splittedFileName := strings.SplitAfter(fileInfo.Name(), ".")
			ext := strings.ToLower(
				splittedFileName[len(splittedFileName)-1],
			)

			// Ignore all "hidden" files
			if fileInfo.Name()[0] == '.' {
				return nil
			}

			for _, fileExt := range extensions {
				if ext == fileExt {
					file := File{
						File:     fileInfo,
						Path:     path,
						Category: category,
					}

					// Check if the file already exists or not
					fileDest := fmt.Sprintf("%s/%s", file.GenerateDestinationPath(conf), fileInfo.Name())

					if _, err := os.Open(fileDest); err == nil && conf.Extract.DestinationConflict == config.DEST_CONFLICT_SKIP {
						skippedFiles++
						return nil
					}

					// Append the file to the list of files
					files = append(files, file)
					break
				}
			}
		}

		return nil
	})

	return files, skippedFiles
}
