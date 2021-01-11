package main
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)
type Count struct {
	Size uint64
	Time int
}
func (cnt *Count) Write(p []byte) (int, error) {
	n := len(p)
	cnt.Size += uint64(n)

	return n, nil
}
func countTime(cnt *Count){
	for {
		time.Sleep(time.Second)
		cnt.Time++
		fmt.Println("Downloading for",cnt.Time,"sec, Downloaded", cnt.Size,"bytes already")
	}
}
func DownloadFile(URL string) error {
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	out, err := os.Create(path.Base(resp.Request.URL.String()))
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	fin := &Count{}
	go countTime(fin)
	if _, err = io.Copy(out, io.TeeReader(resp.Body, fin)); err != nil {
		return err
	}
	out.Close()
	fmt.Println("\nFinish\nFile size:",fin.Size,"bytes\nDownload time:",fin.Time,"sec")
	return nil
}
func main() {
	var fileURL string
	fmt.Println("Enter your link:")
	fmt.Scan(&fileURL)
	err := DownloadFile(fileURL)
	if err != nil {
		panic(err)
	}
}