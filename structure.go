package rwanda

type Province struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Districts []District
}

type District struct{
	Id int `json:"id"`
	Name string `json:"name"`
	ProvinceId int `json:"province_id"`
	Sectors []Sector
}

type Sector struct{
	Id int `json:"id"`
	Name string `json:"name"`
	DistrictId int `json:"district_id"`
	Cells []Cell
}

type Cell struct{
	Id int `json:"id"`
	Name string `json:"name"`
	SectorId int `json:"sector_id"`
	Villages []Village
}

type Village struct{
	Id int `json:"id"`
	Name string `json:"name"`
	CellId int `json:"cell_id"`
}
