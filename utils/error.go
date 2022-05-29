package utils

import (
	"fmt"
	"syscall/js"
)

func ErrorMessage(message string) {
	error := fmt.Sprintf(`<article class="message mb-3  is-danger"><div class="message-header"><p>%s</p></div></article>`, message)
	elem := js.Global().Get("document").Call("getElementById", "status")
	elem.Set("innerHTML",error)
}
