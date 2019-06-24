package main

import (
	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

func main() {
	user32.MessageBox(0, "send mail", "test title", user32.MB_OK)
}
