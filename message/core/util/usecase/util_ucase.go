package usecase

import (
	r "message/domain/repository"
)


type utilUseCase struct {}

func NewUseCase()r.UtilUseCase{
	return &utilUseCase{}
}

func (u *utilUseCase)PaginationValues(p int16)(page int16){
	if p == 1 || p == 0 {
		page = 0
	} else {
		page = p - 1
	}
	return
}


func (h *utilUseCase)GetNextPage(results int8,pageSize int8,page int16) (nextPage int16){
	if results == 0{
	   nextPage = 0
   }else if results != pageSize{
	   nextPage = 0
   } else{
	   nextPage = page + 1
   }
   return
}