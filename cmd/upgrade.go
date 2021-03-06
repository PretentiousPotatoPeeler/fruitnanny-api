/*
Copyright © Lucas Kahlert

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"log"
	"os"

	"github.com/f3anaro/fruitnanny/internal/update"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Install available package upgrades",
	Long: `Install all available upgrades for FruitNanny Debian packages.

This operation requires an Internet connection.`,
	Run: func(cmd *cobra.Command, args []string) {
		updater, err := update.NewUpdater(viper.GetString("libexec-dir"))
		if err != nil {
			log.Println("Error:", err)
			os.Exit(1)
		}
		err = updater.Upgrade()
		if err != nil {
			log.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
