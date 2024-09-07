package Db

type User struct {
	UserID uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100"`
	Passwd string
}

type Admin struct {
	UserID  uint
	Name    string `gorm:"size:100"`
	Passwd  string
	AdminID uint `gorm:"primaryKey;autoIncrement"`
}
type Level struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Score uint
	Mode  uint //难度系数
}

type Images struct {
	ID      uint
	ImageID string
	Port    uint
	Name    string
}

type Container struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	ContainerID string
	Port        uint
	Flag        string
	UserID      uint
}

type Ctf struct {
	Name string `gorm:"size:100"`
	ID   uint   `gorm:"primaryKey;autoIncrement"`
}
