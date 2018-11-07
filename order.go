// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcutil

import (
	"bytes"
	"io"

	"github.com/btcsuite/btcd/wire"
)

// OrderIndexUnknown is the value returned for a transaction index that is unknown.
// This is typically because the transaction has not been inserted into a block
// yet.
const OrderIndexUnknown = -1

// Order ...
type Order struct {
	*Tx
	*wire.MsgOrder
}

// MsgTx returns the underlying wire.MsgTx for the transaction.
func (t *Order) MsgTx() *wire.MsgOrder {
	// Return the cached transaction.
	return t.MsgOrder
}

// NewOrder returns a new instance of an order given an underlying
// wire.MsgOrder.  See Order.
func NewOrder(msgOrder *wire.MsgOrder) *Order {
	return &Order{
		Tx: &Tx{
			msgTx:   msgOrder.MsgTx,
			txIndex: OrderIndexUnknown,
		},
		MsgOrder: &wire.MsgOrder{
			MsgTx: msgOrder.MsgTx,
		},
	}
}

// NewOrderFromBytes returns a new instance of an order given the
// serialized bytes.  See Tx.
func NewOrderFromBytes(serializedTx []byte) (*Order, error) {
	br := bytes.NewReader(serializedTx)
	return NewOrderFromReader(br)
}

// NewOrderFromReader returns a new instance of an order given a
// Reader to deserialize the transaction.  See Tx.
func NewOrderFromReader(r io.Reader) (*Order, error) {
	// Deserialize the bytes into a MsgTx.
	tx, err := NewTxFromReader(r)
	if err != nil {
		return nil, err
	}
	return &Order{
		Tx: tx,
		MsgOrder: &wire.MsgOrder{
			MsgTx: tx.msgTx,
		},
	}, nil
}
