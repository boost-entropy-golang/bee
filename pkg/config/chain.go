// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	// chain ID
	goerliChainID = int64(5)
	xdaiChainID   = int64(100)
	// start block
	// replace this with
	goerliStartBlock = uint64(7457394)
	xdaiStartBlock   = uint64(23862394)
	// factory address
	goerliContractAddress      = common.HexToAddress("0x0c9de531dcb38b758fe8a2c163444a5e54ee0db2")
	xdaiContractAddress        = common.HexToAddress("0x0FDc5429C50e2a39066D8A94F3e2D2476fcc3b85")
	goerliFactoryAddress       = common.HexToAddress("0x73c412512E1cA0be3b89b77aB3466dA6A1B9d273")
	xdaiFactoryAddress         = common.HexToAddress("0xc2d5a532cf69aa9a1378737d8ccdef884b6e7420")
	goerliLegacyFactoryAddress = common.HexToAddress("0xf0277caffea72734853b834afc9892461ea18474")
	// postage stamp
	goerliPostageStampContractAddress = common.HexToAddress("0x07456430a9878626ba42d4A26D5AfDa0A0Ca9D26")
	xdaiPostageStampContractAddress   = common.HexToAddress("0x52102F54CdfA1cC9386E34957FD14A80680e9Ba3")
)

type ChainConfig struct {
	StartBlock         uint64
	LegacyFactories    []common.Address
	PostageStamp       common.Address
	CurrentFactory     common.Address
	PriceOracleAddress common.Address
}

func GetChainConfig(chainID int64) (*ChainConfig, bool) {
	var cfg ChainConfig
	switch chainID {
	case goerliChainID:
		cfg.PostageStamp = goerliPostageStampContractAddress
		cfg.StartBlock = goerliStartBlock
		cfg.CurrentFactory = goerliFactoryAddress
		cfg.LegacyFactories = []common.Address{
			goerliLegacyFactoryAddress,
		}
		cfg.PriceOracleAddress = goerliContractAddress
		return &cfg, true
	case xdaiChainID:
		cfg.PostageStamp = xdaiPostageStampContractAddress
		cfg.StartBlock = xdaiStartBlock
		cfg.CurrentFactory = xdaiFactoryAddress
		cfg.LegacyFactories = []common.Address{}
		cfg.PriceOracleAddress = xdaiContractAddress
		return &cfg, true
	default:
		return &cfg, false
	}
}
