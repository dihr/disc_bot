package mock

import (
	"disc_bot/model"
	"errors"
	"fmt"
	"regexp"
)

type BitcoinMarketSvcMock interface {
	GetCoins(coin string) map[string]string
	GetCoinTicker(coin string) (model.Ticker, error)
	GetOrderBook(coin string) (model.OrderBook, error)
	SetError(value bool)
}

func NewBitcoinMarketMock() BitcoinMarketSvcMock {
	return &bitcoinMarketSvcMock{}
}

type bitcoinMarketSvcMock struct {
	ReturnsError bool
}

func (b *bitcoinMarketSvcMock) SetError(value bool) {
	b.ReturnsError = value
}

func (b *bitcoinMarketSvcMock) GetCoins(coin string) map[string]string {
	if b.ReturnsError {
		return map[string]string{}
	}
	coins := map[string]string{
		"AAVE":       "Aave",
		"ACMFT":      "AC Milan",
		"ACORDO01":   "None",
		"ADA":        "Cardano",
		"ALCX":       "Alchemix",
		"ALGO":       "Algorand",
		"ALICE":      "MyNeighborAlice",
		"ALLFT":      "Alliance",
		"AMFT":       "Aston Martin Cognizant",
		"AMP":        "Amp",
		"ANKR":       "ANKR",
		"ANT":        "Aragon",
		"ARGFT":      "Argentine Football Association",
		"ASRFT":      "AS Roma",
		"ATMFT":      "Fan Token ATM",
		"ATOM":       "Cosmos",
		"AUDIO":      "Audius",
		"AVAX":       "Avalanche",
		"AXS":        "Axie Infinity",
		"BAL":        "Balancer",
		"BAND":       "Band Protocol",
		"BARFT":      "FC Barcelona",
		"BAT":        "Basic Attention token",
		"BCH":        "Bitcoin Cash",
		"BLZ":        "Bluzelle",
		"BNT":        "BANCOR",
		"BTC":        "Bitcoin",
		"CAIFT":      "Fan Token CAI",
		"CHZ":        "Chiliz",
		"CITYFT":     "Manchester City FC",
		"COMP":       "Compound",
		"CRV":        "Curve Dao Token",
		"CSCONS01":   "Consorcio CS01",
		"CTSI":       "Cartesi",
		"CVX":        "Convex Finance",
		"DAI":        "Dai",
		"DOGE":       "Dogecoin",
		"DOT":        "Polkadot",
		"DYDX":       "dYdX",
		"ENJ":        "Enjin Coin",
		"ENS":        "Ethereum Name Service",
		"ETH":        "Ethereum",
		"FET":        "Fetch.ai",
		"FIL":        "Filecoin",
		"FLOKI":      "Floki Inu",
		"GALA":       "Gala",
		"GALFT":      "Galatasaray",
		"GALOFT":     "Clube Atletico Mineiro",
		"GNO":        "Gnosis",
		"GODS":       "Gods Unchained",
		"GRT":        "The Graph",
		"ICP":        "Internet Computer",
		"ILV":        "Illuvium",
		"IMOB01":     "None",
		"IMOB02":     "None",
		"IMX":        "Immutable X'",
		"INTERFT":    "Inter Milan",
		"JUVFT":      "Juventus",
		"KEEP":       "Keep Network",
		"KNC":        "Kyber Network",
		"KP3R":       "Keep3rV1",
		"LDO":        "Lido DAO Token",
		"LINK":       "Chainlink",
		"LOOKS":      "LooksRare",
		"LPT":        "Livepeer",
		"LQTY":       "Liquity",
		"LRC":        "Loopring",
		"LTC":        "Litecoin",
		"MANA":       "MANA (Decentraland)",
		"MATIC":      "Polygon",
		"MBCCSH01":   "Consorcio H01",
		"MBCCSH02":   "Consorcio H02",
		"MBCONS01":   "Cota de Cons??rcio 01",
		"MBCONS02":   "Cota de Cons??rcio 02",
		"MBFP01":     "None",
		"MBFP02":     "None",
		"MBFP03":     "None",
		"MBFP04":     "None",
		"MBFP05":     "None",
		"MBPRK01":    "Precat??rio MB SP01",
		"MBPRK02":    "Precat??rio MB SP02",
		"MBPRK03":    "Precat??rio MB BR03",
		"MBPRK04":    "Precat??rio MB RJ04",
		"MBPRK05":    "Fluxo de Pagamentos 5",
		"MBPRK06":    "Precatorio MB BR06",
		"MBPRK07":    "Precatorio MB SP07",
		"MBSANTOS01": "Token da Vila",
		"MBVASCO01":  "Vasco Token",
		"MC":         "Merit Circle",
		"MCO2":       "Moss Carbon Credit",
		"MENGOFT":    "Flamengo",
		"MIR":        "Mirror Protocol",
		"MKR":        "Maker",
		"NAVIFT":     "Natus Vincere",
		"NFT00":      "Vale do Outback de 100 reais",
		"NFT10":      "Iasy Tata",
		"NFT11":      "NFT Feirante Abaetetubense",
		"NFT12":      "NFT Facas Feitas",
		"NFT13":      "NFT Mandala Yawanawa - Mariri a roda 2",
		"NFT14":      "Dodge Dart Sedan 1970 Verde Imperial",
		"NFT15":      "Dodge Dart Coupe 1971 Vermelho Etrusco",
		"NFT16":      "Dodge Charger LS 1974 White",
		"NFT17":      "Dodge Charger LS 1974 Black",
		"NFT18":      "Combo de Dodges",
		"NFT2":       "NFT Protetores da Floresta",
		"NFT3":       "NFT Protetores da Floresta - Peixe Mandy",
		"NFT4":       "NFT Error",
		"NFT5":       "NFT Simulation",
		"NFT6":       "NFT Cosmovisao Tupinamba da Amazonia",
		"NFT7":       "NFT Barbara Parawara",
		"NFT8":       "NFT Liberdade de sentir",
		"NFT9":       "NFT Pescaria",
		"NFTOKN01":   "Cesta de NFTs",
		"OCEAN":      "Ocean Protocol",
		"OGFT":       "OG eSports",
		"OGN":        "Origin Protocol",
		"OMG":        "Omg Network",
		"OPUL":       "Opulous",
		"OXT":        "Orchid",
		"PAXG":       "PAX Gold",
		"PERP":       "Perpetual Protocol",
		"PFLFT":      "Professional Fighters League",
		"PLA":        "PlayDapp",
		"POLS":       "Polkastarter",
		"PORFT":      "Portugal National Team FT",
		"PSGFT":      "Paris Saint-Germain",
		"QNT":        "Quant",
		"RACA":       "Radio Caca",
		"RAD":        "Radicle",
		"RARI":       "Rarible",
		"REN":        "Ren",
		"REQ":        "Request",
		"RLY":        "Rally",
		"RNDR":       "Render Token",
		"SACI":       "Sport Club Internacional",
		"SAND":       "The Sandbox",
		"SAUBERFT":   "Alfa Romeo Racing ORLEN",
		"SCCPFT":     "Corinthians",
		"SHIB":       "Shiba Inu",
		"SKL":        "SKALE Network",
		"SLP":        "Smooth Love Potion",
		"SNX":        "Synthetix",
		"SOL":        "Solana",
		"SPELL":      "Spell Token",
		"SPFCFT":     "SPFC",
		"STVFT":      "Sint-Truidense Voetbalvereniging",
		"SUPER":      "SuperFarm",
		"SUSHI":      "SushiSwap",
		"SYN":        "Synapse",
		"THFT":       "Team Heretics",
		"TRU":        "TrueFi",
		"UFCFT":      "UFC",
		"UMA":        "Uma",
		"UNI":        "Uniswap",
		"USDC":       "USD Coin",
		"USDP":       "Pax Dollar",
		"VERDAO":     "Sociedade Esportiva Palmeiras",
		"VSPRK01":    "Precatorio VS SP01",
		"WBTC":       "Wrapped Bitcoin",
		"WBX":        "WiBX",
		"WLUNA":      "Wrapped LUNA Token",
		"XLM":        "Stellar",
		"XRP":        "XRP",
		"XTZ":        "Tezos",
		"YBOFT":      "BSC Young Boys",
		"YFI":        "yearn.finance",
		"YGG":        "Yield Guild Games",
		"ZRX":        "0x ",
	}

	result := make(map[string]string)
	for key, value := range coins {
		parameter := fmt.Sprintf("^(?i)%s", coin)
		if ok, _ := regexp.MatchString(parameter, key); ok {
			result[key] = value
		}
	}
	return result
}

