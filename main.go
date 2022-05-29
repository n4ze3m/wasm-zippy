package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"syscall/js"
	"time"
)

type ZipFile struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func createZipfile(input string) {
	// json unmarshal
	var zipFiles []ZipFile
	err := json.Unmarshal([]byte(input), &zipFiles)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("we have", len(zipFiles), "files to zip")
	b := new(bytes.Buffer)
	w := zip.NewWriter(b)
	for _, file := range zipFiles {
		fmt.Println("adding", file.Name)
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println("writing", len(file.Data))
		// convert array of strings to bytes
		// convert base64 string to bytes
		// write bytes to zip file
		// decode base64 string to bytes
		btye, err := base64.StdEncoding.DecodeString(file.Data)
		if err != nil {
			fmt.Println("error:", err)
		}
		f.Write(btye)

		if err != nil {
			fmt.Println("error:", err)
		}
	}

	err = w.Close()
	if err != nil {
		fmt.Println("error:", err)
	}

	data := b.Bytes()
	fmt.Println("we have", len(data), "bytes to send")

	fileName := "zippy-" + time.Now().Format("2006-01-02-15-04-05") + ".zip"
	zipFile := js.Global().Get("document").
		Call("createElement", "a")
	zipFile.Set("href", "data:application/zip;base64,"+base64.StdEncoding.EncodeToString(data))
	zipFile.Set("download", fileName)
	zipFile.Set("innerHTML", fmt.Sprintf("%s &nbsp; %s", `<i class="fas fa-file-archive"></i>`, fileName))
	// add class to zip file
	zipFile.Set("className", "button is-success is-fullwidth mb-3")
	btnHub := js.Global().Get("document").
		Call("getElementById", "btnHub")
	btnHub.Call("appendChild", zipFile)
}

func wasmWrapper() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		input := args[0].String()
		createZipfile(input)
		return nil
	})
}

func main() {

	js.Global().Set("Zippy", wasmWrapper())
	<-make(chan bool)
}
