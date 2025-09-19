package types

import (
	"github.com/dromara/carbon/v2"
)

type UserLoginBody struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Mobile   *string `json:"mobile"`
	Code     *string `json:"code"`
}

type UserLoginData struct {
	Id       uint            `json:"id"`
	RoleId   uint            `json:"role_id"`
	Username string          `json:"username"`
	Password string          `json:"password"`
	Nickname string          `json:"nickname"`
	Avatar   string          `json:"avatar"`
	Mobile   string          `json:"mobile"`
	Uuid     string          `json:"uuid"`
	CreateAt carbon.DateTime `json:"create_at"`
	UpdateAt carbon.DateTime `json:"update_at"`
	State    uint            `json:"state"`
	Token    string          `json:"token"`
}

type BrowseTimeRange struct {
	Start *carbon.Carbon `json:"start"`
	End   *carbon.Carbon `json:"end"`
}

type BrowseQueryParams struct {
	Name        string           `json:"name"`
	Date        string           `json:"date"`
	Current     BrowseTimeRange  `json:"current"`
	Previous    BrowseTimeRange  `json:"previous"`
	SelectSQL   string           `json:"select_sql"`
	GroupBy     string           `json:"group_by"`
	LabelFormat func(int) string `json:"label_format"`
	Length      int              `json:"length"`
	Offset      int              `json:"offset"`
}

type BrowseCroupCount struct {
	Label uint `json:"label"`
	Value uint `json:"value"`
}

type BrowseCountData struct {
	Name  string             `json:"name"`
	Date  string             `json:"date"`
	Trend string             `json:"trend"`
	Count uint               `json:"count"`
	Label []string           `json:"label"`
	Value []uint             `json:"value"`
	Group []BrowseCroupCount `json:"group"`
}

type BrowseProvinceCount struct {
	Province string `json:"province"`
	Count    uint   `json:"count"`
}

type BrowseRegionData struct {
	Rank  uint   `json:"rank"`
	Name  string `json:"name"`
	Value uint   `json:"value"`
}

type ArticleCroupCount struct {
	Label string `json:"label"`
	Value uint   `json:"value"`
}

type ArticleCountData struct {
	Label []string `json:"label"`
	Value []uint   `json:"value"`
}

type ImageColumnBody struct {
	Id   *uint  `json:"id"`
	Name string `json:"name"`
}

type BannerBody struct {
	Id     *uint   `json:"id"`
	TypeId uint    `json:"type_id"`
	Name   string  `json:"name"`
	Src    string  `json:"src"`
	Url    *string `json:"url"`
	Sort   *uint   `json:"sort"`
	State  *uint   `json:"state"`
}
