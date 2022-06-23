# IPFS Go Client Example

This folder contains the suppoting material for a Launchpad exercise that teaches the basics of the IPFS Go Client. The structure of this folder:

* `app`: contains an uncomplete version of the application that you will complete throughout the exercise.
The methods that you must complete are marked with the `TO BE IMPLEMENTED` method.

* `solution`: contains the complete version of the aplication. You can run the application of this folder directly without making any changes. Use the `go run .` command to trigger the application. At the same time, you can generate an executable for your operating system by using the `go build` command.

*IMPORTANT:* The `solution` folder contains the completed application, but you still have to provide your IPFS node's public key and a local path in your computer to make it work.

## Application structure

The application contains two files: `main.go` and `util.go`, along with the `go.mod` and `go.sum` config files.

* The `main.go` file holds the code that interacts with the IPFS node. You only have to complete the functions marked with `TO BE IMPLEMETED`. The main function `func main()` calls the different methods and handles errors (if any).
* The `util.go` contains auxiliary code (performing verification, priting helpers...).

## Content covered

By the end of the exercise, you should be able to:

* Connect to a running IPFS node.
* Add a file.
* Print the content of the file.
* Download the file to your computer.
* Add the file to IPNS.