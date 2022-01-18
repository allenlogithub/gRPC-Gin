package databases

import (
	"fmt"
)

type (
	Register struct {
		Account        string
		HashedPassword string
		Email          string
		Name           string
	}
)

// function: add a new user
// input: Register
// output: int, error
func AddRegister(r *Register) (int64, error) {
	result, err := mysqlDb.Exec("INSERT INTO register (account, hashed_password, email, name) VALUES (?, ?, ?, ?)", r.Account, r.HashedPassword, r.Email, r.Name)
	if err != nil {
		return 0, fmt.Errorf("AddRegister: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddRegister: %v", err)
	}

	return id, nil
}
