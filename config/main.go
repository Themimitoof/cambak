package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type DestinationConflictType string

const (
	DEST_CONFLICT_SKIP  DestinationConflictType = "skip"
	DEST_CONFLICT_MERGE DestinationConflictType = "merge"
)

type Configuration struct {
	Extract ExtractionConfiguration `yaml:"extract"`
}

type ExtractionConfiguration struct {
	ExtractPictures bool `yaml:"pictures"`
	ExtractRaws     bool `yaml:"raws"`
	ExtractMovies   bool `yaml:"movies"`

	SourcePath string `yaml:",omitempty"`

	DestinationPath     string                  `yaml:"destination"`
	DestinationFormat   string                  `yaml:"format"`
	DestinationConflict DestinationConflictType `yaml:"conflict"`

	CameraName     string `yaml:"camera_name"`
	CleanAfterCopy bool   `yaml:"clean_after_copy"`

	DryRunMode bool `yaml:",omitempty"`
}

func NewConfiguration() Configuration {
	return Configuration{
		Extract: ExtractionConfiguration{
			ExtractPictures: true,
			ExtractRaws:     true,
			ExtractMovies:   true,

			DestinationFormat:   "%y/%m-%d/%n/%t",
			DestinationConflict: DEST_CONFLICT_SKIP,

			CameraName:     "Camera",
			CleanAfterCopy: false,
		},
	}
}

func NewConfigurationFile(path string) error {
	conf := NewConfiguration()
	strConf, marshalErr := yaml.Marshal(conf)

	if marshalErr != nil {
		return marshalErr
	}

	err := ioutil.WriteFile(path, strConf, 0755)

	if err != nil {
		return err
	}

	return nil
}

func OpenConfigurationFile(path string) (Configuration, error) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return Configuration{}, err
	}

	conf := NewConfiguration()
	err = yaml.Unmarshal(file, &conf)

	if err != nil {
		return Configuration{}, err
	}
	return conf, nil
}
