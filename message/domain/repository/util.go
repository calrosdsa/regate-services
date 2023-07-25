package repository


type UtilUseCase interface {
	PaginationValues(page int,size int)(int,int)
}