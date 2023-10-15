package models

type Resident struct {
	// Type to store details about a resident
	FirstName string `firestore:"first_name"`
	LastName  string `firestore:"last_name"`
	Email     string `firestore:"email"`
}

type UnitDetail struct {
	// Type for storing details about a particular unit
	Residents []Resident `firestore:"residents"`
	Unit      int16      `firestore:"unit"`
}
