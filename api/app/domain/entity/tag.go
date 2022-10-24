package entity

// TestTag allow you to tag the type of test. Here is the concrete list that will be implemented.
// 1. Official Test.
// 2. Personality.
// 3. Homework.
// 4. Entrance Test.
// 5. Ielts Test - Reading.
type Tag struct {
	ID          int    `db:"id"`   // primary key
	Tag         string `db:"tag"`  // tag name, unique
	Info        string `db:"info"` // info about the tag.
	Active      int    `db:"active"`
	DateCreated int64  `db:"datecreated"`
	DateUpdated int64  `db:"dateupdated"`
}
