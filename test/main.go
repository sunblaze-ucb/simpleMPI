package main

import (
	"fmt"
	"os"
	"simpleMPI/mpi"
)

func main() {
	fmt.Println("Hello, playground", os.Args)
	world := mpi.WorldInit("ip.txt", "id_ed25519", "farmer")
	fmt.Println(world)
	mpi.Close()
}