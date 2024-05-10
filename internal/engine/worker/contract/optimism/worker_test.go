package optimism_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/optimism"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/optimism"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/schema/worker"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Ethereum(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *activityx.Activity
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Deposit ETH from L1 to L2",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x7104a8431d9da16cd21385e840508fa7b657a9a6d9ad7b7b43efadbf5abd060c"),
						ParentHash:   common.HexToHash("0x70a963d867910fd24c284cd6bba787cb8f8db72a5395af9037506bfd2daa2679"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17459684", 0)),
						GasLimit:     30000000,
						GasUsed:      10186487,
						Timestamp:    1686520751,
						BaseFee:      lo.Must(new(big.Int).SetString("16406848556", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
						From:      common.HexToAddress("0x31E7B932C655628fDA6F22f063D550d990dBA726"),
						Gas:       130725,
						GasPrice:  lo.Must(new(big.Int).SetString("16500848556", 10)),
						Hash:      common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
						Input:     hexutil.MustDecode("0xb1a1a8820000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1")),
						Value:     lo.Must(new(big.Int).SetString("100857001000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x7104a8431d9da16cd21385e840508fa7b657a9a6d9ad7b7b43efadbf5abd060c"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17459684", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10165487,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3d786f7ac"),
						GasUsed:           127267,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23"),
									common.HexToHash("0x00000000000000000000000031e7b932c655628fda6f22f063d550d990dba726"),
									common.HexToHash("0x00000000000000000000000031e7b932c655628fda6f22f063d550d990dba726"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000577ac0c0bd437900000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17459684", 0)),
								TransactionHash: common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
								Index:           260,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5"),
									common.HexToHash("0x00000000000000000000000031e7b932c655628fda6f22f063d550d990dba726"),
									common.HexToHash("0x00000000000000000000000031e7b932c655628fda6f22f063d550d990dba726"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000577ac0c0bd437900000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17459684", 0)),
								TransactionHash: common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
								Index:           261,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed"),
								Topics: []common.Hash{
									common.HexToHash("0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32"),
									common.HexToHash("0x00000000000000000000000036bde71c97b33cc4729cf772ae268934f7ab70b2"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001ed00000000000000000000000000000000000000000000000577ac0c0bd437900000000000000000000000000000000000000000000000000577ac0c0bd43790000000000000077d2e00d764ad0b00010000000000000000000000000000000000000000000000000000000008a700000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1000000000000000000000000420000000000000000000000000000000000001000000000000000000000000000000000000000000000000577ac0c0bd43790000000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000a41635f5fd00000000000000000000000031e7b932c655628fda6f22f063d550d990dba72600000000000000000000000031e7b932c655628fda6f22f063d550d990dba72600000000000000000000000000000000000000000000000577ac0c0bd4379000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17459684", 0)),
								TransactionHash: common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
								Index:           262,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1"),
								Topics: []common.Hash{
									common.HexToHash("0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1000000000000000000000000000000000000000000000000000000000000008000010000000000000000000000000000000000000000000000000000000008a70000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000a41635f5fd00000000000000000000000031e7b932c655628fda6f22f063d550d990dba72600000000000000000000000031e7b932c655628fda6f22f063d550d990dba72600000000000000000000000000000000000000000000000577ac0c0bd43790000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17459684", 0)),
								TransactionHash: common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
								Index:           263,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1"),
								Topics: []common.Hash{
									common.HexToHash("0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546"),
									common.HexToHash("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000577ac0c0bd4379000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17459684", 0)),
								TransactionHash: common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
								Index:           264,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1"),
						TransactionIndex: 124,
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x3da542a1670a828838d2916d33be51de1b505e8700ffbb4ef26f31abc4184bc1",
				Network: network.Ethereum,
				Index:   124,
				From:    "0x31E7B932C655628fDA6F22f063D550d990dBA726",
				To:      optimism.AddressL1StandardBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xb1a1a882",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("2100013493176452")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x31E7B932C655628fDA6F22f063D550d990dBA726",
						To:       "0x31E7B932C655628fDA6F22f063D550d990dBA726",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Optimism,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("100857001000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686520751,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit USDC from L1 to L2",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x977ee22b35c6a6300d5a9a7e5c5be3ad9e23e7a671c4e67f42abc0160021cb50"),
						ParentHash:   common.HexToHash("0xd95b173acfa79298158f4c237a2897384d48ca029fd450e54e69f5b600053997"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990"),
						Number:       lo.Must(new(big.Int).SetString("17457885", 0)),
						GasLimit:     30000000,
						GasUsed:      21786195,
						Timestamp:    1686498803,
						BaseFee:      lo.Must(new(big.Int).SetString("15935599330", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
						From:      common.HexToAddress("0x98389CA467c0199D7379fa1B6992A389b56bCc15"),
						Gas:       180729,
						GasPrice:  lo.Must(new(big.Int).SetString("16035599330", 10)),
						Hash:      common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
						Input:     hexutil.MustDecode("0x58a997f6000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c3160700000000000000000000000000000000000000000000000000000002b2d15d800000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x977ee22b35c6a6300d5a9a7e5c5be3ad9e23e7a671c4e67f42abc0160021cb50"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17457885", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 11197369,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3bbcbd3e2"),
						GasUsed:           167049,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x00000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc15"),
									common.HexToHash("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000002b2d15d80"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17457885", 0)),
								TransactionHash: common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
								Index:           232,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396"),
									common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
									common.HexToHash("0x0000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c31607"),
									common.HexToHash("0x00000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc15"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc1500000000000000000000000000000000000000000000000000000002b2d15d8000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17457885", 0)),
								TransactionHash: common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
								Index:           233,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf"),
									common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
									common.HexToHash("0x0000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c31607"),
									common.HexToHash("0x00000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc15"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc1500000000000000000000000000000000000000000000000000000002b2d15d8000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17457885", 0)),
								TransactionHash: common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
								Index:           234,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed"),
								Topics: []common.Hash{
									common.HexToHash("0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32"),
									common.HexToHash("0x00000000000000000000000036bde71c97b33cc4729cf772ae268934f7ab70b2"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000022d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007812e00d764ad0b000100000000000000000000000000000000000000000000000000000000083b00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1000000000000000000000000420000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000e40166a07a0000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c31607000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc1500000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc1500000000000000000000000000000000000000000000000000000002b2d15d8000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17457885", 0)),
								TransactionHash: common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
								Index:           235,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1"),
								Topics: []common.Hash{
									common.HexToHash("0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be10000000000000000000000000000000000000000000000000000000000000080000100000000000000000000000000000000000000000000000000000000083b0000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000e40166a07a0000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c31607000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc1500000000000000000000000098389ca467c0199d7379fa1b6992a389b56bcc1500000000000000000000000000000000000000000000000000000002b2d15d8000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17457885", 0)),
								TransactionHash: common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
								Index:           236,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1"),
								Topics: []common.Hash{
									common.HexToHash("0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546"),
									common.HexToHash("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("17457885", 0)),
								TransactionHash: common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
								Index:           237,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751"),
						TransactionIndex: 143,
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x239293e2a707a187305ab6056ba6bf6fa6279ba05440dce6c3b8f534c4963751",
				Network: network.Ethereum,
				Index:   143,
				From:    "0x98389CA467c0199D7379fa1B6992A389b56bCc15",
				To:      optimism.AddressL1StandardBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x58a997f6",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("2678730832477170")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x98389CA467c0199D7379fa1B6992A389b56bCc15",
						To:       "0x98389CA467c0199D7379fa1B6992A389b56bCc15",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Optimism,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("11590000000"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686498803,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit ETH from L2 to L1",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x785aa867a5a6a44824a11f31010b57c4f86b5f017b87ee3a14e855cb78f3bb36"),
						ParentHash:   common.HexToHash("0xd45ad740fe003078a9500e2523ce40f4bb318ca5c25ae77a1643f939196846a8"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("105376301", 0)),
						GasLimit:     30000000,
						GasUsed:      4567467,
						Timestamp:    1686351379,
						BaseFee:      lo.Must(new(big.Int).SetString("54", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
						From:      common.HexToAddress("0x175C0aD71624A537Cf594751A0A98c2FE85F950C"),
						Gas:       120553,
						GasPrice:  lo.Must(new(big.Int).SetString("116", 10)),
						Hash:      common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
						Input:     hexutil.MustDecode("0x32b7006d000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead00000000000000000000000000000000000000000000000000154f8f552148318000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4200000000000000000000000000000000000010")),
						Value:     lo.Must(new(big.Int).SetString("393114520000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x785aa867a5a6a44824a11f31010b57c4f86b5f017b87ee3a14e855cb78f3bb36"),
						BlockNumber:       lo.Must(new(big.Int).SetString("105376301", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 4406169,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x74"),
						GasUsed:           117071,
						L1GasPrice:        lo.Must(new(big.Int).SetString("18236934242", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("2800", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("34927376460278", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
									common.HexToHash("0x000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead0000"),
									common.HexToHash("0x000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c0000000000000000000000000000000000000000000000154f8f55214831800000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105376301", 0)),
								TransactionHash: common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
								Index:           42,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5"),
									common.HexToHash("0x000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c"),
									common.HexToHash("0x000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000154f8f55214831800000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105376301", 0)),
								TransactionHash: common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
								Index:           43,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000016"),
								Topics: []common.Hash{
									common.HexToHash("0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054"),
									common.HexToHash("0x00010000000000000000000000000000000000000000000000000000000000d2"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
									common.HexToHash("0x00000000000000000000000025ace71c97b33cc4729cf772ae268934f7ab5fa1"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000154f8f5521483180000000000000000000000000000000000000000000000000000000000000046388000000000000000000000000000000000000000000000000000000000000008013024c405de395803bc147a0d88b2c6b256b649490653fe397dc5cccd79d5a9100000000000000000000000000000000000000000000000000000000000001a4d764ad0b00010000000000000000000000000000000000000000000000000000000000d2000000000000000000000000420000000000000000000000000000000000001000000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be10000000000000000000000000000000000000000000000154f8f552148318000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000a41635f5fd000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c0000000000000000000000000000000000000000000000154f8f552148318000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105376301", 0)),
								TransactionHash: common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
								Index:           44,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a"),
									common.HexToHash("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000004200000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000008000010000000000000000000000000000000000000000000000000000000000d2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a41635f5fd000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c000000000000000000000000175c0ad71624a537cf594751a0a98c2fe85f950c0000000000000000000000000000000000000000000000154f8f5521483180000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105376301", 0)),
								TransactionHash: common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
								Index:           45,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000154f8f552148318000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105376301", 0)),
								TransactionHash: common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
								Index:           46,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3"),
						TransactionIndex: 7,
					},
				},
				config: &config.Module{
					Network: network.Optimism,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Optimism),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x7831e3145f8ca8ef0922397bcd5e2da2b1847cc515f2ceb4df74ecc8267505c3",
				Network: network.Ethereum,
				Index:   7,
				From:    "0x175C0aD71624A537Cf594751A0A98c2FE85F950C",
				To:      optimism.AddressL2StandardBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x32b7006d",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("34927390040514")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x175C0aD71624A537Cf594751A0A98c2FE85F950C",
						To:       "0x175C0aD71624A537Cf594751A0A98c2FE85F950C",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Optimism,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("393114520000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686351379,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit CRV from L2 to L1",
			arguments: arguments{
				task: &source.Task{
					Network: network.Optimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x36a395a20d10db49db868d06f7af1325f3375063fc98eef8015edb0ad9075941"),
						ParentHash:   common.HexToHash("0x68103dbb746bb38d4e2a8ddf860b3856d5da5eb6954c19e9d964fd9d66984094"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("105457219", 0)),
						GasLimit:     30000000,
						GasUsed:      1096404,
						Timestamp:    1686513215,
						BaseFee:      lo.Must(new(big.Int).SetString("73", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
						From:      common.HexToAddress("0x7a16fF8270133F063aAb6C9977183D9e72835428"),
						Gas:       137736,
						GasPrice:  lo.Must(new(big.Int).SetString("186", 10)),
						Hash:      common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
						Input:     hexutil.MustDecode("0x32b7006d0000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb530000000000000000000000000000000000000000000000c8152697b9175d9afd000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4200000000000000000000000000000000000010")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x36a395a20d10db49db868d06f7af1325f3375063fc98eef8015edb0ad9075941"),
						BlockNumber:       lo.Must(new(big.Int).SetString("105457219", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 681525,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xba"),
						GasUsed:           129308,
						L1GasPrice:        lo.Must(new(big.Int).SetString("16847193050", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("2704", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("31159490044924", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x0994206dfE8De6Ec6920FF4D779B0d950605Fb53"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000c8152697b9175d9afd"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           8,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x0994206dfE8De6Ec6920FF4D779B0d950605Fb53"),
								Topics: []common.Hash{
									common.HexToHash("0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000c8152697b9175d9afd"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           9,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e"),
									common.HexToHash("0x000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd52"),
									common.HexToHash("0x0000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb53"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           10,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf"),
									common.HexToHash("0x0000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb53"),
									common.HexToHash("0x000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd52"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           11,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000016"),
								Topics: []common.Hash{
									common.HexToHash("0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054"),
									common.HexToHash("0x0001000000000000000000000000000000000000000000000000000000000144"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
									common.HexToHash("0x00000000000000000000000025ace71c97b33cc4729cf772ae268934f7ab5fa1"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004678800000000000000000000000000000000000000000000000000000000000000802f818b6904d6fc0527bf44dd91cf73f419f7bc138d600a11dc211c63e7fb47ec00000000000000000000000000000000000000000000000000000000000001e4d764ad0b0001000000000000000000000000000000000000000000000000000000000144000000000000000000000000420000000000000000000000000000000000001000000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000e40166a07a000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd520000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb530000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           12,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a"),
									common.HexToHash("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000420000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000800001000000000000000000000000000000000000000000000000000000000144000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e40166a07a000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd520000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb530000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           13,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           14,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
						TransactionIndex: 4,
					},
				},
				config: &config.Module{
					Network: network.Optimism,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Optimism),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701",
				Network: network.Optimism,
				Index:   4,
				From:    "0x7a16fF8270133F063aAb6C9977183D9e72835428",
				To:      optimism.AddressL2StandardBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x32b7006d",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("31159514096212")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x7a16fF8270133F063aAb6C9977183D9e72835428",
						To:       "0x7a16fF8270133F063aAb6C9977183D9e72835428",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Optimism,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x0994206dfE8De6Ec6920FF4D779B0d950605Fb53"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("3690872887087038569213"))),
								Name:     "Curve DAO Token",
								Symbol:   "CRV",
								Decimals: 18,
								Standard: 20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686513215,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw ETH from L1 to L2",
			arguments: arguments{
				task: &source.Task{
					Network: network.Optimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x085154c4464955cb8e46541c514667b882681ea4a9cb81d3be29a6d6d1f3ae5c"),
						ParentHash:   common.HexToHash("0x55917f2448ac2df1d78a54b352345521ddc72306684b6232531607b2be9c5cf0"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("114344502", 0)),
						GasLimit:     30000000,
						GasUsed:      9425894,
						Timestamp:    1704287781,
						BaseFee:      lo.Must(new(big.Int).SetString("121384510", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xb0c0b12cc5b5cd67993d4079809c41474ffff2a2de20bbe765f0cec7f07cf9de"),
						From:      common.HexToAddress("0x36BDE71C97B33Cc4729cf772aE268934f7AB70B2"),
						Gas:       490798,
						GasPrice:  lo.Must(new(big.Int).SetString("0", 10)),
						Hash:      common.HexToHash("0xb0c0b12cc5b5cd67993d4079809c41474ffff2a2de20bbe765f0cec7f07cf9de"),
						Input:     hexutil.MustDecode("0xd764ad0b000100000000000000000000000000000000000000000000000000000001472600000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be10000000000000000000000004200000000000000000000000000000000000010000000000000000000000000000000000000000000000000a688906bd8b000000000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000a41635f5fd000000000000000000000000dc548196ba87b19e55d2570d1dc8d4a1dc86238c000000000000000000000000dc548196ba87b19e55d2570d1dc8d4a1dc86238c000000000000000000000000000000000000000000000000a688906bd8b000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4200000000000000000000000000000000000007")),
						Value:     lo.Must(new(big.Int).SetString("12000000000000000000", 0)),
						Type:      126,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x085154c4464955cb8e46541c514667b882681ea4a9cb81d3be29a6d6d1f3ae5c"),
						BlockNumber:       lo.Must(new(big.Int).SetString("114344502", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 142735,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x0"),
						GasUsed:           92234,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
									common.HexToHash("0x000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead0000"),
									common.HexToHash("0x000000000000000000000000dc548196ba87b19e55d2570d1dc8d4a1dc86238c"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000dc548196ba87b19e55d2570d1dc8d4a1dc86238c000000000000000000000000000000000000000000000000a688906bd8b0000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("114344502", 0)),
								TransactionHash: common.HexToHash("0xb0c0b12cc5b5cd67993d4079809c41474ffff2a2de20bbe765f0cec7f07cf9de"),
								Index:           0,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d"),
									common.HexToHash("0x000000000000000000000000dc548196ba87b19e55d2570d1dc8d4a1dc86238c"),
									common.HexToHash("0x000000000000000000000000dc548196ba87b19e55d2570d1dc8d4a1dc86238c"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000a688906bd8b0000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("114344502", 0)),
								TransactionHash: common.HexToHash("0xb0c0b12cc5b5cd67993d4079809c41474ffff2a2de20bbe765f0cec7f07cf9de"),
								Index:           1,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c"),
									common.HexToHash("0x5b32c7fee67ae15a689d3e1bde466af0ba19cff78d6f2bc48d708e14800042f3"),
								},
								Data:            nil,
								BlockNumber:     lo.Must(new(big.Int).SetString("114344502", 0)),
								TransactionHash: common.HexToHash("0xb0c0b12cc5b5cd67993d4079809c41474ffff2a2de20bbe765f0cec7f07cf9de"),
								Index:           2,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0xb0c0b12cc5b5cd67993d4079809c41474ffff2a2de20bbe765f0cec7f07cf9de"),
						TransactionIndex: 1,
					},
				},
				config: &config.Module{
					Network: network.Optimism,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Optimism),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xb0c0b12cc5b5cd67993d4079809c41474ffff2a2de20bbe765f0cec7f07cf9de",
				Network: network.Optimism,
				Index:   1,
				From:    "0x36BDE71C97B33Cc4729cf772aE268934f7AB70B2",
				To:      optimism.AddressL2ETH.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xd764ad0b",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("0")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0xDc548196bA87b19E55D2570d1dC8D4A1dc86238c",
						To:       "0xDc548196bA87b19E55D2570d1dC8D4A1dc86238c",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Optimism,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("12000000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1704287781,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw USDC from L1 to L2",
			arguments: arguments{
				task: &source.Task{
					Network: network.Optimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x749a7d4530b753ed6c614fc38e18518e1155846b157ef48e9119148bb3c1068f"),
						ParentHash:   common.HexToHash("0xa09ac75d78e6b50b34945eb1b31342b7361721e820f4f14b16f5d438c7e73e3f"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("68137544", 0)),
						GasLimit:     15000000,
						GasUsed:      139461,
						Timestamp:    1673901977,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f"),
						From:      common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Gas:       200000,
						GasPrice:  lo.Must(new(big.Int).SetString("0", 10)),
						Hash:      common.HexToHash("0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f"),
						Input:     hexutil.MustDecode("0xcbd4ece9000000000000000000000000420000000000000000000000000000000000001000000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000423e600000000000000000000000000000000000000000000000000000000000000e4662a633a000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c316070000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428000000000000000000000000000000000000000000000000000000007735940000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4200000000000000000000000000000000000007")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x749a7d4530b753ed6c614fc38e18518e1155846b157ef48e9119148bb3c1068f"),
						BlockNumber:       lo.Must(new(big.Int).SetString("68137544", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 139461,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x0"),
						GasUsed:           139461,
						L1GasPrice:        lo.Must(new(big.Int).SetString("28545416243", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("6672", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("190455017173296", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("1")),
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000077359400"),
								BlockNumber:     lo.Must(new(big.Int).SetString("68137544", 0)),
								TransactionHash: common.HexToHash("0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f"),
								Index:           0,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
								Topics: []common.Hash{
									common.HexToHash("0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000077359400"),
								BlockNumber:     lo.Must(new(big.Int).SetString("68137544", 0)),
								TransactionHash: common.HexToHash("0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f"),
								Index:           1,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89"),
									common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
									common.HexToHash("0x0000000000000000000000007f5c764cbc14f9669b88837ca1490cca17c31607"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428000000000000000000000000000000000000000000000000000000007735940000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("68137544", 0)),
								TransactionHash: common.HexToHash("0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f"),
								Index:           2,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c"),
									common.HexToHash("0x3ae08543db87fb943d7b7a2cc0123bd9c202fb3c61db635aafb1bf7c0a075f6e"),
								},
								Data:            nil,
								BlockNumber:     lo.Must(new(big.Int).SetString("68137544", 0)),
								TransactionHash: common.HexToHash("0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f"),
								Index:           3,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Optimism,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Optimism),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xe1a660f3008f200735605f8154e9605a42e9f90687167b2a8f276ead650d850f",
				Network: network.Optimism,
				Index:   0,
				From:    "0x0000000000000000000000000000000000000000",
				To:      optimism.AddressL2ETH.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xcbd4ece9",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("190455017173296")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x7a16fF8270133F063aAb6C9977183D9e72835428",
						To:       "0x7a16fF8270133F063aAb6C9977183D9e72835428",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Optimism,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2000000000"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: 20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1673901977,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit CRV from L2 to L1",
			arguments: arguments{
				task: &source.Task{
					Network: network.Optimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x36a395a20d10db49db868d06f7af1325f3375063fc98eef8015edb0ad9075941"),
						ParentHash:   common.HexToHash("0x68103dbb746bb38d4e2a8ddf860b3856d5da5eb6954c19e9d964fd9d66984094"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("105457219", 0)),
						GasLimit:     30000000,
						GasUsed:      1096404,
						Timestamp:    1686513215,
						BaseFee:      lo.Must(new(big.Int).SetString("73", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
						From:      common.HexToAddress("0x7a16fF8270133F063aAb6C9977183D9e72835428"),
						Gas:       137736,
						GasPrice:  lo.Must(new(big.Int).SetString("186", 10)),
						Hash:      common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
						Input:     hexutil.MustDecode("0x32b7006d0000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb530000000000000000000000000000000000000000000000c8152697b9175d9afd000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x4200000000000000000000000000000000000010")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x36a395a20d10db49db868d06f7af1325f3375063fc98eef8015edb0ad9075941"),
						BlockNumber:       lo.Must(new(big.Int).SetString("105457219", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 681525,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xba"),
						GasUsed:           129308,
						L1GasPrice:        lo.Must(new(big.Int).SetString("16847193050", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("2704", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("31159490044924", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x0994206dfE8De6Ec6920FF4D779B0d950605Fb53"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
									common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000c8152697b9175d9afd"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           8,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x0994206dfE8De6Ec6920FF4D779B0d950605Fb53"),
								Topics: []common.Hash{
									common.HexToHash("0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000c8152697b9175d9afd"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           9,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e"),
									common.HexToHash("0x000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd52"),
									common.HexToHash("0x0000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb53"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           10,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000010"),
								Topics: []common.Hash{
									common.HexToHash("0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf"),
									common.HexToHash("0x0000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb53"),
									common.HexToHash("0x000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd52"),
									common.HexToHash("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e72835428"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           11,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000016"),
								Topics: []common.Hash{
									common.HexToHash("0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054"),
									common.HexToHash("0x0001000000000000000000000000000000000000000000000000000000000144"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
									common.HexToHash("0x00000000000000000000000025ace71c97b33cc4729cf772ae268934f7ab5fa1"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004678800000000000000000000000000000000000000000000000000000000000000802f818b6904d6fc0527bf44dd91cf73f419f7bc138d600a11dc211c63e7fb47ec00000000000000000000000000000000000000000000000000000000000001e4d764ad0b0001000000000000000000000000000000000000000000000000000000000144000000000000000000000000420000000000000000000000000000000000001000000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000e40166a07a000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd520000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb530000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           12,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a"),
									common.HexToHash("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000420000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000800001000000000000000000000000000000000000000000000000000000000144000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e40166a07a000000000000000000000000d533a949740bb3306d119cc777fa900ba034cd520000000000000000000000000994206dfe8de6ec6920ff4d779b0d950605fb530000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000007a16ff8270133f063aab6c9977183d9e728354280000000000000000000000000000000000000000000000c8152697b9175d9afd00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           13,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4200000000000000000000000000000000000007"),
								Topics: []common.Hash{
									common.HexToHash("0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546"),
									common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000010"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("105457219", 0)),
								TransactionHash: common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
								Index:           14,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701"),
						TransactionIndex: 4,
					},
				},
				config: &config.Module{
					Network: network.Optimism,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Optimism),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xf96590b94c25df375eaa0cece0477d2a3b3fe1756d221f9efa1bf48ae9117701",
				Network: network.Optimism,
				Index:   4,
				From:    "0x7a16fF8270133F063aAb6C9977183D9e72835428",
				To:      optimism.AddressL2StandardBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x32b7006d",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("31159514096212")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x7a16fF8270133F063aAb6C9977183D9e72835428",
						To:       "0x7a16fF8270133F063aAb6C9977183D9e72835428",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Optimism,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x0994206dfE8De6Ec6920FF4D779B0d950605Fb53"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("3690872887087038569213"))),
								Name:     "Curve DAO Token",
								Symbol:   "CRV",
								Decimals: 18,
								Standard: 20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686513215,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw USDT from L2 to L1",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x3a1536f32c7c0d0c9da3e9f5d87c0d38256ccd92a7cad30030d3c394b00c5bb3"),
						ParentHash:   common.HexToHash("0x3fffd9c7368352d92c18408c6105118487e2d93174800d6e6c600a2631120509"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("18923235", 0)),
						GasLimit:     30000000,
						GasUsed:      13261374,
						Timestamp:    1704241355,
						BaseFee:      lo.Must(new(big.Int).SetString("17897398074", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
						From:      common.HexToAddress("0x2Ce910fBba65B454bBAf6A18c952A70f3bcd8299"),
						Gas:       500000,
						GasPrice:  lo.Must(new(big.Int).SetString("18097398074", 10)),
						Hash:      common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
						Input:     hexutil.MustDecode("0x8c3152e900000000000000000000000000000000000000000000000000000000000000200001000000000000000000000000000000000000000000000000000000002f7c000000000000000000000000420000000000000000000000000000000000000700000000000000000000000025ace71c97b33cc4729cf772ae268934f7ab5fa10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004678800000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e4d764ad0b0001000000000000000000000000000000000000000000000000000000002f7c000000000000000000000000420000000000000000000000000000000000001000000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000e40166a07a000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec700000000000000000000000094b008aa00579c1307b0ef2c499ad98a8ce58e580000000000000000000000002ce910fbba65b454bbaf6a18c952a70f3bcd82990000000000000000000000002ce910fbba65b454bbaf6a18c952a70f3bcd8299000000000000000000000000000000000000000000000000000001c3fbd4876000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x3a1536f32c7c0d0c9da3e9f5d87c0d38256ccd92a7cad30030d3c394b00c5bb3"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18923235", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10544405,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x436b0613a"),
						GasUsed:           217495,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x00000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1"),
									common.HexToHash("0x0000000000000000000000002ce910fbba65b454bbaf6a18c952a70f3bcd8299"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000001c3fbd48760"),
								BlockNumber:     lo.Must(new(big.Int).SetString("18923235", 0)),
								TransactionHash: common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
								Index:           276,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3"),
									common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
									common.HexToHash("0x00000000000000000000000094b008aa00579c1307b0ef2c499ad98a8ce58e58"),
									common.HexToHash("0x0000000000000000000000002ce910fbba65b454bbaf6a18c952a70f3bcd8299"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000002ce910fbba65b454bbaf6a18c952a70f3bcd8299000000000000000000000000000000000000000000000000000001c3fbd4876000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("18923235", 0)),
								TransactionHash: common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
								Index:           277,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd"),
									common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
									common.HexToHash("0x00000000000000000000000094b008aa00579c1307b0ef2c499ad98a8ce58e58"),
									common.HexToHash("0x0000000000000000000000002ce910fbba65b454bbaf6a18c952a70f3bcd8299"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000002ce910fbba65b454bbaf6a18c952a70f3bcd8299000000000000000000000000000000000000000000000000000001c3fbd4876000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("18923235", 0)),
								TransactionHash: common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
								Index:           278,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1"),
								Topics: []common.Hash{
									common.HexToHash("0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c"),
									common.HexToHash("0x71ace35385c826a1664da4e51a84bb839e0d62c232792f6d84fc5622e1ffe9b3"),
								},
								Data:            nil,
								BlockNumber:     lo.Must(new(big.Int).SetString("18923235", 0)),
								TransactionHash: common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
								Index:           279,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed"),
								Topics: []common.Hash{
									common.HexToHash("0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b"),
									common.HexToHash("0x2e0509242efb8908112741decc6e7112fc2699e3308f66fe7150589b0b35cb0c"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001"),
								BlockNumber:     lo.Must(new(big.Int).SetString("18923235", 0)),
								TransactionHash: common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
								Index:           280,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84"),
						TransactionIndex: 85,
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x90fd9056fa8b2b2412bb0803509d4f7937b287d205fe3823c7b3a02e46166a84",
				Network: network.Ethereum,
				Index:   85,
				From:    "0x2Ce910fBba65B454bBAf6A18c952A70f3bcd8299",
				To:      optimism.AddressL1OptimismPortal.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x8c3152e9",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3936093594104630")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x2Ce910fBba65B454bBAf6A18c952A70f3bcd8299",
						To:       "0x2Ce910fBba65B454bBAf6A18c952A70f3bcd8299",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Optimism,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1941255260000"))),
								Name:     "Tether USD",
								Symbol:   "USDT",
								Decimals: 6,
								Standard: 20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1704241355,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw ETH from L2 to L1",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xa2f39abba7db061bc5653af3f55e72ea79cb4719cbe301523f7c266428ab969d"),
						ParentHash:   common.HexToHash("0x907efce0d393656df83caeead3e5091c0246e2ae05685efa25000f81809a9364"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("18922794", 0)),
						GasLimit:     30000000,
						GasUsed:      12054372,
						Timestamp:    1704236039,
						BaseFee:      lo.Must(new(big.Int).SetString("17398960391", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715"),
						From:      common.HexToAddress("0x7205A3526Ed4F762B61c804e054A499c0f102e1C"),
						Gas:       426931,
						GasPrice:  lo.Must(new(big.Int).SetString("17498960391", 10)),
						Hash:      common.HexToHash("0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715"),
						Input:     hexutil.MustDecode("0x8c3152e900000000000000000000000000000000000000000000000000000000000000200001000000000000000000000000000000000000000000000000000000001268000000000000000000000000420000000000000000000000000000000000000700000000000000000000000025ace71c97b33cc4729cf772ae268934f7ab5fa1000000000000000000000000000000000000000000000000002421df2cb6fe66000000000000000000000000000000000000000000000000000000000004638800000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001a4d764ad0b0001000000000000000000000000000000000000000000000000000000001268000000000000000000000000420000000000000000000000000000000000001000000000000000000000000099c9fc46f92e8a1c0dec1b1747d010903e884be1000000000000000000000000000000000000000000000000002421df2cb6fe66000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000a41635f5fd0000000000000000000000007205a3526ed4f762b61c804e054a499c0f102e1c0000000000000000000000007205a3526ed4f762b61c804e054a499c0f102e1c000000000000000000000000000000000000000000000000002421df2cb6fe66000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa2f39abba7db061bc5653af3f55e72ea79cb4719cbe301523f7c266428ab969d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18922794", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 9539950,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x41304f207"),
						GasUsed:           182081,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631"),
									common.HexToHash("0x0000000000000000000000007205a3526ed4f762b61c804e054a499c0f102e1c"),
									common.HexToHash("0x0000000000000000000000007205a3526ed4f762b61c804e054a499c0f102e1c"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000002421df2cb6fe6600000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("18922794", 0)),
								TransactionHash: common.HexToHash("0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715"),
								Index:           275,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"),
								Topics: []common.Hash{
									common.HexToHash("0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d"),
									common.HexToHash("0x0000000000000000000000007205a3526ed4f762b61c804e054a499c0f102e1c"),
									common.HexToHash("0x0000000000000000000000007205a3526ed4f762b61c804e054a499c0f102e1c"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000002421df2cb6fe6600000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("18922794", 0)),
								TransactionHash: common.HexToHash("0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715"),
								Index:           276,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1"),
								Topics: []common.Hash{
									common.HexToHash("0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c"),
									common.HexToHash("0x449c49fc02b1e10ed800513df59448d28754a4ddea85a8faf649633a8f43e0b3"),
								},
								Data:            nil,
								BlockNumber:     lo.Must(new(big.Int).SetString("18922794", 0)),
								TransactionHash: common.HexToHash("0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715"),
								Index:           277,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed"),
								Topics: []common.Hash{
									common.HexToHash("0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b"),
									common.HexToHash("0x20d3874f883c863b5a1d614c87bec199c2331c2f0d3317d6c8bfee7236c6722b"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001"),
								BlockNumber:     lo.Must(new(big.Int).SetString("18922794", 0)),
								TransactionHash: common.HexToHash("0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715"),
								Index:           278,
								Removed:         false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715"),
						TransactionIndex: 103,
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x7d636b1cbc52a7b9b0836dad3536869f5b4878d8ca250ac4f0ca0687d695c715",
				Network: network.Ethereum,
				Index:   103,
				From:    "0x7205A3526Ed4F762B61c804e054A499c0f102e1C",
				To:      optimism.AddressL1OptimismPortal.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x8c3152e9",
				},
				Platform: workerx.PlatformOptimism.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3186228206953671")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformOptimism.String(),
						From:     "0x7205A3526Ed4F762B61c804e054A499c0f102e1C",
						To:       "0x7205A3526Ed4F762B61c804e054A499c0f102e1C",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Optimism,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("10170341573197414"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1704236039,
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config)
			require.NoError(t, err)

			matched, err := instance.Match(ctx, testcase.arguments.task)
			testcase.wantError(t, err)
			require.True(t, matched)

			activity, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			// t.Log(string(lo.Must(json.MarshalIndent(activity, "", "\x20\x20"))))

			require.Equal(t, testcase.want, activity)
		})
	}
}
