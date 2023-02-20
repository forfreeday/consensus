package config

// BaseConfig defines the base configuration for a Tendermint node
type BaseConfig struct {
	// chainID is une
	chainID string

	// The root directory for all data
	RootDir string `mapstructure:"home"`

	// TCP or UNIX socket address of the ABCI application,
	// or the name of an ABCI application compiled in with the Tendermint binary
	ProxyApp string `mapstructure:"proxy-app"`

	// Mode of Node: full | validator | seed
	Mode string `mapstructure:"Mode"`

	// Data backedn: leveldb | rocksdb
	DBBackend string `mapstructure:"db-backend"`

	// Database directory
	DBPath string `mapstructure:"db-dir"`

	// Output level for logging
	LogLevel string `mapstructure:"log-level"`

	// Output format: 'plain' (colored text) or 'json'
	LogFormat string `mapstructure:"log-format"`

	// Path to the JSON file containing the inital validator set and other mate_data
	Genesis string `mapstructure:"genesis-file"`

	// A JSON file containing the private key to use for p2p authenticated encryption
	NodeKey string `mapstructure:"node-key-file"`

	// Mechanism to connect to the ABCI application: socket | grpc
	ABCI string `mapstructure:"abci"`

	Other map[string]interface{} `mapstructure:",remain"`
}

func DefaultConfig() *Config {

}

// Config defines the top level configuration for a Singularity node
type Config struct {
	BaseConfig `mapstructure:",squash"`

	// Options for services
	RPC     *RPCConfig     `mapstructure:"rpc"`
	P2P     *P2PConfig     `mapstructure:"p2p"`
	Mempool *menpoolConfig `mapstructure:"mempool"`
}

type RPCConfig struct {
	RootDir string `mapstructure:"home"`

	// TCP or UNIX socket address for the RPC server to listen on
	ListenAddress string `mapstructure:"laddr"`
}

type P2PConfig struct {
	RootDir string `mapstructure:"home"`
	// Address to listen for incoming connections
	ListenAddress string `mapstructure:"laddr"`
}

// MempoolConfig defines the configuration options for the Singularity mempool
type MempoolConfig struct {
	Version   string `mapstructure:"version"`
	RootDir   string `mapstructure:"home"`
	Recheck   string `mapstructure:"recheck"`
	Broadcast string `mapstructure:"broadcast"`

	// Maximum number of transaction in the mempool
	Size int `mapstructure:"size"`
}

// DefaultP2PConfig retures a defualt configuration for the peer-to-peer layer
tpye DefaultP2PConfig struct {


}
