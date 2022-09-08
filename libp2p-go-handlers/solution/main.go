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
	buf := bufio.NewReader(s)
	message, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	connection := s.Conn()

	log.Printf("Message from '%s': %s", connection.RemotePeer().String(), message)
	return nil
}

func runTargetNode(nodeInfo chan peer.AddrInfo) {
	ctx, _ := context.WithCancel(context.Background())

	log.Printf("Creating target node...")
	targetNode := createNode()
	log.Printf("Target node created with ID '%s'", targetNode.ID().String())

	// TO BE IMPLEMENTED: Set stream handler for the "/hello/1.0.0" protocol
	targetNode.SetStreamHandler("/hello/1.0.0", func(s network.Stream) {
		log.Printf("/hello/1.0.0 stream created")
		err := readHelloProtocol(s)
		if err != nil {
			s.Reset()
		} else {
			s.Close()
		}
	})

	nodeInfo <- *host.InfoFromHost(targetNode)
	<-ctx.Done()
}

func runSourceNode(targetNodeInfo peer.AddrInfo) {
	log.Printf("Creating source node...")
	sourceNode := createNode()
	log.Printf("Source node created with ID '%s'", sourceNode.ID().String())

	sourceNode.Connect(context.Background(), targetNodeInfo)

	// TO BE IMPLEMENTED: Open stream and send message
	stream, err := sourceNode.NewStream(context.Background(), targetNodeInfo.ID, "/hello/1.0.0")
	if err != nil {
		panic(err)
	}

	message := "Hello from Launchpad!\n"
	log.Printf("Sending message...")
	_, err = stream.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	ch := make(chan peer.AddrInfo)

	go runTargetNode(ch)

	info := <-ch

	go runSourceNode(info)

	<-ctx.Done()
}
