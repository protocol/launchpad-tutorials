package main

import (
	"fmt"
	"time"

	"io"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

// Paste here the local path of your computer where the file will be downloaded
const YourLocalPath = ""

// Paste here your public key
const YourPublicKey = ""

func addFile(sh *shell.Shell, text string) (string, error) {
	// TO BE IMPLEMENTED
}

func readFile(sh *shell.Shell, cid string) (*string, error) {
	// TO BE IMPLEMENTED
}

func downloadFile(sh *shell.Shell, cid string) error {
	// TO BE IMPLEMENTED
}

func addToIPNS(sh *shell.Shell, cid string) error {
	// TO BE IMPLEMENTED
}

func resolveIPNS(sh *shell.Shell) (string, error) {
	// TO BE IMPLEMENTED
}

func main() {
	sh := shell.NewShell("localhost:5001")

	err := performChecks(sh)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 1. Add the "Hello from Launchpad!" text to IPFS
	fmt.Println("Adding file to IPFS")
	cid, err := addFile(sh, "Hello from Launchpad!")
	if err != nil {
		fmt.Println("Error adding file to IPFS:", err.Error())
		return
	}
	fmt.Println("File added with CID:", cid)

	separator()

	// 2. Read the file by using the generated CID
	fmt.Println("Reading file")
	text, err := readFile(sh, cid)
	if err != nil {
		fmt.Println("Error reading the file:", err.Error())
		return
	}
	fmt.Println("Content of the file:", *text)

	separator()

	// 3. Download the file to your computer
	fmt.Println("Downloading file")
	err = downloadFile(sh, cid)
	if err != nil {
		fmt.Println("Error downloading file:", err.Error())
		return
	}
	fmt.Println("File donwloaded")

	separator()

	// 4. Publish the file to IPNS
	fmt.Println("Adding file to IPNS")
	err = addToIPNS(sh, cid)
	if err != nil {
		fmt.Println("Error publishing to IPNS:", err.Error())
		return
	}
	fmt.Println("File added to IPNS")

	separator()

	// 5. Resolve IPNS based on your public key
	fmt.Println("Resolving file in IPNS")
	result, err := resolveIPNS(sh)
	if err != nil {
		fmt.Println("Error resolving IPNS:", err.Error())
		return
	}

	fmt.Println("IPNS is pointing to:", result)
}
