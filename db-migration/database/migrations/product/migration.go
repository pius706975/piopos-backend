package product

import (
	"github.com/pius-microservices/piopos-database/database/models/product"
	"log"

	"github.com/spf13/cobra"
)

var ProductMigrationCMD = &cobra.Command{
	Use:   "migration:product",
	Short: "Product DB migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	ProductMigrationCMD.Flags().BoolVarP(&migUp, "up", "u", true, "run migration up")

	ProductMigrationCMD.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {

	db, err := NewDB()
	if err != nil {
		return err
	}

	if migDown {
		log.Println("Migration down done")
		return db.Migrator().DropTable(&product.Category{}, &product.Supplier{}, &product.Inventory{}, &product.Product{})
	}

	if migUp {
		log.Println("Migration up done")
		return db.AutoMigrate(&product.Category{}, &product.Supplier{}, &product.Inventory{}, &product.Product{})
	}

	return nil
}
