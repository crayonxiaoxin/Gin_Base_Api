package models

import (
	"hello_gin_api/utils"
	"time"
)

// 文章
type Post struct {
	utils.BaseModel
	Uid         uint             `json:"uid" gorm:"not null;default:0"`
	PostTitle   string           `json:"post_title" gorm:"not null"`
	PostContent string           `json:"post_content"`
	PostDate    *utils.LocalTime `json:"post_date" gorm:"not null"`
	PostStatus  string           `json:"post_status" gorm:"not null;default:publish"`
	PostType    string           `json:"post_type" gorm:"not null;default:posts"`
	PostParent  uint             `json:"post_parent" gorm:"default:0"`
}

func init() {
	utils.DB.AutoMigrate(&Post{})
}

func (post *Post) Valid() bool {
	return post.ID > 0
}

type PostListOptions struct {
	utils.ListOptions

	Uid    int
	Date   string
	Status string
	Type   string
}

// 获取所有 posts
func GetPosts(options *PostListOptions) (posts []Post, count int64) {
	options.Prepare()
	tx := utils.DB.Model(Post{})
	if options.Keyword != "" {
		tx = tx.Where("post_title like ? or post_content like ?", options.EscKeyword(), options.EscKeyword())
	}
	if options.Uid > 0 {
		tx = tx.Where("uid = ?", options.Uid)
	}
	if options.Date != "" {
		tx = tx.Where("date(post_date) = ?", options.Date)
	}
	if options.Status != "" {
		tx = tx.Where("post_status = ?", options.Status)
	}
	if options.Type != "" {
		tx = tx.Where("post_type = ?", options.Type)
	}
	tx.Order(options.Order).Count(&count).Limit(options.PageSize).Offset(options.Offset()).Find(&posts)
	return
}

// // 获取所有 posts
// func GetPosts(page int, pageSize int) (posts []Post, count int64) {
// 	if page <= 0 {
// 		page = 1
// 	}
// 	if pageSize <= 0 {
// 		pageSize = 10
// 	}
// 	offset := (page - 1) * pageSize
// 	utils.DB.Model(Post{}).Order("id desc").Count(&count).Limit(pageSize).Offset(offset).Find(&posts)
// 	return
// }

// 通过 id 获取 post
func GetPost(id int) (post *Post) {
	utils.DB.First(&post, id)
	return
}

// 添加 post
func AddPost(post *Post) utils.Result {
	result := utils.Result{}
	if len(post.PostTitle) == 0 {
		result.ResultCode = utils.ERR_PARAMS
	} else {
		if post.PostDate == nil {
			PostDate := utils.LocalTime(time.Now())
			post.PostDate = &PostDate
		}
		if len(post.PostStatus) == 0 {
			post.PostStatus = "publish"
		}
		if len(post.PostType) == 0 {
			post.PostType = "posts"
		}
		err := utils.DB.Create(&post).Error
		if err != nil {
			result.ResultCode = utils.ERR_POST_ADD
		} else {
			result.ResultCode = utils.SUCCESS
			result.Data = *post
		}
	}
	return result
}

// 更新 post
func UpdatePost(post *Post) utils.Result {
	result := utils.Result{}
	if post.ID <= 0 {
		result.ResultCode = utils.ERR_PARAMS
	} else {
		if len(post.PostTitle) == 0 {
			result.ResultCode = utils.ERR_POST_EMPTY_TITLE
		} else {
			getpost := GetPost(int(post.ID))
			if getpost.Valid() {
				post.Uid = getpost.Uid
				if post.PostDate == nil {
					post.PostDate = getpost.PostDate
				}
				if post.PostParent == 0 {
					post.PostParent = getpost.PostParent
				}
				if len(post.PostStatus) == 0 {
					post.PostStatus = getpost.PostStatus
				}
				if len(post.PostType) == 0 {
					post.PostType = getpost.PostType
				}
				err := utils.DB.Updates(&post).Error
				if err != nil {
					result.ResultCode = utils.ERR_POST_UPDATE
				} else {
					result.ResultCode = utils.SUCCESS
					result.Data = *post
				}
			} else {
				result.ResultCode = utils.ERR_POST_NOT_EXISTS
			}
		}
	}

	return result
}

// 删除
func DeletePost(id int) utils.Result {
	var result = utils.Result{}
	if id > 0 {
		post := GetPost(id)
		if post.Valid() {
			err := utils.DB.Delete(&post).Error
			if err != nil {
				result.ResultCode = utils.ERR_POST_DELETE
			} else {
				result.ResultCode = utils.SUCCESS
			}
		} else {
			result.ResultCode = utils.ERR_POST_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}
