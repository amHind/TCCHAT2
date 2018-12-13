package main

import (
"bufio"
"fmt"
"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
		message, _ := bufio.NewReader(conn).ReadString('\n') //pour lire le msg du serveur
		fmt.Print("Message from server: " + message)
		conn.Close()







		}

