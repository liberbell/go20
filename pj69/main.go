package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

func (u User) UnmarshalJSON(b, []byte) error {
	v, err
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@example.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)

	u2 := new(User)
	u3 := new(User)

	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(bs, u3); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
	fmt.Println(u3)

}
