package databases

import (
	// "errors"
	"fmt"
)

type (
	RegisterInfoReply struct {
		Name           string
		HashedPassword string
	}

	RegisterInfoRequest struct {
		Account string
	}
)

func GetRegisterInfo(r *RegisterInfoRequest) (*RegisterInfoReply, error) {
	rows, err := conn.Query("SELECT name, hashed_password FROM register WHERE account = ? LIMIT 1", r.Account)
	if err != nil {
		return &RegisterInfoReply{
			Name:           "",
			HashedPassword: "",
		}, fmt.Errorf("GetRegisterInfo: %v", err)
	}
	var name string
	var hashedPassword string
	for rows.Next() {
		rows.Scan(&name, &hashedPassword)
	}

	rp := RegisterInfoReply{
		Name:           name,
		HashedPassword: hashedPassword,
	}

	return &rp, nil
}
