package bank

import (
	"exam/tools"
	"flag"
	"fmt"
	"os"
	"strconv"
)

/**
 *
 * @author:yianzhou
 * @email:yianzhou88@gmail.com
 * @DateTime:2022/4/14 16:45
 *
 **/
type Cli struct {

}

/*1、登录功能，登录可以写死，不用链接数据库，用户名为自己的姓名，密码为123456。
2、转账功能（存钱、取钱）
3、查看余额
4、系统功能说明文档
*/

func(cli *Cli)Run(){
	switch os.Args[1] {
	case "login":
		cli.login()
	case "saveMoney":
		cli.saveMoney()
	case "useMoney":
		cli.useMoney()
	case "checkMoney":
		cli.checkMoney()
	case "help":
		cli.help()
	default:
		fmt.Println("没有对应的功能")
		os.Exit(1)
	}
}
func (cli *Cli)login(){
	login := flag.NewFlagSet("login", flag.ExitOnError)
	name := login.String("name", "zhouhaohui", "name")
	password := login.String("password","123456","password")
	login.Parse(os.Args[2:])
	//fmt.Println(*name)
	//fmt.Println(*password)
	if(os.Args[3]==*name&&os.Args[5]==(*password)){
		fmt.Println("登录成功")
	}else{
		fmt.Println("登录失败")
	}
}
//func createBK (){
//	exist := tools.FileExist("./bank.db")
//	if exist {
//		fmt.Println("数据库已存在")
//		return
//	}
//	balance, err := createBalance([]byte("0"))
//	defer balance.DB.Close()
//	if err != nil{
//		fmt.Println("创建失败")
//		return
//	}
//	fmt.Println("创建成功")
//}

func(cli *Cli)saveMoney(){
	saveMoney := flag.NewFlagSet("saveMoney", flag.ExitOnError)
	money := saveMoney.String("money", "0","要存入的钱")//要存入的钱
	saveMoney.Parse(os.Args[2:])//数据解析
	//fmt.Println(*money)
	balance, err := createBalance([]byte("0"))//创建数据库
	defer balance.DB.Close()//延迟关闭数据库
	if err != nil{
		fmt.Println("银行创建失败")
		return
	}
	 money1 := string(balance.Balance)
	 if money1 == ""{
	 	money1 = "0"
	 }
	 //fmt.Println("余额",string(balance.Balance))

	 money2, err := strconv.Atoi(money1)//从数据库取出来的
	if err != nil{
		fmt.Println(err)
		return
	}
	money3, err1 := strconv.Atoi(*money)//输入的金额转int
	//fmt.Println("money3",money3)
	if err1 != nil{
		fmt.Println(err)
		return
	}

	if(money3<=0){
		fmt.Println("存入的金额不正确")
	}else{
		money4 := strconv.Itoa(money2 + money3)//输入的金额和余额相加
		//fmt.Println("money4",money4)
		err := balance.AddMoney([]byte(money4))
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println("存入的金额为:",money3)
	}
}
func(cli *Cli)useMoney(){
	useMoney := flag.NewFlagSet("useMoney", flag.ExitOnError)
	money := useMoney.String("money", "0","取出的钱")
	useMoney.Parse(os.Args[2:])
	fileExist := tools.FileExist("./bank.db")
	if !fileExist{
		fmt.Println("银行不存在")
		return
	}

	balance, err := createBalance([]byte("0"))
	defer balance.DB.Close()
	if err != nil{
		fmt.Println("创建银行失败")
		return
	}

	money1 := string(balance.Balance)

	money2, err := strconv.Atoi(money1)//从数据库取出来的
	if err != nil{
		fmt.Println(err)
		return
	}
	money3, err1 := strconv.Atoi(*money)
	if err1 != nil{
		fmt.Println(err)
		return
	}
	if(money3<=0||money3>money2){
		fmt.Println("取出的金额不正确")
	}else{
		money4 := strconv.Itoa(money2 - money3)//输入的金额和余额相减
		balance.AddMoney([]byte(money4))
		fmt.Println("取出的金额为:",money3)
	}
}

func(cl *Cli)checkMoney(){
	balance, err := createBalance(nil)
	defer balance.DB.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("余额为:",string(balance.Balance))
}
func(cli *Cli)help(){
	fmt.Println("1、login用于用户的登入","name表示用户登入的用户名,password表示用户登入的密码")
	fmt.Println("2、saveMoney用于存钱","money表示要存入的余额")
	fmt.Println("3、useMoney用于取钱","money表示要取出的余额")
	fmt.Println("4、checkMoney用于查看余额")
}