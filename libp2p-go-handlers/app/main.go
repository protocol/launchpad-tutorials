package main

import (
	"bufio"
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"log"
)

func createNode() host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	return node
}

func readHelloProtocol(s network.Stream) error {
	// TO BE IMPLEMENTED: Read the stream and print its content
}

func runTargetNode() peer.AddrInfo{	
	log.Printf("Creating target node...")
	targetNode := createNode()
	log.Printf("Target node created with ID '%s'", targetNode.ID().String())

	// TO BE IMPLEMENTED: Set stream handler for the "/hello/1.0.0" protocol

	return *host.InfoFromHost(targetNode)
}

func runSourceNode(targetNodeInfo peer.AddrInfo) {
	log.Printf("Creating source node...")
	sourceNode := createNode()
	log.Printf("Source node created with ID '%s'", sourceNode.ID().String())

	sourceNode.Connect(context.Background(), targetNodeInfo)

	// TO BE IMPLEMENTED: Open stream and send message
}

func main() {
	ctx, _ := context.WithCancel(context.Background())

	info := runTargetNode()
	runSourceNode(info)

	<-ctx.Done()
}
