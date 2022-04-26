// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (b *Block) Transactions(ctx context.Context) ([]*Transaction, error) {
	result, err := b.Edges.TransactionsOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryTransactions().All(ctx)
	}
	return result, err
}

func (b *Block) TransactionReceipts(ctx context.Context) ([]*TransactionReceipt, error) {
	result, err := b.Edges.TransactionReceiptsOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryTransactionReceipts().All(ctx)
	}
	return result, err
}

func (t *Transaction) Block(ctx context.Context) (*Block, error) {
	result, err := t.Edges.BlockOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryBlock().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Transaction) Receipts(ctx context.Context) ([]*TransactionReceipt, error) {
	result, err := t.Edges.ReceiptsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryReceipts().All(ctx)
	}
	return result, err
}

func (tr *TransactionReceipt) Block(ctx context.Context) (*Block, error) {
	result, err := tr.Edges.BlockOrErr()
	if IsNotLoaded(err) {
		result, err = tr.QueryBlock().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (tr *TransactionReceipt) Transaction(ctx context.Context) (*Transaction, error) {
	result, err := tr.Edges.TransactionOrErr()
	if IsNotLoaded(err) {
		result, err = tr.QueryTransaction().Only(ctx)
	}
	return result, MaskNotFound(err)
}
