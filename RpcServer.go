package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"log"
)

//global variables
var myRpc_increment_counter uint64 = 0

func init () {
	
	SetupMyRpcincrement(increment)

}

func increment () (MyRpcProcedure) {

	log.Printf("myRpcService: called increment \n")
	myRpc_increment_counter++
	return &MyRpcincrementReply {myRpc_increment_counter}

}

//Main function
func main () {
	altEthos.LogToDirectory("test/myRpcService")
	listeningFd, status := altEthos.Advertise("MyRpc") 
	if status != syscall.StatusOk { 
		log.Printf("Advertising service failed: %s\n", status)
		altEthos.Exit(status)

	}

	for {

		_, fd, status := altEthos.Import(listeningFd) 
		if status != syscall.StatusOk { 
			log.Printf(" Error calling Import: %v\n", status) 
			altEthos.Exit(status)
			}
		log.Printf("my RpcService: new connection accepted \n")

		t:= MyRpc{}
		altEthos.Handle(fd,&t)

	}
}
