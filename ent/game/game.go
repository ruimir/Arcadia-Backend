// Code generated by entc, DO NOT EDIT.

package game

const (
	// Label holds the string label denoting the game type in the database.
	Label = "game"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCloneof holds the string denoting the cloneof field in the database.
	FieldCloneof = "cloneof"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeDatafile holds the string denoting the datafile edge name in mutations.
	EdgeDatafile = "datafile"
	// EdgeReleases holds the string denoting the releases edge name in mutations.
	EdgeReleases = "releases"
	// EdgeRom holds the string denoting the rom edge name in mutations.
	EdgeRom = "rom"
	// Table holds the table name of the game in the database.
	Table = "games"
	// DatafileTable is the table the holds the datafile relation/edge.
	DatafileTable = "games"
	// DatafileInverseTable is the table name for the Datafile entity.
	// It exists in this package in order to avoid circular dependency with the "datafile" package.
	DatafileInverseTable = "datafiles"
	// DatafileColumn is the table column denoting the datafile relation/edge.
	DatafileColumn = "datafile_games"
	// ReleasesTable is the table the holds the releases relation/edge.
	ReleasesTable = "releases"
	// ReleasesInverseTable is the table name for the Release entity.
	// It exists in this package in order to avoid circular dependency with the "release" package.
	ReleasesInverseTable = "releases"
	// ReleasesColumn is the table column denoting the releases relation/edge.
	ReleasesColumn = "game_releases"
	// RomTable is the table the holds the rom relation/edge.
	RomTable = "roms"
	// RomInverseTable is the table name for the Rom entity.
	// It exists in this package in order to avoid circular dependency with the "rom" package.
	RomInverseTable = "roms"
	// RomColumn is the table column denoting the rom relation/edge.
	RomColumn = "game_rom"
)

// Columns holds all SQL columns for game fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCloneof,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "games"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"datafile_games",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
