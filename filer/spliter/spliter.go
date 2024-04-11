package spliter

import (
	"fmt"
	"io"
	"os"
)

// partSize is the size of each chunk in bytes
// represnted as a 64 bit integer

func SplitFile(filePath string, partSizeInBytes int64) error {
	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	totalParts := fileInfo.Size() / partSizeInBytes
	if fileInfo.Size()%partSizeInBytes != 0 {
		totalParts++
	}

	for i := int64(0); i < totalParts; i++ {
		unProcessed := min(partSizeInBytes, fileInfo.Size()-i*partSizeInBytes)
		outFile, err := os.Create(fmt.Sprintf("%s.part_%d", filePath, i+1))
		if err != nil {
			return err
		}

		_, err = io.CopyN(outFile, file, unProcessed)

		if err != nil {
			return err
		}

		if err = outFile.Close(); err != nil {
			return err
		}
	}

	fmt.Printf("file size -> %d, no of chunks -> %d", fileInfo.Size(), totalParts)

	// write meta data

	if err := SaveMetaData(filePath, partSizeInBytes, totalParts); err != nil {
		return err
	}

	// remove main file

	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil
}

func JoinFiles(filePath string) error {
	mdata, err := LoadMetadata(filePath)
	if err != nil {
		return err
	}

	outFile, err := os.Create(mdata.OriginalFilename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for i := int64(0); i < mdata.TotalParts; i++ {
		partFileName := fmt.Sprintf("%s.part_%d", filePath, i+1)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, partFile)
		if err != nil {
			return err
		}

		_ = os.Remove(partFileName)

		if err = partFile.Close(); err != nil {
			return err
		}
	}

	// remove metadata

	_ = os.Remove(filePath + ".metadata.json")

	return nil
}
