package quran

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	downloadError   = "download error: %s %d"
	minDownloadSize = 1024
)

func LoadTrans(trans string) error {
	resp, err := http.Get("http://tanzil.net/trans/" + trans)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error in downloading - got status", resp.StatusCode)
		return fmt.Errorf(downloadError, "unexpected status", resp.StatusCode)
	}

	out, err := os.Create(trans)
	if err != nil {
		fmt.Println("failed to create file for ", trans, err)
		return err
	}
	defer out.Close()

	n, err := io.Copy(out, resp.Body)

	// unfortunately tanzil.txt doesn't return non 200 response, even when download failed. so we check by size
	if n < minDownloadSize && err == nil {
		os.Remove(trans)
		return fmt.Errorf(downloadError, "unexpected size - too small ", n)
	}

	fmt.Println("downloaded ", n, "bytes")

	return err
}

// TODO: Enable loading of translation directly into sqlite DB
