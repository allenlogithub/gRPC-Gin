package databases

import (
	"context"
	"fmt"
)

type (
	AddFriendRequestRequest struct {
		RequestorUserId int64
		ReceiverUserId  int64
	}

	AddFriendListRequest struct {
		RequestorUserId int64
		ReceiverUserId  int64
	}
)

// func: add a friend request record
func AddFriendRequest(r *AddFriendRequestRequest) error {
	rows, err := connMysql.Query("SELECT COUNT(*) FROM friendrequest WHERE requestor_user_id = ? AND receiver_user_id = ?", r.RequestorUserId, r.ReceiverUserId)
	if err != nil {
		return fmt.Errorf("AddFriendRequestRequest.Exec: %v", err)
	}
	for rows.Next() {
		var len int
		if err := rows.Scan(&len); err != nil {
			return fmt.Errorf("AddFriendRequestRequest.Scan: %v", err)
		}
		if len != 0 {
			return fmt.Errorf("AddFriendRequestRequest.DataExists")
		}
	}
	_, err = connMysql.Exec("INSERT INTO friendrequest (requestor_user_id, receiver_user_id) VALUES (?, ?)", r.RequestorUserId, r.ReceiverUserId)
	if err != nil {
		return fmt.Errorf("AddFriendRequest: %v", err)
	}

	return nil
}

// func: (add a user-friend record) & (delete a friend request record)
func AddFriendList(r *AddFriendListRequest) error {
	rows, err := connMysql.Query("SELECT COUNT(*) FROM friendrequest WHERE requestor_user_id = ? AND receiver_user_id = ?", r.RequestorUserId, r.ReceiverUserId)
	if err != nil {
		return fmt.Errorf("AddFriendListRequest.Exec: %v", err)
	}
	for rows.Next() {
		var len int
		if err := rows.Scan(&len); err != nil {
			return fmt.Errorf("AddFriendListRequest.Scan: %v", err)
		}
		if len == 0 {
			return fmt.Errorf("AddFriendListRequest.DataNotFound")
		}
	}

	ctx := context.Background()
	tx, err0 := connMysql.BeginTx(ctx, nil)
	if err0 != nil {
		return fmt.Errorf("AddFriendListRequest.BeginTx: %v", err0)
	}

	_, err = connMysql.ExecContext(ctx, "INSERT INTO friendlist (user_id, friend_user_id) VALUES (?, ?)", r.RequestorUserId, r.ReceiverUserId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("AddFriendListRequest.INSERT.Forward: %v", err)
	}

	_, err = connMysql.ExecContext(ctx, "INSERT INTO friendlist (user_id, friend_user_id) VALUES (?, ?)", r.ReceiverUserId, r.RequestorUserId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("AddFriendListRequest.INSERT.Backward: %v", err)
	}

	stmt, err1 := connMysql.PrepareContext(ctx, "DELETE FROM friendrequest WHERE requestor_user_id = ? AND receiver_user_id = ?")
	if err1 != nil {
		return fmt.Errorf("AddFriendListRequest.DELETE: %v", err)
	}
	_, err = stmt.ExecContext(ctx, r.RequestorUserId, r.ReceiverUserId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("AddFriendListRequest.DELETE: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("AddFriendListRequest.Commit: %v", err)
	}

	return nil
}
