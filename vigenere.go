package main

/*****************************
** Name: Nisarga Patel
** Class: Operating Systems
** Assignment: Lab
**
*****************************/

import "os"
import "bufio"
import "fmt"
import "strings"
import "io/ioutil"

// check if alphabetical character
func isAlpha(plainChar byte) bool {

	if plainChar >= 65 && plainChar <= 90 {
		return true
	}
	return false

}

func getEncryptedChar(plainChar byte, keyChar byte) string {

	encryptedChar := ((plainChar - 65) + (keyChar - 65)) % 26

	return string(encryptedChar + 65)
}

func getDecryptedChar(cipherChar int, keyChar int) string {

	decryptedChar := (cipherChar - keyChar) % 26

	for decryptedChar < 0 {
		decryptedChar += 26
	}

	return string(decryptedChar + 65)
}

func vigenere(key string, inputtext string, encrypt bool) string {

	resultText := ""
	keyIndex := 0
	count := 0

	// sanitize the input
	keyUpper := strings.ToUpper(key)
	inputUpper := strings.ToUpper(inputtext)

	for i := 0; i < len(inputUpper); i++ {

		if isAlpha(inputUpper[i]) {

			keyIndex = count % len(keyUpper)

			if encrypt == true {
				// encrypt
				resultText += getEncryptedChar(inputUpper[i], keyUpper[keyIndex])
			} else {
				// decrypt
				resultText += getDecryptedChar(int(inputUpper[i]), int(keyUpper[keyIndex]))
			}
			count++

		} else {

			resultText += string(inputUpper[i])

		}

	}

	return resultText

}

func main() {

	if len(os.Args) == 3 {

		key := os.Args[1]

		plaintext := os.Args[2]

		ciphertext := vigenere(key, plaintext, true)
		newplaintext := vigenere(key, ciphertext, false)
		fmt.Println("Ciphertext: " + ciphertext)
		fmt.Println("Plaintext: " + newplaintext)
	} else {

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter a key: ")

		key, _ := reader.ReadString('\n')
		key = strings.Trim(key, "\n")

		fmt.Println("Your key is: " + key)

		fmt.Println("Enter a filename: ")
		filename, _ := reader.ReadString('\n')
		filename = strings.Trim(filename, "\n")

		pwd, _ := os.Getwd()
		fileContents, err := ioutil.ReadFile(pwd + "/" + filename)

		if err != nil {
			fmt.Println(err)
		}

		ciphertext := vigenere(key, string(fileContents), true)
		newplaintext := vigenere(key, ciphertext, false)

		fmt.Println("-----------CIPHERTEXT-----------")
		fmt.Println(ciphertext)
		fmt.Println("-----------PLAINTEXT-----------")
		fmt.Println(newplaintext)

	}

}
