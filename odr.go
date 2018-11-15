// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcutil

import (
	"bytes"
	"io"

	"github.com/btcsuite/btcd/wire"
)

// OdrIndexUnknown is the value returned for a transaction index that is unknown.
// This is typically because the transaction has not been inserted into a block
// yet.
const OdrIndexUnknown = -1

// Odr ...
type Odr struct {
	*Tx
	msgOdr *wire.MsgOdr
}

// MsgOdr returns the underlying wire.MsgOdr for the transaction.
func (t *Odr) MsgOdr() *wire.MsgOdr {
	// Return the cached transaction.
	return t.msgOdr
}

// NewOdr returns a new instance of an order given an underlying
// wire.MsgOdr.  See Odr.
func NewOdr(msgOdr *wire.MsgOdr) *Odr {
	return &Odr{
		Tx: NewTx(msgOdr.MsgTx),
		msgOdr: &wire.MsgOdr{
			MsgTx: msgOdr.MsgTx,
		},
	}
}

// NewOdrFromBytes returns a new instance of an order given the
// serialized bytes.  See Tx.
func NewOdrFromBytes(serializedTx []byte) (*Odr, error) {
	br := bytes.NewReader(serializedTx)
	return NewOdrFromReader(br)
}

// NewOdrFromReader returns a new instance of an order given a
// Reader to deserialize the transaction.  See Tx.
func NewOdrFromReader(r io.Reader) (*Odr, error) {
	// Deserialize the bytes into a MsgOdr.
	tx, err := NewTxFromReader(r)
	if err != nil {
		return nil, err
	}
	return &Odr{
		Tx: tx,
		msgOdr: &wire.MsgOdr{
			MsgTx: tx.msgTx,
		},
	}, nil
}
