package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("test.txt")
			defer f.Close()
			if err != nil {
				fmt.Println(err)
			} else {
				data, err := ioutil.ReadAll(f)
				if err != nil {
					fmt.Println("Read err: ", err)
				}
				fmt.Printf("File contain: %s\n", string(data))
			}
		}()
	}
}
