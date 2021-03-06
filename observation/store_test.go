package observation_test

import (
	"github.com/ONSdigital/dp-dataset-exporter/observation"
	"github.com/ONSdigital/dp-dataset-exporter/observation/observationtest"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStore_GetCSVRows(t *testing.T) {

	Convey("Given an store with a mock DB connection", t, func() {

		filter := &observation.Filter{
			InstanceID: "888",
			DimensionFilters: []*observation.DimensionFilter{
				{Name: "age", Options: []*observation.DimensionOption{
					{Option: "29"},
					{Option: "30"}}},
				{
					Name: "sex", Options: []*observation.DimensionOption{
						{Option: "male"},
						{Option: "female"}},
				},
			},
		}

		expectedQuery := "MATCH (i:`_888_Instance`) RETURN i.header as row " +
			"UNION ALL " +
			"MATCH (age:`_888_age`), (sex:`_888_sex`) " +
			"WHERE age.value IN ['29', '30'] " +
			"AND sex.value IN ['male', 'female'] " +
			"WITH age, sex " +
			"MATCH (o:`_888_observation`)-[:isValueOf]->(age), (o:`_888_observation`)-[:isValueOf]->(sex) " +
			"return o.value as row"

		expectedCSVRow := "the,csv,row"

		mockBoltRows := &observationtest.BoltRowsMock{
			CloseFunc: func() error {
				return nil
			},
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				return []interface{}{expectedCSVRow}, nil, nil
			},
		}

		mockedDBConnection := &observationtest.DBConnectionMock{
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				return mockBoltRows, nil
			},
		}

		store := observation.NewStore(mockedDBConnection)

		Convey("When GetCSVRows is called", func() {

			rowReader, err := store.GetCSVRows(filter)

			Convey("The expected query is sent to the database", func() {

				actualQuery := mockedDBConnection.QueryNeoCalls()[0].Query

				So(len(mockedDBConnection.QueryNeoCalls()), ShouldEqual, 1)
				So(actualQuery, ShouldEqual, expectedQuery)
			})

			Convey("There is a row reader returned for the rows given by the database.", func() {
				So(err, ShouldBeNil)
				So(rowReader, ShouldNotBeNil)
			})
		})
	})
}
