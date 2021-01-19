package transaction

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func Binary(binaryPath string, binaryargs []string) Transaction {
	return func(context interface{}) error {
		command := exec.Command(binaryPath, binaryargs...)

		// set var to get the output
		var out bytes.Buffer

		// set the output to our variable
		command.Stdout = &out
		err := command.Run()
		if err != nil {
			log.Println(err)
		}
		stout, err := ioutil.ReadAll(&out)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(stout))
		return err
	}
}
