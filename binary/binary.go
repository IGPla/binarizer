/* Binary generation and compile functions */
package binary

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

/*
Description: Create a new Go source file containing the script code
*/
func SourceMaker(scriptFile *string, scriptCall *string, encrypt *bool) (*string, string) {
	code, err := ioutil.ReadFile(*scriptFile)
	if err != nil {
		log.Fatalf("Error arised while reading script: %v", err)
	}

	var _template string = commandTemplate
	var _code string = string(code)
	var _key []byte
	if *encrypt {
		_template = commandEncryptedTemplate
		_key = generateKey()
		_code, err = encryptString(_key, _code)
		if err != nil {
			log.Fatalf("Error arised while encrypting code: %v", err)
		}
	}

	commandFields := strings.Fields(*scriptCall)
	for i, field := range commandFields {
		commandFields[i] = fmt.Sprintf("\"%v\"", field)
	}
	dynamicData := map[string]string{
		"command": strings.Join(commandFields, ", "),
		"code":    _code,
	}
	buildTemplate, err := template.New("binary").Parse(_template)
	if err != nil {
		log.Fatalf("Error arised while building template: %v", err)
	}
	rand.Seed(time.Now().UnixNano())
	sourceFile := fmt.Sprintf("/tmp/%v.go", rand.Int63())
	binaryFile, err := os.Create(sourceFile)
	if err != nil {
		log.Fatalf("Error arised while creating source file: %v", err)
	}
	defer binaryFile.Close()

	err = buildTemplate.Execute(binaryFile, dynamicData)
	if err != nil {
		log.Fatalf("Error arised while writing template: %v", err)
	}
	keyEncoded := ""
	if *encrypt {
		keyEncoded = encodeKey(_key)
	}
	return &sourceFile, keyEncoded
}

/*
Description: Compile new generated source file into a binary
*/
func CompileSource(sourceFile *string, binaryFile *string, goos *string, goarch *string) {
	cmdParams := []string{"go", "build", "-o", *binaryFile, *sourceFile}
	cmd := exec.Command(cmdParams[0], cmdParams[1:]...)
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GOOS=%v", *goos),
		fmt.Sprintf("GOARCH=%v", *goarch))
	err := cmd.Run()
	if err != nil {
		log.Fatalf("An error arised while compiling new source: %v", err)
	}
}
