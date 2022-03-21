package tokens

import (
	"database/sql"
	"github.com/spf13/cobra"
	"github.com/ykdundar/budgie/database"
	"github.com/ykdundar/budgie/internal/objects"
)

func FindToken() string {
	record := database.Database.QueryRow("SELECT * FROM tokens ORDER BY id DESC LIMIT 1;")

	token := objects.Token{}

	scanErr := record.Scan(&token.Id, &token.Token)

	if scanErr == sql.ErrNoRows {
		cobra.CheckErr(scanErr)
	}

	return token.Token
}
