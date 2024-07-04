package rwanda

// get all data
func GetAllRegions() *AllData{
	LoadData()
	return Data
}
// get all provinces
func GetAllProvinces() []Province{
	LoadData()
	return Data.Provinces
}
// get a province by id
func GetProvinceById(id string) Province{
	LoadData()
	for _, province := range Data.Provinces{
		if province.Id == id{
			return province
		}
	}
	return Province{}
}
// get all districts
func GetAllDistricts() []District{
	LoadData()
	var districts []District
	for _, province := range Data.Provinces{
		districts = append(districts, province.Districts...)
	}
	return districts
}
// get a district by id
func GetDistrictById(id string) District{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			if district.Id == id{
				return district
			}
		}
	}
	return District{}
}
// get all districts in a province
func GetDistrictsInProvince(provinceId string) []District{
	LoadData()
	for _, province := range Data.Provinces{
		if province.Id == provinceId{
			return province.Districts
		}
	}
	return []District{}
}
// get all sectors
func GetAllSectors() []Sector{
	LoadData()
	var sectors []Sector
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			sectors = append(sectors, district.Sectors...)
		}
	}
	return sectors
}
// get a sector by id
func GetSectorById(id string) Sector{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				if sector.Id == id{
					return sector
				}
			}
		}
	}
	return Sector{}
}
// get all sectors in a district
func GetSectorsInDistrict(districtId string) []Sector{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			if district.Id == districtId{
				return district.Sectors
			}
		}
	}
	return []Sector{}
}
// get all cells
func GetAllCells() []Cell{
	LoadData()
	var cells []Cell
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				cells = append(cells, sector.Cells...)
			}
		}
	}
	return cells
}
// get a cell by id
func GetCellById(id string) Cell{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				for _, cell := range sector.Cells{
					if cell.Id == id{
						return cell
					}
				}
			}
		}
	}
	return Cell{}
}
// get all cells in a sector
func GetCellsInSector(sectorId string) []Cell{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				if sector.Id == sectorId{
					return sector.Cells
				}
			}
		}
	}
	return []Cell{}
}
// get all villages
func GetAllVillages() []Village{
	LoadData()
	var villages []Village
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				for _, cell := range sector.Cells{
					villages = append(villages, cell.Villages...)
				}
			}
		}
	}
	return villages
}
// get a village by id
func GetVillageById(id string) Village{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				for _, cell := range sector.Cells{
					for _, village := range cell.Villages{
						if village.Id == id{
							return village
						}
					}
				}
			}
		}
	}
	return Village{}
}
// get all villages in a cell
func GetVillagesInCell(cellId string) []Village{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				for _, cell := range sector.Cells{
					if cell.Id == cellId{
						return cell.Villages
					}
				}
			}
		}
	}
	return []Village{}
}
// get all villages in a sector
func GetVillagesInSector(sectorId string) []Village{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			for _, sector := range district.Sectors{
				if sector.Id == sectorId{
					var villages []Village
					for _, cell := range sector.Cells{
						villages = append(villages, cell.Villages...)
					}
					return villages
				}
			}
		}
	}
	return []Village{}
}
// get all villages in a district
func GetVillagesInDistrict(districtId string) []Village{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			if district.Id == districtId{
				var villages []Village
				for _, sector := range district.Sectors{
					for _, cell := range sector.Cells{
						villages = append(villages, cell.Villages...)
					}
				}
				return villages
			}
		}
	}
	return []Village{}
}
// get all villages in a province
func GetVillagesInProvince(provinceId string) []Village{
	LoadData()
	for _, province := range Data.Provinces{
		if province.Id == provinceId{
			var villages []Village
			for _, district := range province.Districts{
				for _, sector := range district.Sectors{
					for _, cell := range sector.Cells{
						villages = append(villages, cell.Villages...)
					}
				}
			}
			return villages
		}
	}
	return []Village{}
}
// get all cells in a district
func GetCellsInDistrict(districtId string) []Cell{
	LoadData()
	for _, province := range Data.Provinces{
		for _, district := range province.Districts{
			if district.Id == districtId{
				var cells []Cell
				for _, sector := range district.Sectors{
					cells = append(cells, sector.Cells...)
				}
				return cells
			}
		}
	}
	return []Cell{}
}
// get all cells in a province
func GetCellsInProvince(provinceId string) []Cell{
	LoadData()
	for _, province := range Data.Provinces{
		if province.Id == provinceId{
			var cells []Cell
			for _, district := range province.Districts{
				for _, sector := range district.Sectors{
					cells = append(cells, sector.Cells...)
				}
			}
			return cells
		}
	}
	return []Cell{}
}
// get all sectors in a province
func GetSectorsInProvince(provinceId string) []Sector{
	LoadData()
	for _, province := range Data.Provinces{
		if province.Id == provinceId{
			var sectors []Sector
			for _, district := range province.Districts{
				sectors = append(sectors, district.Sectors...)
			}
			return sectors
		}
	}
	return []Sector{}
}
