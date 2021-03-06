// Code generated by entc, DO NOT EDIT.

package rom

const (
	// Label holds the string label denoting the rom type in the database.
	Label = "rom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldCrc holds the string denoting the crc field in the database.
	FieldCrc = "crc"
	// FieldMd5 holds the string denoting the md5 field in the database.
	FieldMd5 = "md5"
	// FieldSha1 holds the string denoting the sha1 field in the database.
	FieldSha1 = "sha1"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// EdgeFile holds the string denoting the file edge name in mutations.
	EdgeFile = "file"
	// Table holds the table name of the rom in the database.
	Table = "roms"
	// GameTable is the table the holds the game relation/edge.
	GameTable = "roms"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "game_rom"
	// FileTable is the table the holds the file relation/edge.
	FileTable = "files"
	// FileInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FileInverseTable = "files"
	// FileColumn is the table column denoting the file relation/edge.
	FileColumn = "file_rom"
)

// Columns holds all SQL columns for rom fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSize,
	FieldCrc,
	FieldMd5,
	FieldSha1,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "roms"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_rom",
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
