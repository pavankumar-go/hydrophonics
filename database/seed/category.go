package seed

import (
	"log"

	"github.com/hydrophonics/controllers"
	"github.com/jinzhu/gorm"
)

// Category seeds categories which plants can choose from
func Category(db *gorm.DB) {
	categories := []string{"leafy_green", "non_leafy", "fruit"}
	for _, category := range categories {
		_, err := controllers.AddCategory(category)
		if err != nil {
			log.Panic("error seeding category :", err)
			return
		}
	}
}
