// Code generated by entc, DO NOT EDIT.

package staff

const (
	// Label holds the string label denoting the staff type in the database.
	Label = "staff"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldLastName holds the string denoting the lastname field in the database.
	FieldLastName = "last_name"
	// FieldFirstName holds the string denoting the firstname field in the database.
	FieldFirstName = "first_name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldPhoneNumber holds the string denoting the phonenumber field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// Table holds the table name of the staff in the database.
	Table = "staffs"
)

// Columns holds all SQL columns for staff fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldRole,
	FieldLastName,
	FieldFirstName,
	FieldGender,
	FieldPhoneNumber,
	FieldDescription,
	FieldUserId,
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

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// RoleValidator is a validator for the "role" field. It is called by the builders before save.
	RoleValidator func(int32) error
	// LastNameValidator is a validator for the "lastName" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// FirstNameValidator is a validator for the "firstName" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	GenderValidator func(int32) error
	// PhoneNumberValidator is a validator for the "phoneNumber" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)
