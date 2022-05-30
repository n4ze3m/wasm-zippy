package zippy

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/n4ze3m/zippy/utils"
)

func createZipfile(zipFiles []ZipFile) {
	b := new(bytes.Buffer)
	w := zip.NewWriter(b)
	for _, file := range zipFiles {
		fmt.Println("adding", file.Name)
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println("error:", err)
			utils.ErrorMessage(err.Error())
			return
		}
		fmt.Println("writing", len(file.Data))
		btye, err := base64.StdEncoding.DecodeString(file.Data)
		if err != nil {
			fmt.Println("error:", err)
			utils.ErrorMessage(err.Error())
			return
		}
		f.Write(btye)

		if err != nil {
			fmt.Println("error:", err)
			utils.ErrorMessage(err.Error())
			return
		}
	}

	err := w.Close()
	if err != nil {
		fmt.Println("error:", err)
		utils.ErrorMessage(err.Error())
		return
	}

	data := b.Bytes()
	fmt.Println("we have", len(data), "bytes to send")

	fileName := "zippy-" + time.Now().Format("2006-01-02-15-04-05") + ".zip"
	utils.GenerateDownloadButton(fileName, data, "zip")
}
