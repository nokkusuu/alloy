package lib
// download.go
import (
	"os"
	"net/http"
	"io"
)

type Download struct {
	URL string
	Dest string
	Progress chan uint64
	ProgInt uint64
}

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

	resp, err := http.Get(down.URL)
	defer resp.Body.Close()
	if err != nil {
		out.Close()
		return err
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