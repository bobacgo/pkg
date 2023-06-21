package enum

type DimTime uint8

// 时间维度
const (
	DimTimeYear DimTime = iota + 1
	DimTimeMonth
	DimTimeDay
	DimTimeHour
	DimTimeMinute
	DimTimeSecond
)

type DimRegion uint8

// 地区维度
const (
	DimRegionProvince DimRegion = iota + 1
	DimRegionCity
	DimRegionDistrict
)
