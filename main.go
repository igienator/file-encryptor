package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/igienator/file-encryptor/controllers"
)

func main() {

	errTwoOpsAtOnce := errors.New("cant perform encryption and decryption operation at once")

	flagEncrypt := flag.Bool("encrypt", false, "file encryption")
	flagDecrypt := flag.Bool("decrypt", false, "file decryption")
	flagFilename := flag.String("filename", "", "file name to be encrypted/decrypted")
	flagFileoutput := flag.String("output", "", "file name to be encrypted/decrypted")
	flagPassphrase := flag.String("key", "key.txt", "keyfile to encrypt/decrypt files")
	flag.Parse()

	if *flagEncrypt && *flagDecrypt {
		panic(errTwoOpsAtOnce)
	}

	if *flagEncrypt {
		if err := controllers.Encrypt(*flagFilename, *flagFileoutput, *flagPassphrase); err != nil {
			panic(err)
		}
		fmt.Println("File encrypted successfully")
	}

	if *flagDecrypt {
		if err := controllers.Decrypt(*flagFilename, *flagFileoutput, *flagPassphrase); err != nil {
			panic(err)
		}
		fmt.Println("File decrypted successfully")
	}
	fmt.Println(*flagEncrypt, *flagDecrypt, *flagPassphrase, *flagFilename, *flagFileoutput)
}
