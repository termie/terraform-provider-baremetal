// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
)

type DatabaseDatabaseTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDatabaseTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_database_database" "t" {
      database_id = "id"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_database_database.t"
}

func (s *DatabaseDatabaseTestSuite) TestReadDatabase() {
	database := &baremetal.Database{}
	database.DBHomeID = "db_home_id"
	database.ID = "id"

	s.Client.On("GetDatabase", "id").Return(database, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "db_home_id", "db_home_id"),
				),
			},
		},
	},
	)
}

func TestDatabaseDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDatabaseTestSuite))
}
