package mempool

import "sort"

const defaultSize = 5000

type Mempool struct {
	transactions []*Transaction
	maxSize      int
}

func NewMempool(maxSize int) *Mempool {
	if maxSize <= 0 {
		maxSize = defaultSize
	}
	m := Mempool{
		transactions: make([]*Transaction, 0, maxSize+1),
		maxSize:      maxSize,
	}
	return &m
}

func (m *Mempool) AddRawTransaction(rawTx string) error {
	tx, err := ParseTransaction(rawTx)
	if err != nil {
		return err
	}
	m.AddTransaction(tx)
	return nil
}

func (m *Mempool) AddTransaction(tx *Transaction) {
	size := len(m.transactions)

	// binary search of current transaction based on its priority
	// ixFound represents the index of the current transaction if found, otherwise it's the index where it should be inserted
	ixFound := sort.Search(len(m.transactions), func(i int) bool { return m.transactions[i].Priority() <= tx.Priority() })

	if ixFound == len(m.transactions) {
		m.transactions = append(m.transactions, tx)
	} else {
		if m.transactions[ixFound].Priority() == tx.Priority() {
			println("iguales")
			ixFound++
		}
		m.transactions = append(m.transactions, nil)
		copy(m.transactions[ixFound+1:], m.transactions[ixFound:])
		m.transactions[ixFound] = tx
	}

	if size == m.maxSize {
		// we drop the last one, transaction with lowest priority
		m.transactions = m.transactions[:len(m.transactions)-1]
	}
}

func (m *Mempool) GetAllTransactions() []*Transaction {
	return m.transactions
}
