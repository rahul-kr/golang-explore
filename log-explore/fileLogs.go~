package main

import (
  "io"
  "io/ioutil"
  "log"
  "os"
  "fmt"
)

var (
    Error   *log.Logger
)

func Init(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
    // set location of log file
    // To write logs in file
    var errorLogpath = "/tmp/error.log"
    createFile(errorLogpath)
    var errorFile, err = os.OpenFile(errorLogpath, os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
      panic(err)
    }
    Error = log.New(errorFile,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}

func createFile(path string) {
    // detect if file exists
    var _, err = os.Stat(path)
    // create file if not exists
    if os.IsNotExist(err) {
	var file, err = os.Create(path)
	if isError(err) { return }
	defer file.Close()
    } 	
}

func isError(err error) bool {
    if err != nil {
	fmt.Println(err.Error())
    }
    return (err != nil)
}

func main() {
    Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
    Error.Println("Test for writing logs in File.")
}
