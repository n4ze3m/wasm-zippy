package zippy

import (
	"archive/tar"
	"bytes"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/n4ze3m/zippy/utils"
)

func createTarFile(zipFiles []ZipFile, extension string) {
	b := new(bytes.Buffer)
	w := tar.NewWriter(b)
	for _, file := range zipFiles {
		fmt.Println("adding", file.Name)
		btye, err := base64.StdEncoding.DecodeString(file.Data)
		err = w.WriteHeader(&tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(btye)),
		})
		if err != nil {
			fmt.Println("error:", err)
			utils.ErrorMessage(err.Error())
			return
		}
		fmt.Println("writing", len(file.Data))
		if err != nil {
			fmt.Println("error:", err)
			utils.ErrorMessage(err.Error())
			return
		}
		w.Write(btye)

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

	fileName := "zippy-" + time.Now().Format("2006-01-02-15-04-05") + "." + extension
	utils.GenerateDownloadButton(fileName, data, "x-tar")
}
