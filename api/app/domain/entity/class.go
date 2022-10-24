package entity

// Class is the core object of the webapp.
// Class consist of date and general information of a class.
// If a class.Active = 1, the class is still functioning at the moment.
// If a class.Active = 0, the class is shut down.
type Class struct {
	ID           int    `db:"id"` // primary key
	Classname    string `db:"class_name"`
	Info         string `db:"info"`
	Announcement string `db:"announcement"`
	RoomCode     string `db:"room_code"`
	Level        string `db:"level"`
	Active       int    `db:"active"`

	DateCreated int64 `db:"datecreated"`
	DateUpdated int64 `db:"dateupdated"`
	// adding more for on-going time
}

func (s *Class) ClassID(ClassID int) {
	s.ID = ClassID
}
