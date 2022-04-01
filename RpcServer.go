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
        SetupMyRpcTransfer(Transfer)

}

func readCounterVal(accountId uint64) MyType {
        var readData MyType
        status := altEthos.Read(path + "/file_" + strconv.Itoa(int(accountId)), &readData)

        if status != syscall.StatusOk {
                log.Fatalf("Error_writing_%v_%v\n", path, status)
        }
        log.Printf("Val %v\n", readData)
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
         log.Printf("myRpcServer: called with account id %v \n", accountId)
         return &MyRpcGetBalanceReply {counterVal.Amount}

}


func Transfer (from uint64, to uint64, amount float64 ) (MyRpcProcedure) {
        fromAccount := readCounterVal(from)
        toAccount := readCounterVal(to)
        if fromAccount.Amount < amount {
		log.Printf("RpcServer: Sender doesnt have enough balance.")
                return &MyRpcTransferReply {false}
        }
        log.Printf("RpcServer: Initiating Transfer \n")
        fromAccount.Amount -= amount
        toAccount.Amount += amount
        writeCounterVal(fromAccount, from)
        writeCounterVal(toAccount, to)
        log.Printf("RpcServer: Transfer successful from %v to %v \n", from, to)
        fromAccount = readCounterVal(from)
        toAccount = readCounterVal(to)
	log.Printf("RpcServer: Amount %v Transfered from Account ID %v to Account ID %v \n", amount, from, to)
        log.Printf("RpcServer: From Account ID Balance is %v and To Account ID Balance is %v \n", fromAccount.Amount, toAccount.Amount)
        return &MyRpcTransferReply {true}

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
	data1 := MyType {150}
	data2 := MyType {200}
	data3 := MyType {250}
	data4 := MyType {300}

        status = altEthos.DirectoryCreate(path, &data, "boh")

        if status != syscall.StatusOk {
                log.Fatalf("RpcServer: Error creating directory %v because %v\n", path, status)
        }
        writeCounterVal(data, 0);
        writeCounterVal(data1, 1);
        writeCounterVal(data2, 2);
        writeCounterVal(data3, 3);
        writeCounterVal(data4, 4);

        for {

                _, fd, status := altEthos.Import(listeningFd)
                if status != syscall.StatusOk {
                        log.Printf("RpcServer: Error calling Import: %v\n", status)
                        altEthos.Exit(status)
                        }
                log.Printf("my RpcService: new connection accepted \n")

                t:= MyRpc{}
                altEthos.Handle(fd,&t)

        }
}
