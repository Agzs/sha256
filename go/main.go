// main.go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Bytes2Hex returns the hexadecimal encoding of d.
func Bytes2Hex(d []byte) string {
	return hex.EncodeToString(d)
}

// Deprecated: use hexutil.Encode instead.
func ToHex(b []byte) string {
	hex := Bytes2Hex(b)
	if len(hex) == 0 {
		hex = "0"
	}
	return "0x" + hex
}

func main() {
	secretData := []byte{
		14, 0, 0, 0, 0, 0, 0, 0,
		18, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, //
		18, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0} //
	hashData := sha256.Sum256(secretData)

	fmt.Printf("secretData: ")
	fmt.Println(secretData)
	fmt.Printf("hashData: ")
	fmt.Println(hashData)
	fmt.Println(ToHex(hashData[:]))

	fmt.Println("===============")
	value := []byte{1, 0, 0, 0, 0, 0, 0, 0}
	sn := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	r := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	hasher := sha256.New()
	hasher.Write(value)
	hasher.Write(sn)
	hasher.Write(r)
	ID := []byte{}
	res := hasher.Sum(ID[:])
	fmt.Printf("hashData: ")
	fmt.Println(res)
	fmt.Println(ToHex(res[:]))
}
