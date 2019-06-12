package test

import (
	"log"
	"testing"

	"../exchange/bw"
	"./conf"
	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"

	//"github.com/bitontop/gored/exchange/bw"
	"github.com/bitontop/gored/pair"
	//"github.com/bitontop/gored/test/conf"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Bw(t *testing.T) {
	e := InitBw()

	pair := pair.GetPairByKey("USDT|BTC")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)

	//Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
}

func InitBw() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	conf.Exchange(exchange.BW, config)

	ex := bw.CreateBw(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
