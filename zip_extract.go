package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Extract(zipFile *zip.Reader, targetDirectory string) {
	for _, file := range zipFile.File {
		srcFile, _ := file.Open()
		outFilename := filepath.Join(targetDirectory, file.Name)

		outFile, _ := os.Create(outFilename)
		io.Copy(outFile, srcFile)
	}
}

func Example() {
	zipFile, _ := zip.OpenReader("testdata/example.zip")
	Extract(&zipFile.Reader, "/tmp/uploads")
}
