package Db

type User struct {
	UserID  uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:100"`
	Passwd  string
	ContNum int
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
	Score int
	Mode  uint
}

type Active_level struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Level_id int
	Ok       int
	MaxNum   int
	MinNum   int
	NowNum   int
	Mode     int
	CtfId    int
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
	LevelId     uint
}

type Ctf struct {
	Name string `gorm:"size:100"`
	ID   uint   `gorm:"primaryKey;autoIncrement"`
}
