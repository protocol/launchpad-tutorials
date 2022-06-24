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
	return sh.Add(strings.NewReader(text))
}

func readFile(sh *shell.Shell, cid string) (*string, error) {
	reader, err := sh.Cat(fmt.Sprintf("/ipfs/%s", cid))
	if err != nil {
		return nil, fmt.Errorf("Error reading the file: %s", err.Error())
	}

	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("Error reading bytes: %s", err.Error())
	}

	text := string(bytes)

	return &text, nil
}

func downloadFile(sh *shell.Shell, cid string) error {
	return sh.Get(cid, YourLocalPath)
}

func addToIPNS(sh *shell.Shell, cid string) error {
	var lifetime time.Duration = 50 * time.Hour
	var ttl time.Duration = 1 * time.Microsecond

	_, err := sh.PublishWithDetails(cid, YourPublicKey, lifetime, ttl, true)
	return err
}

func resolveIPNS(sh *shell.Shell) (string, error) {
	return sh.Resolve(YourPublicKey)
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
