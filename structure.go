package rwanda

type Provinces []Province

type Province struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Districts []District `json:"districts"`
}
type Village struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CellID int `json:"cell_id"`
}
type Cell struct {
	Id int `json:"id"`
	Name string `json:"name"`
	SectorID int `json:"sector_id"`
	Villages []Village `json:"villages"`
}
type Sector struct {
	Id int `json:"id"`
	Name string `json:"name"`
	DistrictID int `json:"district_id"`
	Cells []Cell `json:"cells"`
}
type District struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ProvinceID string `json:"province_id"`
	Sectors []Sector `json:"sectors"`
}

type StandaloneEntity struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
