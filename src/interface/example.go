package main

import "fmt"

type Stringer interface {
	String() string
}

type Logger interface {
	Log() string
}

type Article struct {
	Title string
	Author string
}

type MySqlConnection struct {
	dataBaseName string
	dataBaseUrl string
	dataBaseUserName string
	dataBasePassword string

}

type SShConnection struct {
	userName string
	url string
	password string
}

func (Connection MySqlConnection) Log() string  {
	if Connection.dataBaseName == "" || Connection.dataBaseUrl=="" || Connection.dataBaseUserName=="" || Connection.dataBasePassword==""  {
		return fmt.Sprintln("Please to create a secure connection , verify that the MySqlConnection data is not empty  ")
	} else {
		return fmt.Sprintln("we can launch a secure connection, congratulation", Connection)
	}

}

func (Connection SShConnection) Log() string  {
	if Connection.userName == "" || Connection.url=="" || Connection.password==""   {
		return fmt.Sprintln("Please to create a secure connection , verify that the SShConnection data is not empty  ")
	} else {
		return fmt.Sprintln("we can launch a secure connection, congratulation", Connection)
	}

}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Radhouen Assakra",
	}
	Connection := MySqlConnection{
		"gomycode",
		"https://www.127.168.10.10",
		"gomycode",
		"123654789/*-+",
	}
	sshConn := SShConnection{
		userName: "Radhouen",
		url:      "192.168.10.126",
		password: "gomycode123654/*-+",
	}
	fmt.Println(a.String())
	fmt.Println(Connection.Log())
	fmt.Println(sshConn.Log())
}
