package sql

type User struct {
	UserID        int64  `gorm:"cloumn:user_id;primaryKey"`
	UserNickName  string `gorm:"cloumn:user_nick_name"`
	FollowCount   int64  `gorm:"cloumn:follow_count"`
	FollowerCount int64  `gorm:"cloumn:follower_count"`
}

type User_inf struct {
	User     User
	IsFollow bool
}

func CheckUserInf(UserID int, FollowerID int) (u User_inf, exist bool) {
	db, _ := SqlConnect()
	result := db.Table("user_info").Where("user_id = ?", UserID).Find(&u.User)
	if result.RowsAffected == 0 {
		return u, false
	}
	u.IsFollow = CheckIsFollow(UserID, FollowerID)

	return u, true
}

func CheckIsFollow(UserID int, FollowerID int) bool {
	var num int64
	db, _ := SqlConnect()
	db.Table("follow_and_follower_list").Where("user_id = ? and follower_id = ?", UserID, FollowerID).Count(&num)

	return num > 0
}

func CreateInfo(UserName string, uid int64) error {
	info := User{UserID: int64(uid), UserNickName: UserName, FollowCount: 0, FollowerCount: 0}
	db, err := SqlConnect()
	db.Table("user_info").Create(&info)

	return err
}
