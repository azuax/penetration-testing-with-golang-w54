package main

import "fmt"

const menu string = `1	Encrypt data
2	Decrypt data
3	Exit`

func main() {
	fmt.Println("XOR Get encrypt/decrypt")
	fmt.Println("Choose what do you want to do:")

	var opt int
	for opt != 3 {
		fmt.Println(menu)
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1:
			fmt.Println("Enter the data to encrypt:")
			var data string
			if _, e := fmt.Scanln(&data); e != nil {
				fmt.Println("There was an error with your input")
				break
			}
			fmt.Println(Encrypt(data))
		case 2:
			fmt.Println("Enter the data to decrypt:")
			var encData string
			if _, e := fmt.Scanln(&encData); e != nil {
				fmt.Println("There was an error with your input")
				break
			}
			fmt.Println(Decrypt(encData))
		}
	}

	fmt.Println("Bye!")
}
