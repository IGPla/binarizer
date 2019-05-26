/* Templates */
package binary

const commandTemplate = `package main
import (
	"os"
	"os/exec"
)
func Command(args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
func main() {
        data := ` + "`{{.code}}`" + `
	Command({{.command}}, string(data))
}`

const commandEncryptedTemplate = `package main
import (
	"os"
	"os/exec"
        "encoding/base64"
      	"crypto/aes"
	"crypto/cipher"
        "log"
        "errors"
)
func command(args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
func decodeKey(key string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		log.Fatalf("Could not decode key. %v", err)
	}
	return decoded
}
func decryptString(key []byte, securemess string) (decodedmess string, err error) {
	cipherText, err := base64.StdEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}
func main() {
        key := decodeKey(os.Getenv("BIN_KEY"))
        data, err := decryptString(key, ` + "`{{.code}}`" + `)
        if err != nil {
                log.Fatalf("Could not execute binary code. Please, be sure that you've provided BIN_KEY env var")
        }
	command({{.command}}, string(data))
}`
