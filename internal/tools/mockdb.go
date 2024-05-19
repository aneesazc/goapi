package tools

import "time"

// mock database
type mockDB struct{}

var mockLoginDetails = map[string]*LoginDetails{
	"anees": {
		AuthToken: "123456",
		Username:  "anees",
	},
	"mike": {
		AuthToken: "123456",
		Username:  "mike",
	},
}

var mockCoinDetails = map[string]*CoinDetails{
	"anees": {
		Coins:    100,
		Username: "anees",
	},
	"mike": {
		Coins:    200,
		Username: "mike",
	},
	"john": {
		Coins:    300,
		Username: "john",
	},
}

func (m *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(1 * time.Second)

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return clientData
}

func (m *mockDB) GetUserUserCoins(username string) *CoinDetails {
	// simulate DB call
	time.Sleep(1 * time.Second)

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return clientData
}

func (m *mockDB) SetupDatabase() error {
	return nil
}
