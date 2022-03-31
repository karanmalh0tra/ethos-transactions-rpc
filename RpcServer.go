package main

import (
        "ethos/altEthos"
        "ethos/syscall"
        "log"
        "strconv"
)

//global variables
var path = "/user/" + syscall.GetUser() + "/accounts"

func init () {

        SetupMyRpcGetBalance(GetBalance)
        SetupMyRpctransfer(transfer)

}

func readCounterVal(accountId uint64) MyType {
        var readData MyType
        status := altEthos.Read(path + "/file_" + strconv.Itoa(int(accountId)), &readData)

        if status != syscall.StatusOk {
                log.Fatalf("Error_writing_%v_%v\n", path, status)
        }
        log.Printf("Val_%v\n", readData)
        return readData
}

func writeCounterVal(data MyType, accountId uint64) {
        status := altEthos.Write(path + "/file_" + strconv.Itoa(int(accountId)), &data)

        if status != syscall.StatusOk {
                log.Fatalf("Error_writing_%v/file_%v\n", path, status)
        }
}

func GetBalance (accountId uint64) (MyRpcProcedure) {
         counterVal := readCounterVal(accountId)
         log.Printf("myRpcService: called increment \n")
         writeCounterVal(counterVal, accountId)
         log.Printf("myRpcService: called with account id %v \n", accountId)
         return &MyRpcGetBalanceReply {counterVal.Amount}

}


func transfer (from uint64, to uint64, amount float64 ) (MyRpcProcedure) {
        fromAccount := readCounterVal(from)
        toAccount := readCounterVal(to)
        if fromAccount.Amount < amount {
                return &MyRpctransferReply {false}
        }
        log.Printf("myRpcService: called transfer \n")
        fromAccount.Amount -= amount
        toAccount.Amount += amount
        writeCounterVal(fromAccount, from)
        writeCounterVal(toAccount, to)
        log.Printf("myRpcService: transfer successful from %v to %v \n", from, to)
        fromAccount = readCounterVal(from)
        toAccount = readCounterVal(to)
        log.Printf("myRpcService: final amounts %v to %v \n", fromAccount.Amount, toAccount.Amount)
        return &MyRpctransferReply {true}

}



//Main function
func main () {
        altEthos.LogToDirectory("test/myRpcService")
        listeningFd, status := altEthos.Advertise("MyRpc")
        if status != syscall.StatusOk {
                log.Printf("Advertising service failed: %s\n", status)
                altEthos.Exit(status)

        }

        data := MyType {100}

        status = altEthos.DirectoryCreate(path, &data, "boh")

        if status != syscall.StatusOk {
                log.Fatalf("Error creating directory %v because %v\n", path, status)
        }
        writeCounterVal(data, 0);
        writeCounterVal(data, 1);
        writeCounterVal(data, 2);
        writeCounterVal(data, 3);
        writeCounterVal(data, 4);

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
