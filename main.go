package rwanda

// get all data
func GetAllRegions() Provinces{
	LoadData()
	return Data
}
// get all provinces
func GetAllProvinces() []StandaloneEntity {
	LoadData()
	var provinces []StandaloneEntity
	for _, province := range Data {
		provinces = append(provinces, StandaloneEntity{
			Id: province.Id,
			Name: province.Name,
		})
	}
	return provinces
}

// get all districts
func GetAllDistricts() []StandaloneEntity{
	LoadData()
	var districts []StandaloneEntity
	for _, province := range Data {
		for _, district := range province.Districts {
			districts = append(districts, StandaloneEntity{
				Id: district.Id,
				Name: district.Name,
			})
		}
	}
	return districts
}

// get all sectors
func GetAllSectors() []StandaloneEntity{
	LoadData()
	var sectors []StandaloneEntity
	for _, province := range Data {
		for _, district := range province.Districts {
			for _, sector := range district.Sectors {
				sectors = append(sectors, StandaloneEntity{
					Id: sector.Id,
					Name: sector.Name,
				})
			}
		}
	}
	return sectors
}

// get all cells
func GetAllCells() []StandaloneEntity{
	LoadData()
	var cells []StandaloneEntity
	for _, province := range Data {
		for _, district := range province.Districts {
			for _, sector := range district.Sectors {
				for _, cell := range sector.Cells {
					cells = append(cells, StandaloneEntity{
						Id: cell.Id,
						Name: cell.Name,
					})
				}
			}
		}
	}
	return cells
}

// get all villages
func GetAllVillages() []StandaloneEntity{
	LoadData()
	var villages []StandaloneEntity
	for _, province := range Data {
		for _, district := range province.Districts {
			for _, sector := range district.Sectors {
				for _, cell := range sector.Cells {
					for _, village := range cell.Villages {
						villages = append(villages, StandaloneEntity{
							Id: village.Id,
							Name: village.Name,
						})
					}
				}
			}
		}
	}
	return villages
}
