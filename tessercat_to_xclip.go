package main

import (
	"context"
	_ "image/png"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/otiai10/gosseract/v2"
	"golang.design/x/clipboard"
)

func main() {
	lang := "eng"
	if len(os.Args) > 1 {
		lang = os.Args[1]
	}

	tempDir := getTempDir()
	imgPath := filepath.Join(tempDir, "tesseract.png")

	client, err := setupOCRClient(lang)
	if err != nil {
		logError("Failed to set up OCR client", err)
		os.Exit(1)
	}
	defer client.Close()

	watchClipboard(client, imgPath)
}

func getTempDir() string {
	if runtime.GOOS == "darwin" {
		// On macOS, use /private/tmp explicitly instead of /tmp
		return "/private/tmp"
	}
	return os.TempDir()
}

func setupOCRClient(lang string) (*gosseract.Client, error) {
	client := gosseract.NewClient()
	err := client.SetLanguage(lang)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func watchClipboard(client *gosseract.Client, imgPath string) {
	ctx := context.Background()
	changes := clipboard.Watch(ctx, clipboard.FmtImage)

	for {
		select {
		case <-changes:
			processClipboardImage(client, imgPath)
		}
	}
}

func processClipboardImage(client *gosseract.Client, imgPath string) {
	imgData := clipboard.Read(clipboard.FmtImage)
	if len(imgData) == 0 {
		logError("Empty image data from clipboard", nil)
		return
	}

	err := os.WriteFile(imgPath, imgData, 0644)
	if err != nil {
		logError("Failed to write image to file", err)
		return
	}

	time.Sleep(500 * time.Millisecond)

	err = client.SetImage(imgPath)
	if err != nil {
		logError("Failed to set image for OCR", err)
		return
	}

	text, err := client.Text()
	if err != nil {
		logError("OCR text extraction failed", err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(text))
}

func logError(message string, err error) {
	if err != nil {
		println(message+":", err.Error())
	} else {
		println(message)
	}
}
