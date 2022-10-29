// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

var (
	Q                   = new(Query)
	Attachment          *attachment
	Category            *category
	Comment             *comment
	CommentBlack        *commentBlack
	FlywaySchemaHistory *flywaySchemaHistory
	Journal             *journal
	Link                *link
	Log                 *log
	Menu                *menu
	Meta                *meta
	Option              *option
	Photo               *photo
	Post                *post
	PostCategory        *postCategory
	PostTag             *postTag
	Tag                 *tag
	ThemeSetting        *themeSetting
	User                *user
)

func SetDefault(db *gorm.DB) {
	*Q = *Use(db)
	Attachment = &Q.Attachment
	Category = &Q.Category
	Comment = &Q.Comment
	CommentBlack = &Q.CommentBlack
	FlywaySchemaHistory = &Q.FlywaySchemaHistory
	Journal = &Q.Journal
	Link = &Q.Link
	Log = &Q.Log
	Menu = &Q.Menu
	Meta = &Q.Meta
	Option = &Q.Option
	Photo = &Q.Photo
	Post = &Q.Post
	PostCategory = &Q.PostCategory
	PostTag = &Q.PostTag
	Tag = &Q.Tag
	ThemeSetting = &Q.ThemeSetting
	User = &Q.User
}

func Use(db *gorm.DB) *Query {
	return &Query{
		db:                  db,
		Attachment:          newAttachment(db),
		Category:            newCategory(db),
		Comment:             newComment(db),
		CommentBlack:        newCommentBlack(db),
		FlywaySchemaHistory: newFlywaySchemaHistory(db),
		Journal:             newJournal(db),
		Link:                newLink(db),
		Log:                 newLog(db),
		Menu:                newMenu(db),
		Meta:                newMeta(db),
		Option:              newOption(db),
		Photo:               newPhoto(db),
		Post:                newPost(db),
		PostCategory:        newPostCategory(db),
		PostTag:             newPostTag(db),
		Tag:                 newTag(db),
		ThemeSetting:        newThemeSetting(db),
		User:                newUser(db),
	}
}

type Query struct {
	db *gorm.DB

	Attachment          attachment
	Category            category
	Comment             comment
	CommentBlack        commentBlack
	FlywaySchemaHistory flywaySchemaHistory
	Journal             journal
	Link                link
	Log                 log
	Menu                menu
	Meta                meta
	Option              option
	Photo               photo
	Post                post
	PostCategory        postCategory
	PostTag             postTag
	Tag                 tag
	ThemeSetting        themeSetting
	User                user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                  db,
		Attachment:          q.Attachment.clone(db),
		Category:            q.Category.clone(db),
		Comment:             q.Comment.clone(db),
		CommentBlack:        q.CommentBlack.clone(db),
		FlywaySchemaHistory: q.FlywaySchemaHistory.clone(db),
		Journal:             q.Journal.clone(db),
		Link:                q.Link.clone(db),
		Log:                 q.Log.clone(db),
		Menu:                q.Menu.clone(db),
		Meta:                q.Meta.clone(db),
		Option:              q.Option.clone(db),
		Photo:               q.Photo.clone(db),
		Post:                q.Post.clone(db),
		PostCategory:        q.PostCategory.clone(db),
		PostTag:             q.PostTag.clone(db),
		Tag:                 q.Tag.clone(db),
		ThemeSetting:        q.ThemeSetting.clone(db),
		User:                q.User.clone(db),
	}
}

type queryCtx struct {
	Attachment          *attachmentDo
	Category            *categoryDo
	Comment             *commentDo
	CommentBlack        *commentBlackDo
	FlywaySchemaHistory *flywaySchemaHistoryDo
	Journal             *journalDo
	Link                *linkDo
	Log                 *logDo
	Menu                *menuDo
	Meta                *metaDo
	Option              *optionDo
	Photo               *photoDo
	Post                *postDo
	PostCategory        *postCategoryDo
	PostTag             *postTagDo
	Tag                 *tagDo
	ThemeSetting        *themeSettingDo
	User                *userDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Attachment:          q.Attachment.WithContext(ctx),
		Category:            q.Category.WithContext(ctx),
		Comment:             q.Comment.WithContext(ctx),
		CommentBlack:        q.CommentBlack.WithContext(ctx),
		FlywaySchemaHistory: q.FlywaySchemaHistory.WithContext(ctx),
		Journal:             q.Journal.WithContext(ctx),
		Link:                q.Link.WithContext(ctx),
		Log:                 q.Log.WithContext(ctx),
		Menu:                q.Menu.WithContext(ctx),
		Meta:                q.Meta.WithContext(ctx),
		Option:              q.Option.WithContext(ctx),
		Photo:               q.Photo.WithContext(ctx),
		Post:                q.Post.WithContext(ctx),
		PostCategory:        q.PostCategory.WithContext(ctx),
		PostTag:             q.PostTag.WithContext(ctx),
		Tag:                 q.Tag.WithContext(ctx),
		ThemeSetting:        q.ThemeSetting.WithContext(ctx),
		User:                q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
