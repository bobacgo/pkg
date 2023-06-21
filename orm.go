package main

import (
	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

type Region struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

func OrmRegion(regions []*Region) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(regions) > 0 {
			var (
				ps  []string
				pcs [][]string
			)
			for _, v := range regions {
				if v.Province != "" && v.City == "" {
					ps = append(ps, v.Province)
				} else if v.Province != "" && v.City != "" {
					pcs = append(pcs, []string{v.Province, v.City})
				}
			}

			if len(ps) > 0 && len(pcs) > 0 {
				db.Where("province IN ? OR (province, city) IN ?", ps, pcs)
			} else if len(pcs) > 0 {
				db.Where("(province, city) IN ?", pcs)
			} else if len(ps) > 0 {
				db.Where("province IN ?", ps)
			}
		}
		return db
	}
}
