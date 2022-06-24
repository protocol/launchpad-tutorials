# IPFS Go Client Example

This folder contains the suppoting material for a Launchpad exercise that teaches the basics of the IPFS Go Client. The structure of this folder:

* `app`: contains an uncomplete version of the application that you will complete throughout the exercise.
The functions that you must complete are in the `main.go` file, and are marked with the `TO BE IMPLEMENTED` comment.

* `solution`: contains the solution files for the exercise (i.e. the complete `main.go` file).
contains the complete version of the aplication.

In order to run the application:
1. Move to the `app` directory.
2. Complete the missing functions (or copy the functions from the `main.go` file in the `solutions` directory).
3. Provide your public key (to work with IPNS) in the `YourPublicKey` constant, and a local path within your computer (to download the file) in the `YourLocalPath` constant.
4. Execute `go run .` to directly run the applicaton, or `go build` to generate an executable.

## Application structure

The application contains two files: `main.go` and `util.go`, along with the `go.mod` and `go.sum` config files.

* The `main.go` file holds the code that interacts with the IPFS node. You only have to complete the functions marked with `TO BE IMPLEMETED`. The main function `func main()` calls the different methods and handles errors (if any).
* The `util.go` contains auxiliary code (performing verifications, helpers...).

## Content covered

By the end of the exercise, you should be able to:

* Connect to a running IPFS node.
* Add a file.
* Print the content of the file.
* Download the file to your computer.
* Add the file to IPNS.