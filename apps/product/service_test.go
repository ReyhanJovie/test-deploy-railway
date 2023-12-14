package product

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	mock := repositoryMock{}
	svc = newService(mock)
}

func TestGetListProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		getListProductMock = func() (products []ProductEntity, err error) {
			return []ProductEntity{
				{
					Id:       1,
					Name:     "Book of Avangers",
					ImageUrl: "",
					Stock:    10,
					Price:    10_000,
				},
			}, nil
		}

		expected := []ProductEntity{
			{
				Id:       1,
				Name:     "Book of Avangers",
				ImageUrl: "",
				Stock:    10,
				Price:    10_000,
			},
		}

		got, err := svc.ListProduct()
		require.Nil(t, err)
		require.Equal(t, len(expected), len(got))
		require.Equal(t, expected[0].Id, got[0].Id)

	})
	t.Run("empty", func(t *testing.T) {

		getListProductMock = func() (products []ProductEntity, err error) {
			return []ProductEntity{}, ErrorProductNotFound
		}

		expected := []ProductEntity{}

		got, err := svc.ListProduct()
		require.NotNil(t, err)
		require.Equal(t, err, ErrorProductNotFound)
		require.Equal(t, len(expected), len(got))

	})
}
