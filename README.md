# ethos-transactions-rpc
RPC calls for Bank Account Transactions in ethOS using Go Programming Language
--------------------------------------------------------------------------------
Karan Malhotra <br/>
CS-587 <br/>
Assignment 1
---------------------------------------------------------------------------------

## Steps To Run
- Launch Ethos on Virtual Machine using the OVA. Ensure it is running on XEN Hypervisor.
- Clone this repository.
- Execute `make install`
- A few files and directories will be created. Execute `cd server`
- Execute the command `sudo ethosRun -t`
- Once that is completed, you can view the logs by executing `ethosLog .`

## Brief Overview
- This repository contains a client-server architecture for a simple Account Transaction.
- The code is written in Golang and executed in Ethos.
- Data is stored in files located on Ethos.
- Edge cases have been handled where a user cannot transfer more money than the balance it has.
- Uses Ethos Type Notations, Ethos File Operations and Ethos Network Operations.
