package main
//
// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"runtime"
// 	"strings"
//
// 	"github.com/dustin/go-humanize"
// )
//
// func main() {
// 	// check OS
// 	if runtime.GOOS == "windows" {
// 		fmt.Println("Hello from Windows")
// 		installWindows()
//
// 	}
// 	if runtime.GOOS == "darwin" {
// 		fmt.Println("Hello from darwin")
// 		installUnix()
// 	}
// 	if runtime.GOOS == "linux" {
// 		fmt.Println("Hello from linux")
// 		installUnix()
// 	}
// }
//
// func installUnix() {
// 	// https://gofi.sh/index.html#install
//
// 	// Code:
// 	// curl -fsSL https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.sh | bash
//
// 	fmt.Println("Download Started")
//
// 	fileURL := "https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.sh"
// 	err := DownloadFile("install.sh", fileURL)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println("Download Finished")
//
// 	// install it
// 	fmt.Println("Install Started")
// 	cmd := exec.Command("/bin/sh", "install.sh")
// 	err = cmd.Run()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Install Finished")
//
// 	fmt.Println("Now run ```gofish init``` ")
//
// }
//
// func installWindows() {
// 	// https://gofi.sh/index.html#install
// 	// evoke From powershell as Administrator
//
// 	// Code:
// 	// Set-ExecutionPolicy Bypass -Scope Process -Force
// 	// iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.ps1'))
//
// 	fmt.Println("Download Started")
//
// 	fileURL := "https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.ps1"
// 	err := DownloadFile("install.ps1", fileURL)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println("Download Finished")
//
// 	// install it
// 	fmt.Println("Install Started")
//
// 	cmd := exec.Command("PowerShell", "-Command", "install.ps1")
// 	err = cmd.Run()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println("Install Finished")
//
// 	fmt.Println("Now run ``` gofish init ``` ")
//
// }
//
// /// DownloadFile stuff
//
// // DownloadFile will download a url to a local file. It's efficient because it will
// // write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// // into Copy() to report progress on the download.
// func DownloadFile(filepath string, url string) error {
//
// 	// Create the file, but give it a tmp file extension, this means we won't overwrite a
// 	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
// 	out, err := os.Create(filepath + ".tmp")
// 	if err != nil {
// 		return err
// 	}
//
// 	// Get the data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		out.Close()
// 		return err
// 	}
// 	defer resp.Body.Close()
//
// 	// Create our progress reporter and pass it to be used alongside our writer
// 	counter := &WriteCounter{}
// 	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
// 		out.Close()
// 		return err
// 	}
//
// 	// The progress use the same line so print a new line once it's finished downloading
// 	fmt.Print("\n")
//
// 	// Close the file without defer so it can happen before Rename()
// 	out.Close()
//
// 	if err = os.Rename(filepath+".tmp", filepath); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// // WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// // and we can pass this into io.TeeReader() which will report progress on each write cycle.
// type WriteCounter struct {
// 	Total uint64
// }
//
// func (wc *WriteCounter) Write(p []byte) (int, error) {
// 	n := len(p)
// 	wc.Total += uint64(n)
// 	wc.PrintProgress()
// 	return n, nil
// }
//
// // PrintProgress outputs the progress of the download.
// func (wc WriteCounter) PrintProgress() {
// 	// Clear the line by using a character return to go back to the start and remove
// 	// the remaining characters by filling it with spaces
// 	fmt.Printf("\r%s", strings.Repeat(" ", 35))
//
// 	// Return again and print current status of download
// 	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
// 	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
// }
