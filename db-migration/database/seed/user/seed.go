package user

import (
	"fmt"
	"github.com/pius-microservices/piopos-database/database/migrations/user"
	models "github.com/pius-microservices/piopos-database/database/models/user"
	"log"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type SeedData struct {
	name  string
	model interface{}
	size  int
}

var UserSeedCMD = &cobra.Command{
	Use: "seed:user",
	Short: "For running user seed",
	RunE: Seed,
}

var seedUP bool
var seedDOWN bool

func init()  {
	UserSeedCMD.Flags().BoolVarP(&seedUP, "Up", "u", true, "run seed up")

	UserSeedCMD.Flags().BoolVarP(&seedDOWN, "Down", "d", false, "run seed down")
}

func Seed(cmd *cobra.Command, args []string) error {
	
	var err error

	db, err := user.NewDB()
	if err != nil {
		return err
	}

	if seedDOWN {
		err = seedDown(db)
		return err
	}

	if seedUP {
		err = seedUp(db)
	}

	return err
}

func seedUp(db *gorm.DB) error {
	
	var err error

	var seedModel = []SeedData{
		{
			name: "role",
			model: RoleSeed,
			size: cap(RoleSeed),
		},
	}

	for _, data := range seedModel {

		log.Println("create seeding data for ", data.name)
		err = db.CreateInBatches(data.model, data.size).Error
	}

	return err
}

func seedDown(db *gorm.DB) error {
	
	var err error

	var seedModel = []SeedData{
		{
			// name: models.User{}.TableName(),
			name: models.Role{}.TableName(),
			model: models.Role{},
		},
	}

	for _, data := range seedModel {
		log.Println("Delete seeding data for ", data.name)
		sql := fmt.Sprintf("DELETE FROM %v ", data.name)
		err = db.Exec(sql).Error
	}

	return err
}
