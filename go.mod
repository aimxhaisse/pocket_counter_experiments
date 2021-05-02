module github.com/pokt-network/relay_counter

go 1.13

require (
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/pokt-network/pocket-core v0.0.0-20210121232307-0152f9076bdf
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.33.7
)

replace github.com/tendermint/tendermint => github.com/pokt-network/tendermint v0.32.11-0.20210113203729-f92374107ace // indirect
