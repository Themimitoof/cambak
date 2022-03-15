package files

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

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

	files.Pictures = append(files.Pictures, ListFiles(source, picturesExtensions, FC_PICTURE)...)
	files.RAW = append(files.Pictures, ListFiles(source, rawExtensions, FC_RAW)...)
	files.Movies = append(files.Pictures, ListFiles(source, moviesExtensions, FC_MOVIE)...)

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
			fmt.Printf("Unable to manage %s (err: %s)", path, err)
			return nil
		}

		path = fmt.Sprintf("%s/%s", sourcePath, path)

		if d.IsDir() {
			files = append(files, ListFiles(path, extensions, category)...)
		} else if d.Type().IsRegular() {
			fileInfo, _ := d.Info()
			splittedFileName := strings.SplitAfter(fileInfo.Name(), ".")
			ext := strings.ToLower(
				splittedFileName[len(splittedFileName)-1],
			)

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
