package db

var poolsSeed = []CreatePoolTxParams{
	{
		ID: "0xd0b53d9277642d899df5c87a3966a349a798f224",
		Token0: Token{
			Decimals: 18,
			ID:       "0x4200000000000000000000000000000000000006",
			Name:     "Wrapped Ether",
			Symbol:   "WETH",
		},
		Token1: Token{
			Decimals: 6,
			ID:       "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913",
			Name:     "USD Coin",
			Symbol:   "USDC",
		},
	},
}
