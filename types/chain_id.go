package types

// Ethereum networks
type ChainID int

const MAINNET ChainID = 1
const ROPSTEN ChainID = 3
const RINKEBY ChainID = 4
const GORLI ChainID = 5
const KOVAN ChainID = 42
const GANACHE ChainID = 1337

var ChainIdToNetwork = map[ChainID]string{
	MAINNET: "MAINNET",
	ROPSTEN: "ROPSTEN",
	RINKEBY: "RINKEBY",
	GORLI:   "GORLI",
	KOVAN:   "KOVAN",
	GANACHE: "GANACHE",
}

var IntToChainId = map[uint]ChainID{
	1:    MAINNET,
	3:    ROPSTEN,
	4:    RINKEBY,
	5:    GORLI,
	42:   KOVAN,
	1337: GANACHE,
}
