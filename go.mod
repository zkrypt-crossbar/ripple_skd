module github.com/zkrypt-crossbar/ripple_skd

go 1.22.0

toolchain go1.23.2

require (
	github.com/bits-and-blooms/bitset v1.17.0
	github.com/btcsuite/btcd/btcec/v2 v2.1.3
	github.com/ethereum/go-ethereum v1.15.2
	github.com/fatih/color v1.16.0
	github.com/golang/glog v1.0.0
	github.com/gorilla/websocket v1.4.2
	github.com/juju/testing v0.0.0-20210324180055-18c50b0c2098
	github.com/rubblelabs/ripple v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.32.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
)

require (
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/juju/loggo v0.0.0-20200526014432-9ce3a2e09b5e // indirect
	github.com/juju/mgo/v2 v2.0.0-20210302023703-70d5d206e208 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/rubblelabs/ripple => github.com/zkrypt-crossbar/ripple_skd v1.0.3
