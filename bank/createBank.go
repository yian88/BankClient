package bank

import (
	"errors"
	"fmt"
	"github.com/bolt"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/4/16 17:36
 *
 **/
/**
创建银行
*/
const Bank_DB_PATH = "./bank.db"//要保存的文件地址
const BANK_MONEY = "bank_money"//
const MONEY_KEY = "key"//用于存放余额的key

type Bank struct {
	DB *bolt.DB
	Balance []byte
}

func createBalance(Money []byte)(*Bank,error){
	db, err := bolt.Open(Bank_DB_PATH, 0600, nil)
	if err != nil{
		return nil,err
	}
	var balance []byte
	err = db.Update(func(tx *bolt.Tx) error {
		bKMoney := tx.Bucket([]byte(BANK_MONEY))
		if bKMoney == nil{
			bKMoney, err = tx.CreateBucket([]byte(BANK_MONEY))
			if err != nil{
				fmt.Println(err)
				return err
			}
			bKMoney.Put([]byte(MONEY_KEY),Money)//把钱存入到数据库中
		}else{
			bkMoney1 := tx.Bucket([]byte(BANK_MONEY))

			balance = bkMoney1.Get([]byte(MONEY_KEY))
		}
		return nil
	})
	bk := Bank{
		DB: db,
		Balance: balance,
	}
	return &bk,err
}

func (bk *Bank)AddMoney(money []byte)error{
	err := bk.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BANK_MONEY))
		if bucket == nil{
			return errors.New("没有银行")
		}
		bucket.Put([]byte(MONEY_KEY), money)
		bk.Balance = bucket.Get([]byte(MONEY_KEY))
		return nil
	})
	return err
}