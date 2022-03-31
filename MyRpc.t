MyRpc interface {
        GetBalance(accountId uint64)(balance float64)
        Transfer(from uint64, to uint64, amount float64) (status bool)
}

MyType struct {
        Amount float64
}
