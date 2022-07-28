package main

import (
	"context"
	"github.com/otiai10/gosseract/v2"
	"golang.design/x/clipboard"
	_ "image/png"
	"io/ioutil"
	"os"
)

var filename string = "/tmp/tesseract.png"

func main() {

	lng := "eng"
	if len(os.Args) > 1 {
		lng = os.Args[1]
	}
	var client *gosseract.Client = gosseract.NewClient()
	defer func(client *gosseract.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	err := client.SetLanguage(lng)
	if err != nil {
		panic(err)
	}

	for {
		changed := clipboard.Watch(context.Background(), clipboard.FmtImage)
		select {
		case <-changed:

			d := clipboard.Read(clipboard.FmtImage)
			_ = ioutil.WriteFile(filename, d, 0644)

			if err != nil {
				println("fcu")
				panic(err)
			} else {
				readimage(client)
				text, _ := client.Text()
				clipboard.Write(clipboard.FmtText, []byte(text))
			}
		}
	}

}

func readimage(client *gosseract.Client) {

	err := client.SetImage(filename)
	if err != nil {
		panic(err)
	}
}
