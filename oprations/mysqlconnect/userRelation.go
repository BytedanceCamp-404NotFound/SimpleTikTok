package mysqlconnect

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type follow_and_follower_list struct {
	UserID     int64     `gorm:"cloumn:user_id;"`
	FollowerId int64     `gorm:"cloumn:follower_id;"`
	RecordTime time.Time `gorm:"cloumn:record_time;"`
}
type RelationUser struct {
	Id            int64  `from:"id" gorm:"column:user_id"`
	Name          string `from:"name" gorm:"column:user_nick_name"`
	FollowCount   int64  `form:"follow_count" gorm:"column:follow_count"`
	FollowerCount int64  `from:"follower_count" gorm:"column:follower_count"`
	IsFollow      bool   `from:"is_follow"`
}

func RelationAction(UserID int64, ToUserID int64, ActionType int32) (ok bool, err error) {
	db := GormDB
	if ActionType == 1 { //关注
		//#关注账号 user_id：被关注的账号  follower_id：哪个账号要关注
		insertData := follow_and_follower_list{UserID: UserID, FollowerId: ToUserID, RecordTime: time.Now()}
		if err := db.Table("follow_and_follower_list").Create(&insertData).Error; err != nil {
			return false, err
		}

		if err := db.Table("user_info").Where("user_id = ?", UserID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return true, nil
		}

	} else if ActionType == 2 { //取消关注
		//取消关注.  user_id：要被取消关注的账号   follower_id：哪个账号要取消关注

		if err := db.Table("follow_and_follower_list").Where("user_id = ? and follower_id = ?", ToUserID, UserID).Delete(nil).Error; err != nil {
			return false, err
		}

		if err := db.Table("user_info").Where("user_id = ?", UserID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return true, nil
		}
	} else {
		return false, errors.New("无效动作")
	}
	return true, nil

}

func GetFollowerList(LoginUserID int64, UserID int64) (list []RelationUser, err error) {

	testUserList_ := make([]RelationUser, 0)

	tempUserList := RelationUser{}

	db := GormDB
	//#查找某个账号的粉丝列表   user_id：账号
	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM user_info where user_id in(SELECT follower_id FROM follow_and_follower_list where user_id=%d)", UserID)).Rows()

	if err != nil {
		return nil, err
	} else {
		for userInfoList.Next() {
			userInfoList.Scan(&tempUserList.Id, &tempUserList.Name, &tempUserList.FollowCount, &tempUserList.FollowerCount, &tempUserList.IsFollow) //！！err :查询出来的列数不同、数据格式不同时会报错，不影响格式正确的变量
			testUserList_ = append(testUserList_, tempUserList)
		}
		//查询一遍上面查出来的id，是否已被当前登录的账号关注
		for i := 0; i < len(testUserList_); i++ {
			result, _ := CheckIsFollow(int(testUserList_[i].Id), int(LoginUserID))
			if result {
				testUserList_[i].IsFollow = true
			}
		}
	}

	return testUserList_, nil
}

func GetFollowList(LoginUserID int64, UserID int64) (list []RelationUser, err error) {

	testUserList_ := make([]RelationUser, 0)

	tempUserList := RelationUser{}

	db := GormDB
	//#查询某个账号关注的所有账号   follower_id：账号
	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM user_info where user_id in(SELECT user_id FROM follow_and_follower_list where follower_id = %d)", UserID)).Rows()
	if err != nil {
		return nil, nil
	} else {
		for userInfoList.Next() {
			userInfoList.Scan(&tempUserList.Id, &tempUserList.Name, &tempUserList.FollowCount, &tempUserList.FollowerCount, &tempUserList.IsFollow) //！！err :查询出来的列数不同、数据格式不同时会报错，不影响格式正确的变量
			testUserList_ = append(testUserList_, tempUserList)
		}
		//查询一遍上面查出来的id，是否已被当前登录的账号关注
		for i := 0; i < len(testUserList_); i++ {
			result, _ := CheckIsFollow(int(testUserList_[i].Id), int(LoginUserID))
			if result {
				testUserList_[i].IsFollow = true
			}
		}
	}
	return testUserList_, nil
}
