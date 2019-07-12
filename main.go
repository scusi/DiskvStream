// this is only a proof of concept for now.
// nothing really usable
package main

import (
  "flag"
  "bufio"
  "os"
  "io"
  "fmt"
  "log"
  "path/filepath"
  "github.com/peterbourgon/diskv"
)

var (
  fileName string
)

func init() {
  flag.StringVar(&fileName, "f", "", "file to read")
}

func main() {
  flag.Parse()
  s := diskv.New(diskv.Options{
    BasePath: "data-store",
  })

  // open the file specified by fileName
  f, err := os.Open(fileName)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  rdr := bufio.NewReader(f)
  key := filepath.Base(fileName)
// write file to data-store
  err = s.WriteStream(key, rdr, true)
  if err != nil {
    log.Fatal(err)
  }

// read file from data-storage
rdrc, err := s.ReadStream(key, true)
if err != nil {
  log.Fatal(err)
}
out, err := os.Create("outfile")
if err != nil {
  log.Fatal(err)
}
defer out.Close()
written, err := io.Copy(out, rdrc)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("%d bytes written to 'outfile'", written)
}
