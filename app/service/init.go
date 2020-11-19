package service

var (
	CategorySer *categoryService
)

func Init()  {
	CategorySer = NewCategoryService()
}
