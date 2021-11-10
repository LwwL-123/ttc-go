package test

import (
	"fmt"
	"github.com/LwwL-123/go-substrate-crypto/crypto"
	"github.com/LwwL-123/go-substrate-rpc-client/v3/scale"
	"github.com/LwwL-123/go-substrate-rpc-client/v3/types"
	"github.com/LwwL-123/ttc-go/client"
	"github.com/LwwL-123/ttc-go/expand"
	"github.com/LwwL-123/ttc-go/tx"
	"testing"
)

func Test_Transaction(t *testing.T) {
	// 1. 初始化rpc客户端
	c, err := client.New("ws://127.0.0.1:9944")
	if err != nil {
		t.Fatal(err)
	}
	//2. 如果某些链（例如：chainX)的地址的字节前面需要0xff,则下面这个值设置为false
	//expand.SetSerDeOptions(false)
	from := "5GxaFgy8zsS8GGF5TjCoYW24F6e83bn2Hg77ch13jyqqQBxf"
	to := "5GC2XVjtoPsCGezSQ2DtU3K2HrLU2TwwCmHEgc9Vdsh4QyE3"
	amount := uint64(10000000000)
	//3. 获取from地址的nonce
	acc, err := c.GetAccountInfo(from)
	if err != nil {
		t.Fatal(err)
	}
	nonce := uint64(acc.Nonce)
	//4. 创建一个substrate交易，这个方法满足所有遵循substrate 的交易结构的链
	transaction := tx.NewSubstrateTransaction(from, nonce)
	//5. 初始化metadata的扩张结构
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	//6. 初始化Balances.transfer的call方法
	call, err := ed.BalanceTransferCall(to, amount)
	if err != nil {
		t.Fatal(err)
	}
	/*
		//Balances.transfer_keep_alive  call方法
		btkac,err:=ed.BalanceTransferKeepAliveCall(to,amount)
	*/

	/*
		toAmount:=make(map[string]uint64)
		toAmount[to] = amount
		//...
		//true: user Balances.transfer_keep_alive  false: Balances.transfer
		ubtc,err:=ed.UtilityBatchTxCall(toAmount,false)
	*/

	//7. 设置交易的必要参数
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call) //设置call
	//8. 签名交易
	sig, err := transaction.SignTransaction("0xe0b6b6bd3c798dc93ce711bb5c7e1d7f5c7f4a1bdf71372ca42e2de21d03059b", crypto.Sr25519Type)
	if err != nil {
		t.Fatal(err)
	}
	//9. 提交交易
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		t.Fatal(err)
	}
	//10. txid
	txid := result.(string)
	fmt.Println(txid)
}

func Test_WorkerRegister(t *testing.T) {
	// 1. 初始化rpc客户端
	c, err := client.New("ws://127.0.0.1:9944")
	if err != nil {
		t.Fatal(err)
	}
	//2. 如果某些链（例如：chainX)的地址的字节前面需要0xff,则下面这个值设置为false
	//expand.SetSerDeOptions(false)
	from := "5GxaFgy8zsS8GGF5TjCoYW24F6e83bn2Hg77ch13jyqqQBxf"

	//3. 获取from地址的nonce
	acc, err := c.GetAccountInfo(from)
	if err != nil {
		t.Fatal(err)
	}
	nonce := uint64(acc.Nonce)
	//4. 创建一个substrate交易，这个方法满足所有遵循substrate 的交易结构的链
	transaction := tx.NewSubstrateTransaction(from, nonce)
	//5. 初始化metadata的扩张结构
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	//6. 初始化Balances.transfer的call方法
	//totalStorage := types.U128{
	//	Int: new(big.Int).SetInt64(1000000),
	//}
	//
	//usedStorage := types.U128{
	//	Int: new(big.Int).SetInt64(100),
	//}
	res := "111"
	cu := types.Bytes(res)
	call, err := ed.TestCall(cu)
	if err != nil {
		t.Fatal(err)
	}
	/*
		//Balances.transfer_keep_alive  call方法
		btkac,err:=ed.BalanceTransferKeepAliveCall(to,amount)
	*/

	/*
		toAmount:=make(map[string]uint64)
		toAmount[to] = amount
		//...
		//true: user Balances.transfer_keep_alive  false: Balances.transfer
		ubtc,err:=ed.UtilityBatchTxCall(toAmount,false)
	*/

	//7. 设置交易的必要参数
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call) //设置call
	//8. 签名交易
	sig, err := transaction.SignTransaction("0xe0b6b6bd3c798dc93ce711bb5c7e1d7f5c7f4a1bdf71372ca42e2de21d03059b", crypto.Sr25519Type)
	if err != nil {
		t.Fatal(err)
	}
	//9. 提交交易
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		t.Fatal(err)
	}
	//10. txid
	txid := result.(string)
	fmt.Println(txid)
}

type MyVal struct {
	Value interface{}
}

func (a *MyVal) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		var u uint8
		err = decoder.Decode(&u)
		a.Value = u
	} else if b == 1 {
		var s string
		err = decoder.Decode(&s)
		a.Value = s
	}

	if err != nil {
		return err
	}

	return nil
}

func (a MyVal) Encode(encoder scale.Encoder) error {
	var err1, err2 error

	switch v := a.Value.(type) {
	case uint8:
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(v)
	case string:
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(v)
	default:
		return fmt.Errorf("unknown type %T", v)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

func Test_rpc(t *testing.T) {
	c, err := client.New("ws://127.0.0.1:9944")
	if err != nil {
		t.Fatal(err)
	}

	var result interface{}
	var res string
	res = "1"

	//if !strings.HasPrefix(res, "0x") {
	//	res = "0x" + res
	//}

	cu := types.Bytes(res)

	err = c.C.Client.Call(&result, "virtualMachine_getVirtualMachineInfo", fmt.Sprintf("0x%s", cu))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
