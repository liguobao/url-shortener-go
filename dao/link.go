package dao

import (
	"time"
	utilities "url-shortener-go/utilities"
)

// DBLink 短链表
type DBLink struct {
	ID         int64     `gorm: "column:id;primary_key";json: "id";`
	OpenID     string    `gorm: "column:open_id;type:varchar(100)";json: "openId";`
	ShortCode  string    `gorm: "column:short_code;type:varchar(30)";json :"shortCode";`
	URL        string    `gorm: "column:url;type:varchar(170)";json: "url";`
	CreateTime time.Time `gorm: "column:create_time";json: "createTime";`
	UpdateTime time.Time `gorm: "column:update_time";json: "updateTime";`
}

// TableName sets the insert table name for this struct type
func (m *DBLink) TableName() string {
	return "link"
}

// CreateLink 创建一个Link
func CreateLink(link *DBLink) (err error) {
	err = DB().Create(link).Error
	if link.ID > 0 {
		link.ShortCode = utilities.DecimalToAny(int(link.ID), 62)
		err = DB().Model(&link).Update("ShortCode", link.ShortCode).Error
	}
	return
}

// GetLinkByID 查找一个Link
func GetLinkByID(linkID int64) (link *DBLink, err error) {
	link = new(DBLink)
	err = DB().Where("id = ?", linkID).First(link).Error
	return
}

// GetLinkByURLAndOpenID 查找一个Link
func GetLinkByURLAndOpenID(openID string, url string) (link *DBLink, err error) {
	link = new(DBLink)
	err = DB().Where("open_id = ? and url = ?", openID, url).First(link).Error
	return
}

// ContainsLink 是否存在此记录
func ContainsLink(openID string, url string) bool {
	queryCount := new(int)
	DB().Where("open_id = ? and url = ?", openID, url).Count(&queryCount)
	return *queryCount > 0
}

// GetAllLinks 查找所有的Links
func GetAllLinks() (links []*DBLink, err error) {
	err = DB().Find(&links).Error
	return
}
