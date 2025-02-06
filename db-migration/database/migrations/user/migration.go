package user

import (
	"github.com/pius-microservices/piopos-database/database/models/user"
	"log"

	"github.com/spf13/cobra"
)

var UserMigrationCMD = &cobra.Command{
	Use:   "migration:user",
	Short: "User DB migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	UserMigrationCMD.Flags().BoolVarP(&migUp, "up", "u", true, "run migration up")

	UserMigrationCMD.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {

	db, err := NewDB()
	if err != nil {
		return err
	}

	if migDown {
		log.Println("User DB migration down done")
		return db.Migrator().DropTable(&user.Role{}, &user.User{}, &user.RefreshToken{})
	}

	if migUp {
		log.Println("User DB migration up done")
		return db.AutoMigrate(&user.Role{}, &user.User{}, &user.RefreshToken{})
	}

	return nil
}
