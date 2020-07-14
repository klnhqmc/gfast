// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package user_post

import (
	"gfast/app/model/admin/sys_post"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

//删除对应用户的岗位信息
func DeleteByUserId(userId int64) error {
	_, err := Model.Delete(Columns.UserId, userId)
	return err
}

//添加用户岗位
func AddUserPost(postIds []int64, userId int64) (err error) {
	data := g.List{}
	for _, v := range postIds {
		data = append(data, g.Map{
			Columns.UserId: userId,
			Columns.PostId: v,
		})
	}
	_, err = Model.Data(data).Insert()
	return
}

//获取用户岗位
func GetAdminPosts(userId int) (postIds []int64, err error) {
	list, e := Model.All(Columns.UserId, userId)
	if e != nil {
		g.Log().Error(e)
		err = gerror.New("获取用户岗位信息失败")
		return
	}
	for _, entity := range list {
		postIds = append(postIds, entity.PostId)
	}
	return
}

//根据用户id获取岗位信息详情
func GetPostsByUserId(userId int) ([]*sys_post.Entity, error) {
	model := g.DB().Table(Table)
	datas := ([]*sys_post.Entity)(nil)
	err := model.As("a").InnerJoin("sys_post b", "a.post_id = b.post_id").Fields("b.*").Where(Columns.UserId, userId).Structs(&datas)

	if err != nil {
		return nil, err
	}

	return datas, nil
}
