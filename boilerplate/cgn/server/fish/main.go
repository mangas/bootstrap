package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/dustin/go-humanize"

	"github.com/KnicKnic/go-powershell/pkg/logger/kloghelper"
	"github.com/KnicKnic/go-powershell/pkg/powershell"
)

func main() {
	// check OS
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
		installWindows()

	}
	if runtime.GOOS == "darwin" {
		fmt.Println("Hello from darwin")
		installUnix()
	}
	if runtime.GOOS == "linux" {
		fmt.Println("Hello from linux")
		installUnix()
	}
}

func installUnix() {
	// https://gofi.sh/index.html#install

	// Code:
	// curl -fsSL https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.sh | bash

	fmt.Println("Download Started")

	fileURL := "https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.sh"
	err := DownloadFile("install.sh", fileURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")

	// install it
	fmt.Println("Install Started")
	cmd := exec.Command("/bin/sh", "install.sh")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Install Finished")

	fmt.Println("Now run ```gofish init``` ")

}

func installWindows() {
	// https://gofi.sh/index.html#install
	// evoke From powershell as Administrator

	// Code:
	// Set-ExecutionPolicy Bypass -Scope Process -Force
	// iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.ps1'))

	fmt.Println("Download Started")

	fileURL := "https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.ps1"
	err := DownloadFile("install.ps1", fileURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")

	// install it
	fmt.Println("Install Started")

	runspace := powershell.CreateRunspace(kloghelper.Klog{VerboseLevel: 1, DebugLevel: 2}, callbackTest{})
	defer runspace.Close()

	PrintAndExecuteCommand(runspace, "install.ps1", *useLocalScope)

	// Just tests we can talk to Powershell and the output encoding is not ok.
	fmt.Println("With encoding change:")
	stdout, stderr, err := posh.Execute("install.ps1")
	fmt.Println(stdout)
	fmt.Println(stderr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Install Finished")

	fmt.Println("Now run ``` gofish init ``` ")

}

/// DownloadFile stuff

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

// PrintProgress outputs the progress of the download.
func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

/// Powershell stuff

type callbackTest struct{}

func (c callbackTest) Callback(_ powershell.Runspace, str string, input []powershell.Object, results powershell.CallbackResultsWriter) {
	fmt.Println("\tIn callback:", str)
	results.WriteString(str)
	for i, object := range input {
		if object.IsNull() {
			fmt.Println("\tIn callback: index", i, "Object Is Null") // ToString and Type are still valid
		}
		fmt.Println("\tIn callback: index", i, "type:", object.Type(), "with value:", object.ToString())
		results.Write(object, false)
	}
}

// PrintAndExecuteCommand executes a command in powershell and prints the results
func PrintAndExecuteCommand(runspace powershell.Runspace, command string, useLocalScope bool) {
	fmt.Println("Executing powershell command:", command)

	// determine if executing just a .ps1 file, if so use command, otherwise script
	var results *powershell.InvokeResults
	if strings.HasSuffix(command, ".ps1") {
		results = runspace.ExecCommand(command, useLocalScope, nil)
	} else {
		results = runspace.ExecScript(command, useLocalScope, nil)
	}
	defer results.Close()

	fmt.Println("Completed Executing powershell command:", command)
	if !results.Success() {
		fmt.Println("\tCommand threw exception of type", results.Exception.Type(), "and ToString", results.Exception.ToString())
	} else {
		fmt.Println("Command returned", len(results.Objects), "objects")
		for i, object := range results.Objects {
			fmt.Println("\tObject", i, "is of type", object.Type(), "and ToString", object.ToString())
		}
	}
}
