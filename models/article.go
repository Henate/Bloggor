package models

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`	//TagID即外键, gorm:index，用于声明这个字段为索引
	Tag   Tag `json:"tag"`	//实际是一个嵌套的struct，它利用TagID与Tag模型相互关联，通过Related进行关联查询

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}


func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface {}) (count int){
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

/*
Preload就是一个预加载器，它会执行两条SQL，
分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，
在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询
 */
func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface {}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface {}) bool {
	db.Create(&Article {
		TagID : data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

/*
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())		//time.Now().Unix()返回当前的时间戳

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())	////time.Now().Unix()返回当前的时间戳

	return nil
}
*/