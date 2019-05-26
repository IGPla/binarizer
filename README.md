# Binarizer

Binarizer is a Go project to create binary packages from scripts. 

One of the main concerns when using scripting languages is that all knowledge is delivered as is, in plain text. Packaging solutions often involve many steps, 3rd party libraries and confusing configurations for a basic goal: obfuscate code and deliver it in a native OS wrapper.

Binarizer will accomplish it with a simple yet powerful approach: wrap a given script into a go binary, ready to be used in the target platform

## Configuration options

- -script-type: The scripting language selected. Available values are {python, bash}
- -script-file: path to your script file (default ./myscript.sh)
- -binary-file: path to the new binary that will be created (default ./mybinary)
- -goos: OS to be compiled against (default linux)
- -goarch: Architecture to be compiled against (default amd64)
- -encrypt: Encrypt (or not) your binary. It will prevent anyone to, even if they decompile the binary, to see the code that will be executed.

## Build

```
go build -o /usr/local/bin/binarizer
```

## Usage

Basic binary

```
binarizer -script-type bash -script-file myfile.sh -binary-path /tmp/mybinary
/tmp/mybinary
```

Encrypted basic binary
```
binarizer -script-type bash -script-file myfile.sh -binary-path /tmp/mybinary -encrypt
BIN_KEY=<key output from previous step> /tmp/mybinary
```
