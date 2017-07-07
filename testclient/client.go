package main

import "github.com/russelltg/gosrt"
import "net"
import "fmt"

func main() {
	sock := gosrt.NewSocket(gosrt.INET_6)
	
	sock.SetBoolSockOpt(gosrt.OPT_TSBPDMODE, true)
    sock.SetBoolSockOpt(gosrt.OPT_SENDER, true)
	
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}
	
	ipv4 := ips[0]
	
	// bind to localhsot port 1234
	err = sock.Connect(ipv4, 1234)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Connected!")
	
	// send
	for {
		
        message := "Hello World!"
        
		err := sock.SendMsg([]byte(message))
		if err != nil {
			panic(err)
		}
	}
	
}
