package main

import (
	"log"
	"net/http"

	"github.com/daithi-coombes/bot-bridge-tec-forum/pkg/dao"
	"github.com/daithi-coombes/bot-bridge-tec-forum/pkg/discourse"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const endpoint = "wss://rinkeby.infura.io/ws/v3/"

// const endpoint = "wss://rpc.xdaichain.com/wss"

func main() {

	for {
		// 1. create channel
		proposal := make(chan dao.ProposalAdded)

		// 2. create client for chain
		client, err := ethclient.Dial(endpoint)
		if err != nil {
			log.Fatal(err)
		}

		// 3. subscribe for new proposal events
		contractAddress := common.HexToAddress("0x9C8963a1B7e84dED384881a918AB31Ca7C7d64fd")
		go dao.ListenNewProposal(contractAddress, proposal, client)

		// 3. create discorse instance
		d := discourse.NewDiscourse("http://localhost:9292", "", &http.Client{})
		go d.HandleProposal(proposal)
	}
}
