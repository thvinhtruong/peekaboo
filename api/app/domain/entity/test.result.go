package entity

// Set unique constraint TestClassID, StudentID, StudentID.
type TestResult struct {
	ID          int    `db:"id"`            // Primary key
	TestClassID int    `db:"test_class_id"` // foreign key
	UserID      int    `db:"user_id"`       // foreign key
	EntityCode  int    `db:"entity_code"`   // foreign key
	DateCreated int64  `db:"datecreated"`
	Score       int    `db:"score"`
	Comment     string `db:"comment"`
	ResultNote  string `db:"result_note"`
	Active      int    `db:"active"`

	DateUpdated int64 `db:"dateupdated"`
}
