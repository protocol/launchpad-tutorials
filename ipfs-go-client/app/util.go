package main

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
)

func separator() {
	fmt.Println("--------------------------")
}

func performChecks(sh *shell.Shell) error {
	if YourLocalPath == "" {
		return fmt.Errorf("'YourLocalPath' constant is NOT defined. Please, provide a local path in your computer where the file will be downloaded")
	}

	if YourPublicKey == "" {
		return fmt.Errorf("'YourPublicKey' constant is NOT defined. Please, provide the public key of your IPFS node")
	}

	if !sh.IsUp() {
		return fmt.Errorf("You do not have an IPFS node running at port 5001")
	}

	return nil
}
