package pack

import (
	"tiktok/cmd/comment/dal/db"
	relationdb "tiktok/cmd/relation/dal/db"
	userdb "tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/commentpart"
)

// 打包成可以直接返回的评论信息
func CommentInfo(commentRaw *db.Comment, user *userdb.User) *commentpart.Comment {
	comment := &commentpart.Comment{
		Id: int64(commentRaw.ID),
		User: &commentpart.User{
			// ID:            int64(user.ID),
			Name: user.Name,
			// FollowCount:   user.FollowCount,
			// FollowerCount: user.FollowerCount,
			IsFollow: false,
		},
		Content:    commentRaw.Content,
		CreateDate: commentRaw.CreatedAt.Format("2023-01-01 18:00:00"),
	}
	return comment
}

func CommentList(currentId int64, comments []*db.Comment, userMap map[int64]*userdb.User, relationMap map[int64]*relationdb.Following) []*commentpart.Comment {
	commentList := make([]*commentpart.Comment, 0)
	for _, commentRaw := range comments {
		commentUser, ok := userMap[commentRaw.UserId]
		if !ok {
			commentUser = &userdb.User{
				// ID:            0,
				Name: "name",
				// FollowCount:   0,
				// FollowerCount: 0,
				// IsFollow:      false,
			}
		}

		var isFollow bool = false
		if currentId != -1 {
			_, ok := relationMap[commentRaw.UserId]
			if ok {
				isFollow = true
			}
		}

		commentList = append(commentList, &commentpart.Comment{
			Id: int64(commentRaw.ID),
			User: &commentpart.User{
				// Id:            int64(commentUser.ID),
				Name: commentUser.Name,
				// FollowCount:   commentUser.FollowCount,
				// FollowerCount: commentUser.FollowerCount,
				IsFollow: isFollow,
			},
			Content:    commentRaw.Content,
			CreateDate: commentRaw.CreatedAt.Format("2023-01-01 18:00:00"),
		})
	}
	return commentList
}
