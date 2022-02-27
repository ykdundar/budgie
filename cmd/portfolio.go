package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/internal"
	"strings"
)

// flags
var (
	name     string
	currency string
	active   bool
	rename   string
)

// Portfolio struct
type Portfolio struct {
	Id       int
	Name     string
	Currency string
	Active   bool
}

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "portfolio commend creates, updates, deletes and lists portfolios by using sub commends",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			fmt.Sprintf("'%s' is created succesfully", name),
		)

		activeValue := internal.ConvertBoolToInt(active)

		createPortfolio, _ := dataBase.Prepare(
			"INSERT INTO portfolios (name, currency, active) VALUES (?, ?, ?)",
		)

		defer createPortfolio.Close()

		_, insertErr := createPortfolio.Exec(name, currency, activeValue)
		cobra.CheckErr(insertErr)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			fmt.Sprintf("'%s' is updated succesfully", name),
		)
		activeValue := internal.ConvertBoolToInt(active)

		var queryStr []string

		if rename != "" {
			queryStr = append(queryStr, fmt.Sprintf("name='%s'", rename))
		}

		if currency != "" {
			queryStr = append(queryStr, fmt.Sprintf("currency='%s'", currency))
		}

		queryStr = append(queryStr, fmt.Sprintf("active=%d", activeValue))

		updateSql := strings.Join(queryStr[:], ",")

		updatePortfolio, _ := dataBase.Prepare(
			fmt.Sprintf("UPDATE portfolios SET %s WHERE name = '%s'", updateSql, name),
		)

		defer updatePortfolio.Close()

		_, updateErr := updatePortfolio.Exec()

		cobra.CheckErr(updateErr)

	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			fmt.Sprintf("'%s' is deleted succesfully", name),
		)
		deletePortfolio, _ := dataBase.Prepare(
			fmt.Sprintf("DELETE FROM portfolios WHERE name= '%s'", name),
		)
		defer deletePortfolio.Close()

		_, deleteErr := deletePortfolio.Exec()
		cobra.CheckErr(deleteErr)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all portfolios",
	Run: func(cmd *cobra.Command, args []string) {
		records, queryErr := dataBase.Query("SELECT * FROM portfolios")
		cobra.CheckErr(queryErr)

		defer records.Close()

		portfolio := Portfolio{}

		for records.Next() {
			scanErr := records.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Currency, &portfolio.Active)

			cobra.CheckErr(scanErr)

			fmt.Println(portfolio)
		}
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows an active portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		record := dataBase.QueryRow("SELECT * FROM portfolios WHERE name=?", name)
		portfolio := Portfolio{}
		scanErr := record.Scan(&portfolio.Id, &portfolio.Name, &portfolio.Currency, &portfolio.Active)
		if scanErr == sql.ErrNoRows {
			cobra.CheckErr(scanErr)
		}
		fmt.Println(portfolio)
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
	portfolioCmd.AddCommand(createCmd, updateCmd, deleteCmd, listCmd, showCmd)

	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	createCmd.PersistentFlags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	createCmd.PersistentFlags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	createCmd.MarkPersistentFlagRequired("name")

	updateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	updateCmd.PersistentFlags().StringVarP(&rename, "rename", "r", "", "Update portfolio name (required)")
	updateCmd.PersistentFlags().StringVarP(&currency, "currency", "c", "USD", "Portfolio currency")
	updateCmd.PersistentFlags().BoolVarP(&active, "active", "a", true, "Set to true if default portfolio")
	updateCmd.MarkPersistentFlagRequired("name")

	deleteCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name (required)")
	deleteCmd.MarkPersistentFlagRequired("name")

	showCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Portfolio name")
}
