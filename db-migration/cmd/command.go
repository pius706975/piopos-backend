package cmd

import (
	"github.com/pius-microservices/piopos-database/database/migrations/product"
	"github.com/pius-microservices/piopos-database/database/migrations/user"

	userSeed "github.com/pius-microservices/piopos-database/database/seed/user"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "DB Migration",
	Long: "Database migration service",
}

func init() {
	initCommand.AddCommand(user.UserMigrationCMD)
	initCommand.AddCommand(product.ProductMigrationCMD)

	initCommand.AddCommand(userSeed.UserSeedCMD)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}