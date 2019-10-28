package service

import (
	"fmt"  
    "os"
    "strings"
    "log"
    "agenda/entity"
)

var my_name, my_password string
var Login_flag bool 
var All_name []string

var log_file *os.File

func GetFlag() bool {
	return Login_flag
}

func Init() {
	entity.Init()
    logFile,err := os.OpenFile("service/agenda.log",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    log_file = logFile
    if err != nil {
        log.Fatalln("open file error !")
    }
	tmp := entity.LN_ReadFromFile()
	if (len(tmp)==0) {
		Login_flag = false
	} else {
		Login_flag = true
		my_name = strings.Replace(tmp[0],"\n","",-1)
	}
}

func RegisterUser(name string, password string, email string, phone string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	i := entity.RegisterUser(name, password, email, phone)
	if (i) {
		debugLog.Println(name, " register successfully!")
	} else {
		debugLog.Println(name, " register failed!")
	}
	defer log_file.Close()
}

func Log_in(name string, password string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp_u, flag, _:= entity.Query_user(name)
	if flag == true {
		my_name = name
		my_password = password
		if (entity.GetPassword(tmp_u) != password) {
			debugLog.Println(name, " log in failed!")
			fmt.Println("Password is wrong!")
		} else {
			debugLog.Println(name, " log in successfully!")
			fmt.Println("Log in successfully!")
		}
	} else {
		debugLog.Println(name, " log in failed!")
		fmt.Println("You don't register!")
	}
	entity.LN_WriteToFile(name)
	defer log_file.Close()
}

func Log_out() {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	debugLog.Println(my_name, " log out successfully!")
	fmt.Println("Log out successfully!")
	entity.Empty_login()
	defer log_file.Close()
}


func Query_user(name string) {
	debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp_u, flag, _ := entity.Query_user(name)
	if !flag {
		debugLog.Println(my_name, " query user ", name, " failed!")
		fmt.Println(name," doesn't exists!")
	} else {
		debugLog.Println(my_name, " query user ", name, " successfully!")
		fmt.Println("Name : ", entity.GetName(tmp_u))
		fmt.Println("Email : ", entity.GetEmail(tmp_u))
		fmt.Println("Phone : ", entity.GetPhone(tmp_u))
	}
	defer log_file.Close()
}


