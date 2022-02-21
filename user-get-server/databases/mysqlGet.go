package databases

import (
	"fmt"

	proto "user-get-server/proto"
)

type (
	GetFriendListRequest struct {
		UserId int64
	}

	SearchUserRequest struct {
		UserId       int64
		SearchString string
	}

	GetFriendRequestListRequest struct {
		UserId int64
	}
)

// func: get friend list
func GetFriendList(r *GetFriendListRequest) (*proto.GetFriendListReply, error) {
	q := `
	SELECT friendlist.friend_user_id, register.name
	FROM friendlist
		LEFT JOIN register
			ON friendlist.friend_user_id=register.id
	WHERE friendlist.user_id=?
	`
	rows, err := connMysql.Query(q, r.UserId)
	if err != nil {
		return nil, fmt.Errorf("GetFriendList.Query: %v", err)
	}
	rp := proto.GetFriendListReply{}
	for rows.Next() {
		var friendUserName string
		var friendUserId int64
		if err := rows.Scan(&friendUserId, &friendUserName); err != nil {
			return nil, fmt.Errorf("GetFriendList.Scan: %v", err)
		}
		fi := proto.FriendInfo{}
		fi.FriendUserName = friendUserName
		fi.FriendUserId = friendUserId
		rp.Items = append(rp.Items, &fi)
	}

	return &rp, nil
}

// query steps:
// 	search column:register.name contains SearchString
// 	join table:friendlist to filter out user of which is in friend list
// 	join table:friendrequest to filter out user of which has received/ requested a friend request
func SearchUser(r *SearchUserRequest) (*proto.SearchUserReply, error) {
	q := `
	SELECT fl.id, fl.name
	FROM (
		SELECT friendlist.friend_user_id, r.id, r.name, friendlist.user_id
		FROM (
			SELECT register.id, register.name
			FROM register
			WHERE register.name LIKE ?
		)AS r
			LEFT JOIN friendlist ON friendlist.user_id=?
				AND r.id=friendlist.friend_user_id
		WHERE friendlist.friend_user_id IS NULL
		) AS fl
			LEFT JOIN friendrequest ON (
				fl.id=friendrequest.requestor_user_id
				AND friendrequest.receiver_user_id=?
			) OR (
				fl.id=friendrequest.receiver_user_id
				AND friendrequest.requestor_user_id=?
			)
		WHERE requestor_user_id IS NULL
			AND fl.id<>?
	`
	rows, err := connMysql.Query(q, string("%")+r.SearchString+string("%"), r.UserId, r.UserId, r.UserId, r.UserId)
	if err != nil {
		return nil, fmt.Errorf("SearchUser.Query: %v", err)
	}
	rp := proto.SearchUserReply{}
	for rows.Next() {
		var userName string
		var userId int64
		if err := rows.Scan(&userId, &userName); err != nil {
			return nil, fmt.Errorf("SearchUser.Scan: %v", err)
		}
		ui := proto.UserInfo{}
		ui.UserName = userName
		ui.UserId = userId
		rp.Items = append(rp.Items, &ui)
	}

	return &rp, nil
}

func GetFriendRequestList(r *GetFriendRequestListRequest) (*proto.GetFriendRequestListReply, error) {
	q := `
	SELECT register.id, register.name
	FROM friendrequest
		LEFT JOIN register
			ON friendrequest.requestor_user_id=register.id
	WHERE friendrequest.receiver_user_id=?
	`
	rows, err := connMysql.Query(q, r.UserId)
	if err != nil {
		return nil, fmt.Errorf("GetFriendRequestList.Query: %v", err)
	}
	rp := proto.GetFriendRequestListReply{}
	for rows.Next() {
		var userId int64
		var userName string
		if err := rows.Scan(&userId, &userName); err != nil {
			return nil, fmt.Errorf("GetFriendRequestList.Scan: %v", err)
		}
		fr := proto.FriendRequest{}
		fr.UserName = userName
		fr.UserId = userId
		rp.Items = append(rp.Items, &fr)
	}

	return &rp, nil
}
