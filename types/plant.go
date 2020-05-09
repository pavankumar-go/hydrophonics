package types

// PlantDTO creates a plant
type PlantDTO struct {
	Name           string `json:"name" binding:"required"`
	NumberOfPlants int    `json:"numberOfPlants" binding:"required"`
	CategoryID     uint   `json:"categoryID" binding:"required"`
}

// UpdatePlantDTO updates a plant
// TODO add more fields
type UpdatePlantDTO struct {
	Name           string `json:"name"`
	NumberOfPlants int    `json:"numberOfPlants"`
	CategoryID     uint   `json:"categoryID"`
}
