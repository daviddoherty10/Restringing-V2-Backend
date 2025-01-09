package database

import "fmt"

func testGet1User(s *service) {
	user, err := s.getUser1()
	if err != nil {
		fmt.Printf("There was an error", err)
	}
	fmt.Printf("User", user)
}
