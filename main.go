package main

import (
	"bufio"
	"log"
	"os"

	"github.com/yanpozka/kava/mempool"
)

const (
	transactionsFile            = "transactions.txt"
	prioritizedTransactionsFile = "prioritized-transactions.txt"

	numTransactions = 5000
)

func main() {
	tf, err := os.Open(transactionsFile)
	// tf, err := os.Open(prioritizedTransactionsFile)
	checkErr(err)
	defer tf.Close()
	reader := bufio.NewReader(tf)

	mp := mempool.NewMempool(numTransactions)
	var c int
	for {
		line, err := reader.ReadString('\n')
		if err != nil { // EOF
			break
		}

		if err := mp.AddRawTransaction(string(line)); err != nil {
			log.Printf("error parsing transaction: %v", err)
			continue
		}
		c++
	}
	log.Printf("Read %d transactions", c)

	ptf, err := os.Create(prioritizedTransactionsFile)
	checkErr(err)
	defer ptf.Close()
	writer := bufio.NewWriter(ptf)

	transactions := mp.GetAllTransactions()
	log.Printf("Saving %d transactions into %s file ...", len(transactions), prioritizedTransactionsFile)
	c = 0
	for _, tx := range transactions {
		if _, err := writer.WriteString(tx.String()); err != nil {
			log.Println("write error:", err)
			continue
		}
		c++
	}
	if err := writer.Flush(); err != nil {
		log.Println("write/flush error:", err)
		return
	}

	log.Printf("%d transactions were store successfully!", c)
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
