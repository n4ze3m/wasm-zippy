package zippy

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/n4ze3m/zippy/utils"
)

type ZipFile struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func createArchivefile(input, output string) {
	var zipFiles []ZipFile
	err := json.Unmarshal([]byte(input), &zipFiles)
	if err != nil {
		fmt.Println("error:", err)
		utils.ErrorMessage(err.Error())
		return
	}
	fmt.Println("we have", len(zipFiles), "files to archive")

	if output == "zip" {
		createZipfile(zipFiles)
	} else {
		createTarFile(zipFiles, output)
	}
}

func CreateArchiveFile() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		input := args[0].String()
		output := args[1].String()
		createArchivefile(input, output)
		return nil
	})
}
