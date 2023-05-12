package entities

import "time"

type Address struct {
	Street  string `bson:"street,omitempty"`
	City    string `bson:"city,omitempty"`
	State   string `bson:"state,omitempty"`
	Country string `bson:"country,omitempty"`
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
