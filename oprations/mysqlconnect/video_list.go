package mysqlconnect

func VideoNum(AuthorID int) (n int64) {
	db, _ := SqlConnect()
	db.Table("video_info").Where("author_id = ?", AuthorID).Count(&n)
	return n
}

func GetVideoList(AuthorID int) (list []VideoInfo) {
	db, _ := SqlConnect()
	n := VideoNum(AuthorID)
	if n == 0 {
		return []VideoInfo{}
	}
	db.Table("video_info").Where("author_id = ?", AuthorID).Find(&list)
	return list
}