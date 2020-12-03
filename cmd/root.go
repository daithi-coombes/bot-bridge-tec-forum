package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/daithi-coombes/bot-bridge-tec-forum/pkg/dao"
	"github.com/daithi-coombes/bot-bridge-tec-forum/pkg/discourse"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile, endpoint, discourseAPI, discourseEndpoint, daoAddress string
var rootCmd = &cobra.Command{
	Use:   "my-calc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		// 1. create channel
		proposal := make(chan dao.ProposalAdded)
		d := discourse.NewDiscourse(discourseEndpoint, discourseAPI, &http.Client{})
		client, err := ethclient.Dial(endpoint)
		if err != nil {
			log.Fatal(err)
		}

		// 2. subscribe for new proposal events
		contractAddress := common.HexToAddress(daoAddress)
		go dao.ListenNewProposal(contractAddress, proposal, client)

		// 3. handle new proposals
		for {

			select {
			case p := <-proposal:
				if err := d.HandleProposal(p); err != nil {
					log.Fatal(err)
				}

			}
			time.Sleep(time.Second * 1)
		}

	},
}

// Execute Run the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-calc.yaml)")
	rootCmd.PersistentFlags().StringVar(&discourseEndpoint, "discourse-endpoint", "", "The url (fqdn) to the discourse instance")
	rootCmd.PersistentFlags().StringVar(&discourseAPI, "discourse-key", "", "The discourse api key")
	rootCmd.PersistentFlags().StringVar(&daoAddress, "dao", "", "The contract address for the DAO (ie conviction voting contract)")
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "wss://rinkeby.infura.io/ws/v3", "The endpoint to the JSON-RPC service")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".my-calc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".my-calc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
