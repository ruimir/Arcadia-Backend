// Code generated by entc, DO NOT EDIT.

package datafile

const (
	// Label holds the string label denoting the datafile type in the database.
	Label = "datafile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeHeader holds the string denoting the header edge name in mutations.
	EdgeHeader = "header"
	// EdgeGames holds the string denoting the games edge name in mutations.
	EdgeGames = "games"
	// Table holds the table name of the datafile in the database.
	Table = "datafiles"
	// HeaderTable is the table the holds the header relation/edge.
	HeaderTable = "headers"
	// HeaderInverseTable is the table name for the Header entity.
	// It exists in this package in order to avoid circular dependency with the "header" package.
	HeaderInverseTable = "headers"
	// HeaderColumn is the table column denoting the header relation/edge.
	HeaderColumn = "datafile_header"
	// GamesTable is the table the holds the games relation/edge.
	GamesTable = "games"
	// GamesInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GamesInverseTable = "games"
	// GamesColumn is the table column denoting the games relation/edge.
	GamesColumn = "datafile_games"
)

// Columns holds all SQL columns for datafile fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}