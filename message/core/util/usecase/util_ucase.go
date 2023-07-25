package usecase

import (
	r "message/domain/repository"
)


type utilUseCase struct {}

func NewUseCase()r.UtilUseCase{
	return &utilUseCase{}
}

func (u *utilUseCase)PaginationValues(p int,s int)(page int,size int){
	if s == 0 {
		size = 10
	}else{
		size = s
	}
	if p == 1 || p == 0 {
		page = 0
	} else {
		page = p - 1
	}
	return
}