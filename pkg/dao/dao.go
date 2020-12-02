package dao

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/davecgh/go-spew/spew"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ProposalAdded The json struct for proposals added to the chain
type ProposalAdded struct {
	// msg.sender, proposalCounter, _title, _link, _requestedAmount, _beneficiary
	Entity      common.Address `json:"entity"`
	ID          *big.Int       `json:"id"`
	Title       string         `json:"title"`
	Link        []byte         `json:"link"`
	Amount      *big.Int       `json:"amount"`
	Beneficiary common.Address `json:"beneficiary"`
}

// ListenNewProposal Listens for new proposals addded on chain
func ListenNewProposal(contractAddr common.Address, proposal chan ProposalAdded, client *ethclient.Client) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(ConvictionBetaABI)))
	if err != nil {
		log.Fatal(err)
	}

	logs := make(chan types.Log)

	// subscribe
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// infinite loop handle chanel pipline for logs & errors
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			spew.Dump(vLog)
			spew.Dump(vLog.Topics)
			// event, err := contractAbi.Unpack("ProposalAdded", vLog.Data)
			var event ProposalAdded
			err = contractAbi.UnpackIntoInterface(&event, "ProposalAdded", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			event.Entity = common.HexToAddress(vLog.Topics[1].String())
			event.ID = big.NewInt(vLog.Topics[2].Big().Int64())
			spew.Dump(event)
			fmt.Println(vLog) // pointer to event log
		}
	}
}
