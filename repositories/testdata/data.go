package testdata

import "github.com/ikiat/work/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID: 1,
		Title: "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum: 9,
	},
	models.Article{
		ID: 2,
		Title: "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum: 4,
	},
	models.Article{
		ID: 3,
		Title: "insert test",
		Contents: "Can I insert data correctly?",
		UserName: "saki",
		NiceNum: 0,
	},
}