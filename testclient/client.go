package main

import "github.com/russelltg/gosrt"
import "net"

func main() {
	sock := gosrt.NewSocket(gosrt.INET_4)
	
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

	
	// send
	for {
		
        message := "Hello World!"
        
		err := sock.SendMsg([]byte(message))
		if err != nil {
			panic(err)
		}
	}
	
}
