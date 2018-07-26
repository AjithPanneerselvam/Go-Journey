/*
Pipe writers and readers
Types io.PipeWriter and io.PipeReader model IO operations as in memory pipes. 
Data is written to the pipe’s writer-end and is read on the pipe’s reader-end 
using separate go routines. The following creates the pipe reader/writer pair 
using the io.Pipe() which is then used to copy data from a buffer proverbs to io.Stdout
*/

package main 

import (
	"os"
	"bytes"
	"io"
)

func main(){
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	piperR, piperW := io.Pipe()

	go func(){
		defer piperW.Close()
		io.Copy(piperW, proverbs)
	}()

	io.Copy(os.Stdout, piperR)
	piperR.Close()
}