package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"log"
)

func init() {
	
	SetupMyRpcincrementReply(incrementReply)

}

func incrementReply(count uint64) (MyRpcProcedure) {

	log.Printf("myRpcClient: Received Increment Reply: %v\n", count)
	return nil

}

//Main function
func main() {
	for i := 1; i < 5; i++ {
	altEthos.LogToDirectory("test/myRpcClient")
	log.Printf("myRpcClient:_before_call\n")
	fd, status := altEthos.IpcRepeat("MyRpc","",nil)
	if status != syscall.StatusOk { 
		log.Printf("Ipc_failed:_%v\n", status)
		altEthos.Exit(status)

	}

		call := MyRpcincrement{}
		status = altEthos.ClientCall(fd, &call)
		
		if status != syscall.StatusOk {
			log.Printf("clientCall failed: %v\n", status)
			altEthos.Exit(status)
	}

	}

	log.Printf("myRpcClient: _done\n")

}
