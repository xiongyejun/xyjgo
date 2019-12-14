package main

import (
	"github.com/Afrostream/afrostream-media-server/src/mp4"
	//        "log"
)

func main() {
	mp4.Debug(true)
	mp4.ParseFile("1.mp4", "eng")
}
