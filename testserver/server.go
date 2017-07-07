package main

import "github.com/russelltg/gosrt"
import "net"
import "fmt"

func main() {
	bindSock := gosrt.NewSocket(gosrt.INET_6)
	
	bindSock.SetBoolSockOpt(gosrt.OPT_TSBPDMODE, true)
	
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}
	
	ipv4 := ips[0]
	
	fmt.Printf("IP: %s\n", ipv4.String())
	
	// bind to localhsot port 1234
	err = bindSock.Bind(ipv4, 1234)
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
