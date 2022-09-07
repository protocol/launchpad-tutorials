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
	buf := bufio.NewReader(s)
	message, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	s.Conn().RemotePeer().String()

	log.Printf("Message from '%s': %s", s.Conn().RemotePeer().String(), message)
	return nil
}

func runTargetNode(nodeInfo chan peer.AddrInfo) {
	ctx, _ := context.WithCancel(context.Background())

	targetNode := createNode()

	// TO BE IMPLEMENTED
	targetNode.SetStreamHandler("/hello/1.0.0", func(s network.Stream) {
		log.Println("target new stream")
		if err := doEcho(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	nodeInfo <- *host.InfoFromHost(targetNode)
	<-ctx.Done()
}

func runSourceNode(targetNodeInfo peer.AddrInfo) {
	sourceNode := createNode()

	sourceNode.Connect(context.Background(), targetNodeInfo)

	// TO BE IMPLEMENTED
	stream, err := sourceNode.NewStream(context.Background(), targetNodeInfo.ID, "/hello/1.0.0")
	if err != nil {
		panic(err)
	}

	s := "Hola\ndsdssdsds"
	_, err = stream.Write([]byte(s))
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
