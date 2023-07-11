package service

type ExcelService struct{}

var ExcelServiceImpl = &ExcelService{}

func (*ExcelService) Test() error {

	return nil
}
