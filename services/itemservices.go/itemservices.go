package itemservices

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Item()
	Save()
}

type itemsService struct{}

func (s *itemsService) Item() {

}

func (s *itemsService) Save() {

}
