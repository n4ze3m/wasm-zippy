package utils

import (
	"encoding/base64"
	"fmt"
	"syscall/js"
)

func GenerateDownloadButton(fileName string, data []byte, mimeType string) {
	zipFile := js.Global().Get("document").
		Call("createElement", "a")
	zipFile.Set("href", fmt.Sprintf(`"data:application/%s;base64,%s"`, mimeType, base64.StdEncoding.EncodeToString(data)))
	zipFile.Set("download", fileName)
	zipFile.Set("innerHTML", fmt.Sprintf("%s &nbsp; %s", `<i class="fas fa-file-archive"></i>`, fileName))
	zipFile.Set("className", "button is-success is-fullwidth mb-3")
	btnHub := js.Global().Get("document").
		Call("getElementById", "btnHub")
	btnHub.Call("appendChild", zipFile)
}
