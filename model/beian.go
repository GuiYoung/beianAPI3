package model

import (
	"gorm.io/gorm"
)

type Beian struct {
	gorm.Model
	SiteName   string `gorm:"type:varchar(200);not null" json:"domainName"`
	SiteDomain string `gorm:"type:varchar(200)" json:"websiteName"`
	SiteEntity string `gorm:"type:varchar(200)" json:"entityClass"`
	SiteClass  string `gorm:"type:varchar(200)" json:"websiteClass"`

	OperatorName string `gorm:"type:varchar(200)" json:"entityName"`
	RecordICP    string `gorm:"type:varchar(200)" json:"psICP"`
	RecordeDept  string `gorm:"type:varchar(200)" json:"ICIssuer"`
	RecordTime   string `gorm:"type:varchar(200)" json:"ICPDate"`
}

// insert beian info
func InsertBeian(beian *Beian) (err error) {
	if err = Db.Create(beian).Error; err != nil {
		return
	}
	return
}

// find beian info by name
func FindBeianByName(domainName string) (beian *Beian, err error) {
	if err = Db.Where("site_domain = ?", domainName).First(&beian).Error; err != nil {
		return
	}
	return
}

// update beian info
func UpdateBeianInfo(beian *Beian) (err error) {
	if err = Db.Model(&Beian{}).Where("site_domain = ?", beian.SiteDomain).Updates(beian).Error; err != nil {
		return
	}
	return
}
