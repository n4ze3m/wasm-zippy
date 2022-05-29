package zippy

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"syscall/js"
	"time"

	"github.com/n4ze3m/zippy/utils"
)

type ZipFile struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func createZipfile(input string) {
	var zipFiles []ZipFile
	err := json.Unmarshal([]byte(input), &zipFiles)
	if err != nil {
		fmt.Println("error:", err)
		utils.ErrorMessage(err.Error())
		return
	}
	fmt.Println("we have", len(zipFiles), "files to zip")
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

	err = w.Close()
	if err != nil {
		fmt.Println("error:", err)
		utils.ErrorMessage(err.Error())
		return
	}

	data := b.Bytes()
	fmt.Println("we have", len(data), "bytes to send")

	fileName := "zippy-" + time.Now().Format("2006-01-02-15-04-05") + ".zip"
	zipFile := js.Global().Get("document").
		Call("createElement", "a")
	zipFile.Set("href", "data:application/zip;base64,"+base64.StdEncoding.EncodeToString(data))
	zipFile.Set("download", fileName)
	zipFile.Set("innerHTML", fmt.Sprintf("%s &nbsp; %s", `<i class="fas fa-file-archive"></i>`, fileName))
	zipFile.Set("className", "button is-success is-fullwidth mb-3")
	btnHub := js.Global().Get("document").
		Call("getElementById", "btnHub")
	btnHub.Call("appendChild", zipFile)
}

func CreateZipFile() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		input := args[0].String()
		createZipfile(input)
		return nil
	})
}
