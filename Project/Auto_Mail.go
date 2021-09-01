/*
Author : Pakshal Shahikant Jain 
Date : 28/05/2021
Program/Project : AutoMatic Mail Sending Pogram Sending Through/Using Gmail Account
*/
package main

import (
	"fmt"
	"crypto/tls"
	"gomail"
	"gocron"
)

func SendMail() {
	m := gomail.NewMessage()
	m.SetHeader("From","thechainsmokers78@gmail.com")
	m.SetHeader("To","thechainsmokers78@gmail.com")
	m.SetHeader("Subject","Gomail test subject")
	m.SetBody("text/plain","Jay Ganesh.............")
	m.Attach("Demo.go")
	d := gomail.NewDialer("smtp.gmail.com",587,"thechainsmokers78@gmail.com","Your Gmail Account Password")

	d.TLSConfig = &tls.Config{InsecureSkipVerify : true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	fmt.Println("Jay Ganesh........")
	
	gocron.Every(1).Minutes().Do(SendMail)

	<- gocron.Start()
	return 
}
