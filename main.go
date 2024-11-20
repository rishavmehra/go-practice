package main

// I use this for small-small tests, for better understanding
import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func main() {
	// numCPUs := runtime.NumCPU()
	// fmt.Println(numCPUs)

	type version struct {
		Ver        int
		BestHeight int
		AddrFrom   string
	}

	result := gobEncode(version{0, 0, "rishav"})
	// request := append(CommandToBytes("version"), result...)
	fmt.Printf("%v", result)

}

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()

}

// func CommandToBytes(command string) []byte {
// 	var bytes [12]byte

// 	for i, c := range command {
// 		bytes[i] = byte(c)
// 	}
// 	return bytes[:]

// }
