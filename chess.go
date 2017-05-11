package main

import (
	"fmt"
	"flag"
	"io/ioutil"
)

// build class for board
// class will be used to "setup" the move list requested and propose a move to follow with

    // func validate move list is valid

    // func needed to check if move valid


// func needed to evaluate total value capture pieces in a branch of moves

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

    boolPtr := flag.Bool("verbose", false, "Verbose mode")
    movesAheadPtr := flag.Int("movesAhead", 3, "Number of moves to think ahead on.")
    fileStringPtr := flag.String("filename","moves.txt","A file with a list of chess moves in it")
    flag.Parse()

    fmt.Println("Verbose mode: ", *boolPtr)
    fmt.Println("movesAhead: ", *movesAheadPtr)
    fmt.Println("File with moves: ", *fileStringPtr)

    dat, err := ioutil.ReadFile("./moves.txt")
    check(err)

    launch(string(dat))
}
