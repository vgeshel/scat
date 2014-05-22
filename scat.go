package main

import (
	//"fmt"
	"flag"
	//"os"
	"net"
	"log"
	"io"
	"time"
)

var verbose bool = false

func main() {
	flag.BoolVar(&verbose, "v", false, "verbose")
	
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || len(args) % 2 != 0 {
		log.Fatal("must be an even number of parameters: local remote ...")
	}

	for i := 0; i < len(args); i += 2 {
		local := args[i]
		remote := args[i + 1]

		log.Printf("%s -> %s", local, remote)

		err := scat(remote, local)
		
		if (err != nil) {
			panic(err)
		}
	}

	for {
		time.Sleep(60 * time.Second)
	}
}

func scat(remote string, local string) error {
	ln, err := net.Listen("tcp", local)

	if err != nil {
		log.Printf("ERROR listening on %s: %s", local, err)
		return err
	}

	if verbose {
		log.Printf("listening on %s", local)
	}

	go acceptLoop(ln, local, remote)

	return nil
}

func acceptLoop(ln net.Listener, local string, remote string) {
	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Printf("ERROR in accept on %s %s: %s", local, ln, err)
			continue
		}

		if verbose {
			log.Printf("accepted connection on %s", local)
		}

		go handleConnection(conn, remote)
	}
}

func handleConnection(localConn net.Conn, remote string) {
	log.Printf("connecting to %s", remote)
	
	remoteConn, err := net.Dial("tcp", remote)

	if (err != nil) {
		log.Printf("ERROR connecting to remote %s: %s", remote, err)
		return
	}

	if verbose {
		log.Printf("connected to %s", remote)
	}

	go relay(localConn, remoteConn)

	go relay(remoteConn, localConn)
}

func relay(to net.Conn, from net.Conn) {
	_, err := io.Copy(to, from)

	if (err != nil) {
		log.Printf("ERROR relaying: %s", err)
	}

	to.Close()
	from.Close()
}
