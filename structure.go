package rwanda

type Province struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Districts []District
}

type District struct{
	Id string `json:"id"`
	Name string `json:"name"`
	ProvinceId string `json:"province_id"`
	Sectors []Sector
}

type Sector struct{
	Id string `json:"id"`
	Name string `json:"name"`
	DistrictId string `json:"district_id"`
	Cells []Cell
}

type Cell struct{
	Id string `json:"id"`
	Name string `json:"name"`
	SectorId string `json:"sector_id"`
	Villages []Village
}

type Village struct{
	Id string `json:"id"`
	Name string `json:"name"`
	CellId string `json:"cell_id"`
}
