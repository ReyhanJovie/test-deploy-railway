package product

type repositoryMock struct{}

// list function that will be mock
var (
	getListProductMock func() (products []ProductEntity, err error)
)

// GetListProduct implements Repository.
func (repositoryMock) GetListProduct() (products []ProductEntity, err error) {
	return getListProductMock()
}
