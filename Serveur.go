package main

import (
	"fmt"
	"net"
	"strings"
)


func main() {
	fmt.Println("Serveur en cours de synchronisation...")
	ln, _ := net.Listen("tcp", ":8081") //le serveur écoute sur toutes les interfaces
	for {
		conn, _:= ln.Accept() // accepte les connexions
		//cmdLine := make([]byte,(1024 * 4)) // buffer pour stocker les données en attente de la part d'autres client
		message:="Bienvenue sur TCCHAT, vous êtes connectés :-)"
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}





	//fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	//var names []string
	//names = strings.Split("Ta,b,c", ",")
	//fmt.Printf("%q\n", names)

	//for _, name := range names {
	//	fmt.Println(name)

	//}
	//for name := range names {
	//	fmt.Println(name)

	//}
}

