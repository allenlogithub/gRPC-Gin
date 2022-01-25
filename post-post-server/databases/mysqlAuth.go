package databases

import (
	"fmt"
)

type (
	RegisterInfoReply struct {
		Id             int64
		Name           string
		HashedPassword string
	}

	RegisterInfoRequest struct {
		Account string
	}
)

func GetRegisterInfo(r *RegisterInfoRequest) (*RegisterInfoReply, error) {
	rows, err := connMysql.Query("SELECT id, name, hashed_password FROM register WHERE account = ? LIMIT 1", r.Account)
	if err != nil {
		return &RegisterInfoReply{
			Id:             0,
			Name:           "",
			HashedPassword: "",
		}, fmt.Errorf("GetRegisterInfo: %v", err)
	}
	var id int64
	var name string
	var hashedPassword string
	for rows.Next() {
		rows.Scan(&id, &name, &hashedPassword)
	}

	rp := RegisterInfoReply{
		Id:             id,
		Name:           name,
		HashedPassword: hashedPassword,
	}

	return &rp, nil
}
