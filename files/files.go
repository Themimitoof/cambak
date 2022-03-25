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
	Pictures []File
	RAW      []File
	Movies   []File
}

func CollectFiles(conf config.Configuration) Files {
	files := Files{}
	source := conf.Extract.SourcePath

	if conf.Extract.ExtractPictures {
		files.Pictures = append(files.Pictures, ListFiles(source, picturesExtensions, FC_PICTURE)...)
	}

	if conf.Extract.ExtractRaws {
		files.RAW = append(files.RAW, ListFiles(source, rawExtensions, FC_RAW)...)
	}

	if conf.Extract.ExtractMovies {
		files.Movies = append(files.Movies, ListFiles(source, moviesExtensions, FC_MOVIE)...)
	}

	return files
}

func ListFiles(sourcePath string, extensions []string, category FileCategory) []File {
	var files []File
	fileSystem := os.DirFS(sourcePath)

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
					// TODO: Add a check if the file already exists or not.

					files = append(
						files,
						File{
							File:     fileInfo,
							Path:     path,
							Category: category,
						},
					)
					break
				}
			}
		}

		return nil
	})

	return files
}
