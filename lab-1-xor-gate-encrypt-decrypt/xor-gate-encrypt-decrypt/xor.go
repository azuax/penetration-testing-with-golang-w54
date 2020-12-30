package main

// not so secure to have it on code!
const key = "5678"

// Encrypt XOR encryption
func Encrypt(s string) string {
	var encSlice []int
	sASCII := toASCII(s)
	keyASCII := toASCII(key)
	for i, e := range sASCII {
		encSlice = append(encSlice, e^keyASCII[i%len(keyASCII)])
	}
	return toString(encSlice)
}

// Decrypt XOR decryption
func Decrypt(s string) string {
	var decSlice []int
	sASCII := toASCII(s)
	keyASCII := toASCII(key)
	for i, e := range sASCII {
		decSlice = append(decSlice, e^keyASCII[i%len(keyASCII)])
	}
	return toString(decSlice)
}

func toASCII(s string) []int {
	var asciiVals []int
	for _, e := range s {
		asciiVals = append(asciiVals, int(e))
	}
	return asciiVals
}

func toString(is []int) string {
	var st string
	for _, s := range is {
		st += string(byte(s))
	}
	return st
}
