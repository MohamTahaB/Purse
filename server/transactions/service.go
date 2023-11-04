package transactions

import "fmt"

// Fetches the transaction info from the transaction info map (temporary until a DB is made), knowing the id, outputs an error if not found.
func (t *Transaction) FetchTransactionInfoFromTransaction(transactionInfoMap map[string]*TransactionInfo) (*TransactionInfo, error) {
	transactionInfoOutput, exists := transactionInfoMap[t.Id]
	if exists {
		return transactionInfoOutput, nil
	} else {
		return nil, fmt.Errorf("info of transaction of id : %s not found", t.Id)
	}
}

// Fetches the transaction from the transaction map (temporary until a DB is made), knowing the id, outputs an error if not found.
func (ti *TransactionInfo) FetchTransactionFromTransactionInfo(transactionMap map[string]*Transaction) (*Transaction, error) {
	transactionOutput, exists := transactionMap[ti.Id]
	if exists {
		return transactionOutput, nil
	} else {
		return nil, fmt.Errorf("Transaction of id : %s not found", ti.Id)
	}
}
