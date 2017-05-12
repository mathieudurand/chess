package main

import (
	"fmt"
	"flag"
	"io/ioutil"
)

type loc struct{
    X, Y string
}

// func will be used to "setup" a board. use as needed
func buildboard (pieces string)(map[string]loc) {
    fmt.Println("board is setting up")

    var m = make(map[string]loc)

    m["wR"] = loc{"a","1"}
    m["wK"] = loc{"a","5"}
    fmt.Println("testing where is wK:", m["wK"])
    fmt.Println("testing where is wR:", m["wR"])
    return m
}

// this func is called to test a possible move (maybe against a possible CHECK)
func checkMove(){
    fmt.Println("starting checkMove")
}

// func validate move list is valid

// func needed to check if move valid

// func needed to evaluate total value capture pieces in a branch of moves

// main user loop function
func launch(pieces string){
    fmt.Println(pieces)

    var initBoard = buildboard(pieces)
    fmt.Println(initBoard)
    fmt.Println("type \"quit\" when finished")

    var i int = 0
    for i < 1 {
        fmt.Print("next move: ")
        var input string
        fmt.Scanln(&input)
        fmt.Println(input)
        if input == "quit"{
            fmt.Println("bye bye")
            i=1
        }
    }

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
    fileStringPtr := flag.String("filename","setpieces.txt","A file with a list of chess pieces in it")
    flag.Parse()

    fmt.Println("Verbose mode: ", *boolPtr)
    fmt.Println("movesAhead: ", *movesAheadPtr)
    fmt.Println("File with setup: ", *fileStringPtr)

    dat, err := ioutil.ReadFile("./moves.txt")
    check(err)

    launch(string(dat))
}
