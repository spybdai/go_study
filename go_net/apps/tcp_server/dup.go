package tcp_server

import (
	"fmt"
	"net"
	"strconv"
)

func main() {

	listenPort := "10000"

	sc := make(chan bool, 1)
	go server(sc, listenPort)

	cc := make(chan bool, 1)

	port, _ := strconv.Atoi(listenPort)
	go client(cc, port)

}

func client(stopChan chan bool, listenPort int) {

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP: net.IP("127.0.0.1"),
		Port: listenPort,
	})

	if err != nil {
		fmt.Println(err)
	}

	for index := 0; ; index++ {
		_, err = conn.Write([]byte(fmt.Sprintf("send: %d", index)))
		if err != nil {
			fmt.Println(err)
		}

		var buff []byte
		_, err = conn.Read(buff)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Get: %v", buff)
	}
}

func server(stopChan chan bool, address string) {

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		handleConn(conn)

		switch <- stopChan {
		case true:
			break
		default:
		}
	}
}


func handleConn(conn net.Conn) {
	var buff []byte
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Write(buff)
	if err != nil {
		fmt.Println(err)
	}
}
