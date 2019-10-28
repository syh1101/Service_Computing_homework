package entity

import (
	"fmt"  
    "regexp"

) 


var uData []User

func Show() {
	fmt.Println(uData)
}
func Init() {

	tmp_u := User_ReadFromFile()
	for i := 0; i < len(tmp_u); i++ {
		uData = append(uData, tmp_u[i])
	}
}


func IsEmail(str string) bool {  
    var b bool  
    b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", str)  
    if false == b {  
        return b  
    }
    return b  
}

func IsCellphone(str string) bool {  
    var b bool  
    b, _ = regexp.MatchString("^1[0-9]{10}$", str)  
    if false == b {  
        return b  
    }  
    return b  
}  

func Query_user(name string) (User,bool, int){
	for i := 0; i< len(uData); i++ {
		if uData[i].Name == name {
			return uData[i], true, i
		}
	}
	return User{"1","2","3","4"}, false, 0
}

func RegisterUser(name string, password string, email string, phone string) bool {
	var user User
	err := false
	if (IsEmail(email) == false) {
		fmt.Println("Email is error!")
		err = true
	}
	if (IsCellphone(phone) == false) {
		fmt.Println("Phone is error!")
		err = true
	}
	if (len(password) < 6) {
		fmt.Println("The length of password can't be less than 6!")
		err = true
	}

	if (err) {
		return false
	}
	user.Name = name
	user.Password = password
	user.Email = email
	user.Phone = phone
	uData = append(uData,user)
	User_WriteToFile(uData)
	fmt.Println("Register successfully!")
	return true
}



