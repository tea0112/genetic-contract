package main

import (
	"backend-genetic/contracts"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	ctx := context.Background()
	blockchain, err := ethclient.DialContext(ctx, "http://127.0.0.1:8545/")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")

	balance, err := blockchain.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	log.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := blockchain.PendingBalanceAt(context.Background(), account)
	log.Println(pendingBalance)

	header, err := blockchain.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(header.Number.String())

	block, err := blockchain.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block number:", block.Number().Uint64())        // 5671744
	fmt.Println("block time", block.Time())                      // 1527211625
	fmt.Println("block difficulty", block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println("block hash hex", block.Hash().Hex())            // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("block transaction", len(block.Transactions()))  // 144

	count, err := blockchain.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144

	for _, tx := range block.Transactions() {
		fmt.Println("transaction hex", tx.Hash().Hex())
		fmt.Println("ether amount", tx.Value().String())
		fmt.Println("gas", tx.Gas())
		fmt.Println("gas price", tx.GasPrice().Uint64())
		fmt.Println("nonce", tx.Nonce())
	}

	//Account #19: 0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199 (10000 ETH)
	//Private Key: 0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e
	privateKey, err := crypto.HexToECDSA("df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e")
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := blockchain.NetworkID(ctx)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	controller, err := contracts.NewController(account, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	// need to connect rpc
	//var wg sync.WaitGroup
	//wg.Add(1)
	//creatingEvent := make(chan *contracts.ControllerCreatingMovieEvent)
	//_, err = controller.WatchCreatingMovieEvent(&bind.WatchOpts{}, creatingEvent)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//go func() {
	//	e := <-creatingEvent
	//	fmt.Println("Received Event: ", e.Message)
	//	wg.Done()
	//}()

	account10TxOpts := &bind.TransactOpts{
		From:   common.HexToAddress("0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"), // Account 19
		Signer: auth.Signer,
	}

	tx, err := controller.CreateNewMovie(account10TxOpts, big.NewInt(1201), "Lords of the ring", big.NewInt(2))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created movie with data:", string(tx.Data()), ", At:", tx.Time())

	log.Println("Start Getting movie id: 1201")
	controllerMovie, err := controller.GetMovie(&bind.CallOpts{}, big.NewInt(1201))
	if err != nil {
		log.Fatal(err)
	}

	PrintMovieInfo(controllerMovie)

	fmt.Println()
	tx, err = controller.DeleteMovie(account10TxOpts, big.NewInt(1201))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted movie %d\n", 1201)

	fmt.Println()
	log.Println("Start Getting movie id: 1201")
	controllerMovie, err = controller.GetMovie(&bind.CallOpts{}, big.NewInt(1201))
	if err != nil {
		log.Fatal(err)
	}

	PrintMovieInfo(controllerMovie)
	//wg.Wait()
}

func PrintMovieInfo(m contracts.ControllerMovie) {
	log.Println("Receive movie id:", m.Id)
	log.Println("Receive movie title:", m.Title)
	log.Println("Receive movie length (hour):", m.Length)
}
