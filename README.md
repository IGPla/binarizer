# Binarizer

Binarizer is a Go project to create binary packages from scripts. 

One of the main concerns when using scripting languages is that all knowledge is delivered as is, in plain text. Packaging solutions often involve many steps, 3rd party libraries and confusing configurations for a basic goal: obfuscate code and deliver it in a native OS wrapper.

Binarizer will accomplish it with a simple yet powerful approach: wrap a given script into a go binary, ready to be used in the target platform

## Requirements

This project is already builded in Releases (see releases section to download a binary compilation). 

However, to be able to work, it requires go development environment to run (it compiles go binaries with scripts embedded, so it requires go environment to do that compilation). To install it in your system, please go to official Go install page

https://golang.org/doc/install

Note: Binarizer project requires go dev environment to compile final binaries. That final binaries DO NOT require that environment, and they are safe to be run in the environment that you've compiled them for

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
