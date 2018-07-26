package main 

import (
	"fmt"
	"io.Reader"
	"strings"
	"os"
)

type alphaReader struct{
	reader io.Reader 
}

func NewAlphaReader(reader io.Reader) *alphaReader{
	return &(alphaReader{reader: reader})
}

func alpha(char byte) byte{
	if(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'){
		return char 
	}
	return 0
}

func(a *alphaReader) Read(p []byte) (int, error){
	n, err := a.reader.Read(p)
	if(err != nil){
		return 0, err
	}

	buffer := new([]byte, 4)
	
	for i := 0; i < n; i++ {
		if char = alpha(p[i]); char != 0{
			buffer[i] = char 
		}
	}

	copy(buffer, p)
	return 0, nil 
}

func main(){
	reader := NewAlphaReader(strings.NewReader("Hey there, Hope you are doing great in your job!!"))
	p := make([]byte, 4)

	for{
		n, err := reader.Read(p)
		if(err == io.EOF){
			break
		}
		fmt.Println(p[:n])
	}
	fmt.Println()
}