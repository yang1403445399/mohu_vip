package model

import (
	"github.com/dromara/carbon/v2"
)

type Menu struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentId uint            `gorm:"default:0;size:10" json:"parent_id"`
	Name     string          `gorm:"size:50" json:"name"`
	Path     string          `gorm:"size:255" json:"path"`
	Icon     string          `gorm:"size:50" json:"icon"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type Browse struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Ip       string          `gorm:"size:20" json:"ip"`
	Province string          `gorm:"size:20" json:"province"`
	City     string          `gorm:"size:20" json:"city"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type User struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleId   uint            `gorm:"default:1;size:10" json:"role_id"`
	Username string          `gorm:"size:50" json:"username"`
	Password string          `gorm:"size:50;->:false;<-" json:"password"`
	Nickname string          `gorm:"size:50" json:"nickname"`
	Avatar   string          `gorm:"size:255" json:"avatar"`
	Mobile   string          `gorm:"size:11" json:"mobile"`
	Uuid     string          `gorm:"size:36;unique" json:"uuid"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type Log struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId   uint            `gorm:"default:1;size:10" json:"user_id"`
	User     User            `gorm:"foreignKey:UserId;references:Id" json:"user"`
	Route    string          `gorm:"size:255" json:"route"`
	Method   string          `gorm:"size:10" json:"method"`
	Params   string          `json:"params"`
	Remark   string          `gorm:"size:255" json:"remark"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type Article struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId   uint            `gorm:"default:1;size:10" json:"user_id"`
	CateId   uint            `gorm:"size:10" json:"cate_id"`
	Title    string          `gorm:"size:100" json:"title"`
	Keywords string          `gorm:"size:255" json:"keywords"`
	Intro    string          `gorm:"size:255" json:"intro"`
	Thumb    string          `gorm:"size:255" json:"thumb"`
	Author   string          `gorm:"size:255" json:"author"`
	Source   string          `gorm:"size:255" json:"source"`
	Content  string          `gorm:"" json:"content"`
	Hot      uint            `gorm:"default:0" json:"hot"`
	Top      uint            `gorm:"default:0" json:"top"`
	Hits     uint            `gorm:"default:0" json:"hits"`
	Sort     uint            `gorm:"default:0" json:"sort"`
	Uuid     string          `gorm:"size:36;unique" json:"uuid"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type Basic struct {
	Site  string `gorm:"size:10" json:"site"`
	Label string `gorm:"size:50" json:"label"`
	Value string `json:"value"`
}

type Image struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId   uint            `gorm:"default:1;size:10" json:"user_id"`
	ColumnId uint            `gorm:"size:10" json:"column_id"`
	Name     string          `gorm:"size:255" json:"name"`
	Url      string          `gorm:"size:255" json:"url"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type ImageColumn struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId   uint            `gorm:"default:1;size:10" json:"user_id"`
	Name     string          `gorm:"size:255" json:"name"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type Banner struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	TypeId   uint            `gorm:"default:1;size:10" json:"type_id"`
	Type     BannerType      `gorm:"foreignKey:TypeId;references:Id" json:"type"`
	Name     string          `gorm:"size:255" json:"name"`
	Src      string          `gorm:"size:255" json:"src"`
	Url      string          `gorm:"size:255" json:"url"`
	Sort     uint            `gorm:"default:0" json:"sort"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}

type BannerType struct {
	Id       uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string          `gorm:"size:255" json:"name"`
	CreateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt carbon.DateTime `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
	State    uint            `gorm:"default:1" json:"state"`
}
