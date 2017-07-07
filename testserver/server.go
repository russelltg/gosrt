package main

import (
	"github.com/russelltg/gosrt"
	"net"
	"fmt"
)

func main() {
	bindSock, err := gosrt.NewSocket(gosrt.INET_4)
	if err != nil {
		panic(err)
	}
	
	bindSock.SetBoolSockOpt(gosrt.OPT_TSBPDMODE, true)
	
	ips, err := net.LookupIP("127.0.0.1")
	if err != nil {
		panic(err)
	}
	
	ipv4 := ips[0].To4()
	
	// print it out
	fmt.Printf("IP %s:%d\n", ipv4.String(), 7654)
	
	// bind to localhsot port 7654
	err = bindSock.Bind(ipv4, 7654)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("bound")
	
	// listen for errors
	err = bindSock.Listen()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("listening")
	
	_, _, socket, err := bindSock.Accept()
	if err != nil {
		panic(err)
	}
	
	
	fmt.Println("accepted")
	
	// recieve
	for {
		
		data, _, err := socket.RecvMsg()
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("Packet recieved: Size: %i: %s\n", len(data), data)
	}
	
}
