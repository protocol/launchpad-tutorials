package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"strings"
)

func createSourceNode() host.Host {
	// TO BE IMPLEMENTED
}

func createTargetNode() host.Host {
	// TO BE IMPLEMENTED
}

func connectToTargetNode(sourceNode host.Host, targetNode host.Host) {
	// TO BE IMPLEMENTED
}

func countSourceNodePeers(sourceNode host.Host) int {
	// TO BE IMPLEMENTED
}

func printNodeID(host host.Host) {
	println(fmt.Sprintf("ID: %s", host.ID().String()))
}

func printNodeAddresses(host host.Host) {
	addressesString := make([]string, 0)
	for _, address := range host.Addrs() {
		addressesString = append(addressesString, address.String())
	}

	println(fmt.Sprintf("Multiaddresses: %s", strings.Join(addressesString, ", ")))
}

func main() {
	sourceNode := createSourceNode()
	println("-- SOURCE NODE INFORMATION --")
	printNodeID(sourceNode)
	printNodeAddresses(sourceNode)

	targetNode := createTargetNode()
	println("-- TARGET NODE INFORMATION --")
	printNodeID(targetNode)
	printNodeAddresses(targetNode)

	connectToTargetNode(sourceNode, targetNode)

	println(fmt.Sprintf("Source node peers: %d", countSourceNodePeers(sourceNode)))
}
