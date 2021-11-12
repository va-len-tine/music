package model

type Mysql struct {
	Hostname	 string `json:"hostname"`
	Port 		string `json:"port"`
	Dbname       string `json:"dbname"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type Gin struct {
	Host	 string `json:"host"`
	Port 		string `json:"port"`
}

type Log struct{
	Outputfile bool	`json:"outputfile"`
	Logname string	`json:"logname"`
	Logprefix string	`json:"logprefix"`
}

type Local struct {
	Defaultmusic string `json:"defaultmusic"`
	Tips string `json:"tips"`
}

type ConfigServer struct {
	Mysql Mysql `structure:"mysql" json:"mysql" yaml:"mysql"`
	Gin      Gin    `json:"gin"`
	Log Log `json:"log"`
	Local Local `json:"local"`
}
