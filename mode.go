package main

// Mode specifies the mode in which you want to in,
// either encrypt or decrypt.
type Mode int

// EncryptMode is used to indicate you would like to encrypt
// DecryptMode is used to specify you would like to decrypt
const (
	EncryptMode Mode = iota
	DecryptMode
)

func (m Mode) String() string {
	return [...]string{"EncryptMode", "DecryptMode"}[m]
}
