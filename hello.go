package main

import "fmt" 

func main() {

    fmt.Printf( "\x1bc\x1b[40;37m\n %s\n",retstr())

}

func retstr() string{
    return "hello world"
}