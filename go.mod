module github.com/endurio/btcutil

replace github.com/btcsuite/btcd => ../ndrd

replace github.com/btcsuite => ../

require (
	github.com/btcsuite/btcd v0.0.0-20181130015935-7d2daa5bfef2
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9
)
