package gosrt

import (
	"testing"
	"net"
)

func TestLocalConnection(t *testing.T) {

	// server
	go func() {
		
		serverBindSock, err := NewSocket(INET_4)
		if err != nil {
			t.Error("Failed to create a SRT socket", err)
		}
		
		err = serverBindSock.SetBoolSockOpt(OPT_SENDER, true)
		if err != nil {
			t.Error("Failed to set sender mode", err)
		}
		
		// create IP
		ips, err := net.LookupIP("localhost")
		if err != nil {
			t.Error("Failed to get localhost IP", err)
		}
		
		err = serverBindSock.Bind(ips[0].To4(), 9761)
		if err != nil {
			t.Error("Failed to bind sock to 9761", err)
		}
		
		_, _, sock, err := serverBindSock.Accept()
		if err != nil {
			t.Error("Failed to accept", err)
		}
		
		// send a few messages
		err = sock.SendMsg(([]byte)("Message 1"))
		if err != nil {
			t.Error("Failed to send", err)
		}
		
		sock.SendMsg(([]byte)("Message 2"))
		if err != nil {
			t.Error("Failed to send", err)
		}
		
		// close 
		sock.Close()
		
	}()
	
	// client
	go func() {
		sock, err := NewSocket(INET_4)
		if err != nil {
			t.Error("Failed to create a SRT socket", err)
		}
		
		// create IP
		ips, err := net.LookupIP("localhost")
		if err != nil {
			t.Error("Failed to get localhost IP", err)
		}
		
		err = sock.Connect(ips[0].To4(), 9761)
		if err != nil {
			t.Error("Failed to connect", err)
		}
		
		data, _, err := sock.RecvMsg()
		if err != nil {
			t.Error("Failed to rcv", err)
		}
		if (string)(data) != "Message 1" {
			t.Error("Mismatching packet data", data)
		}
		
		data, _, err = sock.RecvMsg()
		if err != nil {
			t.Error("Failed to rcv", err)
		}
		if (string)(data) != "Message 2" {
			t.Error("Mismatching packet data", data)
		}
	}()
}
