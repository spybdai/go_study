package main

import (
	"bufio"
	"fmt"
	eh "github.com/spybdai/go_study/utils/error_handle"
	log "github.com/spybdai/go_study/utils/log"
	"os"
	"runtime"
	"strings"
)

func GetFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return fmt.Sprintf("%s:", f.Name())
}

// #############
// Scanner.Scan

// tip-1: need check Err() after scan

var Delim = "#####"

func ScanWithoutCheckErr(path string) {
	f, err := os.Open(path)
	eh.PanicError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineNum := strings.Split(line, Delim)[0]
		fmt.Println(log.GetFuncName(), lineNum)
	}
	// no check error, cannot find the error
}

func ScanWithCheckErr(path string) {
	f, err := os.Open(path)
	eh.PanicError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineNum := strings.Split(line, Delim)[0]
		fmt.Println(log.GetFuncName(), lineNum)
	}

	err = scanner.Err()
	eh.PanicError(err)
}

// tip-1.1: why the "too long" error?
// there is attribute named maxTokenSize which will control the buf size

// tip-1.2: how to resolve?

func UseBufferBeforeScan(path string) {
	f, err := os.Open(path)
	eh.PanicError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	//reset the buffer and the MaxScanTokenSize
	scanner.Buffer(make([]byte, 1), bufio.MaxScanTokenSize*1000)

	for scanner.Scan() {
		line := scanner.Text()
		lineNum := strings.Split(line, Delim)[0]
		fmt.Println(log.GetFuncName(), lineNum)
	}

	err = scanner.Err()
	eh.PanicError(err)

}

// or

func UseReadString(path string) {
	f, err := os.Open(path)
	eh.PanicError(err)
	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err != nil { // nee check this. especially if no '\n' at end of the last line
			fmt.Println(err)
			break
		}

		lineNum := strings.Split(line, Delim)[0]
		fmt.Println(log.GetFuncName(), lineNum)
	}
}

func main() {
	longLineFile := "long_line_file.txt"

	ScanWithoutCheckErr(longLineFile)
	//ScanWithCheckErr(longLineFile)
	UseBufferBeforeScan(longLineFile)
	UseReadString(longLineFile)

	longLineFileWithNewLine := "long_line_file_with_newline.txt"
	UseReadString(longLineFileWithNewLine)
}
