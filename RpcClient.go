package main

import (
        "ethos/altEthos"
        "ethos/syscall"
        "log"
)

func init() {

        SetupMyRpcGetBalanceReply(GetBalanceReply)
        SetupMyRpctransferReply(transferReply)
}

func GetBalanceReply(balance float64) (MyRpcProcedure) {

        log.Printf("myRpcClient: Received Balance Reply: %v\n", balance)
        return nil

}

func transferReply(status bool) (MyRpcProcedure) {

        log.Printf("myRpcClient: Transfer completed successfully: %v\n", status)
        return nil
}

//Main function
func main() {
        altEthos.LogToDirectory("test/myRpcClient")
        log.Printf("myRpcClient: before call\n")
	for i := 1; i < 5; i++ {
                fd, status := altEthos.IpcRepeat("MyRpc","",nil)
                if status != syscall.StatusOk {
                        log.Printf("Ipc failed: %v\n", status)
                        altEthos.Exit(status)

                }
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

        call := MyRpctransfer{uint64(0), uint64(1), 50}
        status = altEthos.ClientCall(fd, &call)

        if status != syscall.StatusOk {
                log.Printf("clientCall failed: %v\n", status)
                altEthos.Exit(status)
        }

        log.Printf("myRpcClient: done\n")

}
