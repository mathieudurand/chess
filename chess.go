package main

import (
	"fmt"
	"flag"
	"io/ioutil"
)

// main loop function
func launch(moves string){
    fmt.Print(moves)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// argument processing and sourcing move list
func main() {

    boolPtr := flag.Bool("verbose", false, "a bool")
    fileStringPtr := flag.String("filename","moves.txt","A file with a list of chess moves in it")
    flag.Parse()

    fmt.Println("Verbose mode: ", *boolPtr)
    fmt.Println("File with moves: ", *fileStringPtr)

    dat, err := ioutil.ReadFile("./moves.txt")
    check(err)

    launch(string(dat))
}
