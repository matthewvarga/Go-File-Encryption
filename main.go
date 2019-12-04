package main

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"os"

	"fmt"
)

// takes a plaintext string, and uses SHA256 alg to hash it
// returns the 64 digit hex hash
func hashString(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func main() {

	// load flags
	passwordFlag, pathFlag, err := loadFlags()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Password: ", *passwordFlag)
	fmt.Println("Path: ", *pathFlag)

	// load curve params
	curve := elliptic.P256()

	// hash password string
	passHash := hashString(*passwordFlag)

	fmt.Println("Password Hash: ", passHash)

	// convert hash to hex byte string
	//hexByteHash := passHash.Sum(nil)
	// convert hex byte string to int string (64 digits)
	intStr := hex.EncodeToString(passHash)

	// convert int string to int bytes: []uint8
	intBytes := []byte(intStr)

	fmt.Println("int Str: ", intStr)

	// calculate the point P = kG, where k = hashed value of password
	secretX, secretY := curve.ScalarMult(curve.Params().Gx, curve.Params().Gy, intBytes)

	fmt.Println("\nX: ", secretX)
	fmt.Println("Y: ", secretY)

	// concnat x,y secret coords into single number, and convert to byte array
	secret := secretX.String() + "" + secretY.String()
	secretBytesSlice := []byte(secret)
	var secretBytes [32]byte

	copy(secretBytes[:], secretBytesSlice)
	fmt.Println("secret bytes: ", secretBytes)

	plainText := []byte("Hello World!")

	encryptedMsg, _ := Encrypt(plainText, &secretBytes)

	decryptedMsg, _ := Decrypt(encryptedMsg, &secretBytes)

	fmt.Println("Encrypted Message: ")
	fmt.Printf("%x", encryptedMsg)

	fmt.Println("\nDecrypted Message: ", string(decryptedMsg))

}