func (b *bitcoinMarketSvcMock) GetCoinTicker(coin string) (model.Ticker, error) {
	if b.ReturnsError {
		return model.Ticker{}, errors.New("fail to get coin ticker")
	}
	response := model.Ticker{
		High: "14481.47000000",
		Low:  "13706.00002000",
		Vol:  "443.73564488",
		Last: "14447.01000000",
		Buy:  "14447.00100000",
		Sell: "14447.01000000",
		Date: 1502977646,
	}
	return response, nil
}

func (b *bitcoinMarketSvcMock) GetOrderBook(coin string) (model.OrderBook, error) {
	if b.ReturnsError {
		return model.OrderBook{}, errors.New("fail to get orderbook")
	}
	result := model.OrderBook{
		Asks: [][]float64{
			{10410.00006000, 2.09190016},
			{10420.00000000, 0.00997000},
			{10488.99999000, 0.46634897},
			{10410.00006000, 2.09190016},
			{10420.00000000, 0.00997000},
			{10488.99999000, 0.46634897},
			{10410.00006000, 2.09190016},
			{10420.00000000, 0.00997000},
			{10488.99999000, 0.46634897},
			{10410.00006000, 2.09190016},
		},
		Bids: [][]float64{
			{10410.00006000, 2.09190016},
			{10420.00000000, 0.00997000},
			{10488.99999000, 0.46634897},
			{10410.00006000, 2.09190016},
			{10420.00000000, 0.00997000},
			{10488.99999000, 0.46634897},
			{10410.00006000, 2.09190016},
			{10420.00000000, 0.00997000},
			{10488.99999000, 0.46634897},
			{10410.00006000, 2.09190016},
		},
	}
	return result, nil
}
