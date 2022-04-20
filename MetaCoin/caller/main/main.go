package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// "io/ioutil"
	"math/big"
	"xg-pro/MetaCoin/caller/erc721"
)

const (
	toAddress    = "0x477257735cCEF9C7d42856649c7149865D04eDeb"
	chainId      = 1337
	chtAddress   = "0xb9bD7405797CfFBc7e57309444b4af89c39cA92c" // 合约地址
	privateKey   = "\"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCdY4M0d9nwo+fC\\nEVvzIsMfirtmuzJMT7S+KTWyR5hCYd46QZIc80TuzloOxGaHfv8cpVLY1vxiSzOt\\nkemCXwcWo0ILto7xZ9trSlVWDTsC/yQqyebI2mqNMVfruFTwH/FtRf7H21irf0wE\\n+ayQtRSvmUyKmEfU0bT+gE4LZyIsOaoe458ALZyv2qzGURLxb5XrP2qpBMLIJOdl\\nD0h+i0eKcRihoyHmyYeIHBNM+36ilOZIGpPPu6OHlwP7MS5tvkk4stihSR0BPJd2\\nRXs5H4QUq39UEpniIeOcnqrWsYDsRyCyehD3V4IsZBklRU1im3yet46B9rWt1fKa\\n/99q+06DAgMBAAECggEAdzbORirWOPFd9eDCPjtBgx3JbzoyEh15sWRzmNOkH/wT\\noRrTuvwFZcGF85OOeemXATK0uyy0xRtjTICuWCL0so/80fR496AMUotecizZWx65\\nXXPLTK4scUBD8XeRVsVLUNLo9qdN1bE5erdHn+CZh5zdnklUd35U00WKWBbJiqb+\\n7AWSpLjEjc52qSMEPYj6HqLutaY3IO/Ko8daIpHIG+musLqbDLsemRWVNeyPUk1N\\ncxO9bVIJwxPE4MlS8aDhT7OG2JB3gfS9qUYFfDcF6gwg/kvwCeqqebQwd1bh5iUo\\nlA2ikSi5LEdO9cCZf4KMpZqD6VRBdo3DBi3I9a+TmQKBgQDNUfX7Ftc8sgosNHQq\\ngIzNgjzdaqrMwTHGPeDj4tSY1kCDA5P7dnU6IsvUniR3y8WbbTSRaD1LauA6HCCa\\nMGVKrJf/uUe4T9DbQcxEiHudlE0xNaHMRD6diAQ2CLfzSN6UGECbVyaih6514B8O\\nIpTpI4QouI3IPH8TCw4BF75SpQKBgQDEPMwQiZ7wJeZi4K9Bio5yVMIPsqiCRY/F\\ncAJBI3z+D9DWAYVwItlonmACJFq0ehfuL72sdkwFgEshdJd7rp9l++Ox+bPnSZUz\\nHEiUNiM4uYdVkZEEdRObydKxWbvOlKNsXA9/quPnPAo+I/dZB45Vg3y+TDD0RtrG\\n2aEfyTkcBwKBgAq/urf9dyE3VO8Bg/9hoX10zwjsd6qmhpuRS7/CdUzXEqOJQWsN\\nxY9YmQE1kkUvRtYfPnxKT8MuAjUomdBYrkTdikuaAJY2n4GLDU2dM5OJWw0zeJgA\\nDqMips3JYQ/8haKNLnyzcFNb9Rc4t4d+6frWoCmdl9aezCxMDNyjyJP9AoGAGrlg\\n3Puw7mlq55pLo7RX09AhIXwmIJ+ShcPMOsy3b+39dbd9UZkdXHb9Ai2rQJrD+Yb8\\n4Ki3j6Q6FxNGsexE/uF/z7P+wQevTueSJsT8pPP3LzsEscOz6OpiTA65Wde6Lb6X\\nyb4fDVWK242QWMqDoremXFAJ9qpTjvMCU07W/S8CgYEAs1ZzM7C38xjc5O5BzLCU\\n/qnee8SgyyQOjIi75HCWhM9SRPvoWfQ2/ypCaVzOy1qpV1Uy2dFOrXJbyhGiN4tf\\ncmv8GENvkqgVZDrgTQgQhVta9AZKCPdJC+CAquxUM1uACF9WJL5xWpS2vu+b3U/4\\n0ykG96txtASC+/rK9CZ+9is=\\n-----END PRIVATE KEY-----\"\n"
)

func main() {
	// 连接rpc
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		fmt.Println("rpc Dial err")
		panic(err)
	}
	addr := common.HexToAddress("0x18F687c92CaBF3530C8242F8e9D5a720Aa4fFc03")
	acc0 := common.HexToAddress("0xaD86BF99178295423d16d86b51d39C659EB54719")
	acc9 := common.HexToAddress("0x3b4dA041F3573a90007c761A7daF36C68f46D181")

	// 创建合约对象
	erc, err := erc721.NewErc721(addr, client)
	if err != nil {
		fmt.Println("NewErc721 err")
		panic(err)
	}

	balance0, err := erc.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, acc0)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance0:", balance0)

	// 读取私钥文件
	// keyjson, err := ioutil.ReadFile("/home/weilijie/go/src/xg-pro/MetaCoin/caller/main/1.json")
	// if err != nil {
	// 	panic(err)
	// }
	// // 调用其他方法，需要组装from
	// fromKey, err := keystore.DecryptKey(keyjson, "12345")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	privateKey, err := crypto.HexToECDSA("7ee3209b6640887b180e464b8e713f0935f409bb1d2404cba5d2d92873f33190")
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(chainId))
	if err != nil {
		panic(err)
	}

	trade, err := erc.Mint(&bind.TransactOpts{
		From: acc0,
		Signer: auth.Signer,
	}, acc0, big.NewInt(10))
	if err != nil {
		panic(err)
	}
	fmt.Println(trade)

	balance9, err := erc.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, acc9)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance9:", balance9)
	//

	_,err = erc.SetApprovalForAll(&bind.TransactOpts{
		From: acc0,
		Signer: auth.Signer,
	}, acc9,acc0, true)
	if err != nil {
		panic(err)
	}

}

// 设置签名
// func MakeAuth(addr, pass string) (*bind.TransactOpts, error) {
// 	keystorePath := "/Users/yekai/eth/rungeth/data/keystore"
// 	fileName, err := GetFileName(string([]rune(addr)[2:]), keystorePath)
// 	if err != nil {
// 		fmt.Println("failed to GetFileName", err)
// 		return nil, err
// 	}
//
// 	file, err := os.Open(keystorePath + "/" + fileName)
// 	if err != nil {
// 		fmt.Println("failed to open file ", err)
// 		return nil, err
// 	}
// 	auth, err := bind.NewTransactor(file, pass)
// 	if err != nil {
// 		fmt.Println("failed to NewTransactor  ", err)
// 		return nil, err
// 	}
// 	auth.GasLimit = 300000
// 	return auth, err
// }

func NewKeyStore()  {
	ks := keystore.NewKeyStore("./1.key", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123456"
	account, err := ks.NewAccount(password)
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Address.Hex())
}
