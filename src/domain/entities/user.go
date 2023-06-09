package entities

import "time"

type Address struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}

type User struct {
	Id        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Email     string   `json:"email,omitempty"`
	Password  string   `json:"password,omitempty"`
	BirthDate string   `json:"birthDate,omitempty"`
	Age       int      `json:"age,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// This is a method defined on the `User` struct that sets the `Age` field of the user based on their
// `BirthDate`. Calculate the difference between the current date and the user's birth date in years,
// and then sets the `Age` field to that value. If there is an error parsing the `BirthDate`, it returns the error.
func (u *User) SetAge() (err error) {
	today := time.Now()
	pattern := "2006-01-02"
	birthDate, err := time.Parse(pattern, u.BirthDate)
	if err != nil {
		return
	}
	yearInHours := 8760
	age := today.Sub(birthDate).Hours() / float64(yearInHours)
	u.Age = int(age)
	return
}
