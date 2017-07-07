package main

import "github.com/russelltg/gosrt"
import "net"
import "fmt"
import "time"

func main() {
	sock, err := gosrt.NewSocket(gosrt.INET_4)
	if err != nil {
		panic(err)
	}
	
	sock.SetBoolSockOpt(gosrt.OPT_TSBPDMODE, true)
    sock.SetBoolSockOpt(gosrt.OPT_SENDER, true)
	
	ips, err := net.LookupIP("127.0.0.1")
	if err != nil {
		panic(err)
	}
	
	ipv4 := ips[0]
	
	err = sock.Connect(ipv4.To4(), 7654)
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
		
		time.Sleep(time.Millisecond * 100)
	}
	
}
