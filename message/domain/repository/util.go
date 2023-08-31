package repository


type UtilUseCase interface {
	PaginationValues(page int16) int16
	GetNextPage(results int8, pageSize int8, page int16) (nextPage int16)
}