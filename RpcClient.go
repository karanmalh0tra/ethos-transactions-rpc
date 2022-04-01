package main

import (
        "ethos/altEthos"
        "ethos/syscall"
        "log"
)

func init() {

        SetupMyRpcGetBalanceReply(GetBalanceReply)
        SetupMyRpcTransferReply(TransferReply)
}

func GetBalanceReply(balance float64) (MyRpcProcedure) {

        log.Printf("RpcClient: Received Balance Reply for Account: %v\n", balance)
        return nil

}

func TransferReply(status bool) (MyRpcProcedure) {

        log.Printf("RpcClient: Transfer completed successfully: %v\n", status)
        return nil
}

//Main function
func main() {
        altEthos.LogToDirectory("test/myRpcClient")
        log.Printf("RpcClient: Before Call\n")
	for i := 0; i < 5; i++ {
                fd, status := altEthos.IpcRepeat("MyRpc","",nil)
                if status != syscall.StatusOk {
                        log.Printf("Ipc failed: %v\n", status)
                        altEthos.Exit(status)

                }
		// call to get Balance of each User
                call := MyRpcGetBalance{uint64(i)}
                status = altEthos.ClientCall(fd, &call)

                if status != syscall.StatusOk {
                        log.Printf("clientCall failed: %v\n", status)
                        altEthos.Exit(status)
                }

        }

        fd, status := altEthos.IpcRepeat("MyRpc","",nil)
        if status != syscall.StatusOk {
                log.Printf("Ipc failed: %v\n", status)
                altEthos.Exit(status)

        }

	// call to show what happens if you try to transfer more than balance
	call := MyRpcTransfer{uint64(0), uint64(1), 150}
	status = altEthos.ClientCall(fd, &call)
	if status != syscall.StatusOk {
		log.Printf("RpcClient: clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}


	fd1, status1 := altEthos.IpcRepeat("MyRpc","",nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status1)
		altEthos.Exit(status1)
	}
	
	// call to show a successful transfer
        call1 := MyRpcTransfer{uint64(0), uint64(1), 75}
        status1 = altEthos.ClientCall(fd1, &call1)

        if status1 != syscall.StatusOk {
                log.Printf("RpcClient: clientCall failed: %v\n", status1)
                altEthos.Exit(status1)
        }

        log.Printf("RpcClient: DONE\n")

}
