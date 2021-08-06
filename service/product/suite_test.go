package product

import (
	"context"
	"github.com/pupi94/madara/config"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductSuite struct {
	suite.Suite
	ctx context.Context
}

func TestBaseSuite(t *testing.T) {
	suite.Run(t, new(ProductSuite))
}

func (suite *ProductSuite) SetupSuite() {
	config.InitDB()
}

func (suite *ProductSuite) SetupTest() {
}

func (suite *ProductSuite) TearDownSuite() {
	//ClearDB()
}
