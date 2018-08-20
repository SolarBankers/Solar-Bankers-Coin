package main

import (
	_ "net/http/pprof"

	"github.com/skycoin/skycoin/src/skycoin"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/skycoin/skycoin/src/visor"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.24.1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "07d2490521bce42eb44de7e18e72b1e98b472c1de0707365f086310422657c4969249bd0554473cca59616f48d071e447b4901a88d91246e83b42fd815b5d39e01"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "NVLndZjvMj4ErpKw9F2cn86vjZxzB87fEy"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "02c656f4b2ab967f24109b01fb7bce30d76ce9cf1e28ef05b73cc3a5b0bd0f6447"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = "3a9a608e2f27d52fdbddd2ba2cfd7a0aabb6adb11dc9c1a4c8455080397b72eb"

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1522660056
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 300e12

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
	"172.104.173.74:30100",
"172.104.173.74:30001",
"172.104.173.74:30002",
"172.104.173.74:30000",
"172.104.52.230:30000",
"18.218.142.16:30000",
"13.58.196.172:30000",
	}
)
func main() {
	// get node config
	nodeConfig := skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "https://solarbankers.com/peers.txt",
		Port:                30005,
		WebInterfacePort:    7220,
		DataDirectory:       "$HOME/.solarbankerscoin",
		ProfileCPUFile:      "skycoin.prof",
	})

	// create a new fiber coin instance
	coin := skycoin.NewCoin(
		skycoin.Config{
			Node: *nodeConfig,
			Build: visor.BuildInfo{
				Version: Version,
				Commit:  Commit,
				Branch:  Branch,
			},
		},
		logger,
	)

	// parse config values
	coin.ParseConfig()

	// run fiber coin node
	coin.Run()
}
