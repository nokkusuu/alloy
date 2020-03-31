package lib
// download.go
import (
	"os"
	"net/http"
	"io"
	"strconv"
)

func (down *Download) Write(p []byte) (int, error) {
	n := len(p)
	down.ProgInt += uint64(n)
	down.Progress <- down.ProgInt
	return n, nil
}

func (down *Download) Do() error {
	out, err := os.Create(down.Dest + ".tmp")
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", down.URL, nil)
	req.Header.Add("User-Agent", UserAgent)

	resp, err := HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		out.Close()
		return err
	}

	v, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	down.TotalSize = uint64(v)
	if down.TotalSize == 0 {
		down.TotalSize = 1 // prevent divide by zero error
	}

	if _, err = io.Copy(out, io.TeeReader(resp.Body, down)); err != nil {
		out.Close()
		return err
	}
	out.Close()

	if err = os.Rename(down.Dest+".tmp", down.Dest); err != nil {
		return err
	}

	close(down.Progress)
	return nil
}

func NewDownload(url string, dest string) Download {
	return Download{
		URL: url,
		Dest: dest,
		Progress: make(chan uint64),
	}
}