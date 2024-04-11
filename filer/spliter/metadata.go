package spliter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Metadata struct {
	OriginalFilename string `json:"original_filename"`
	PartSize         int64  `json:"part_size"`
	TotalParts       int64  `json:"total_parts"`
}

func SaveMetaData(filePath string, partSize, totalParts int64) error {
	mdata := Metadata{
		OriginalFilename: filepath.Base(filePath),
		PartSize:         partSize,
		TotalParts:       totalParts,
	}

	mbs, err := json.Marshal(mdata)
	if err != nil {
		return err
	}

	metadataFilename := filePath + ".metadata.json"

	if err = os.WriteFile(metadataFilename, mbs, 0644); err != nil {
		return err
	}

	return nil
}

func LoadMetadata(filePath string) (*Metadata, error) {
	metadataFileName := filePath + ".metadata.json"

	mbs, err := os.ReadFile(metadataFileName)
	if err != nil {
		return nil, err
	}

	var metadata Metadata

	if err = json.Unmarshal(mbs, &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil

}
