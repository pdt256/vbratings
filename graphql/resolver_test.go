package graphql_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pdt256/vbratings/app"
	"github.com/pdt256/vbratings/graphql"
	"github.com/stretchr/testify/assert"
)

func Test_Query_PlayerRatingsQueries_GetTopPlayerRatings(t *testing.T) {
	// Given
	configuration := app.NewConfiguration(":memory:")
	application := app.New(configuration)
	handler := graphql.NewHandler(application)
	query := `query ($year: Int!, $gender: String!, $limit: Int!) {
		playerRatingQueries {
			getTopPlayerRatings(year: $year, gender: $gender, limit: $limit) {
				player {
					Name
				}
			}
		}
	}`
	variables := `{
		"year": 2018,
		"gender": "male",
		"limit": 10
	}`
	request := getRequest(query, variables)
	response := httptest.NewRecorder()

	// When
	handler.ServeHTTP(response, request)

	// Then
	expectedBody := `{"data":{"playerRatingQueries":{"getTopPlayerRatings":[]}}}`
	assert.Equal(t, expectedBody, response.Body.String())
}

func Test_Mutation_PlayerCommands_Create(t *testing.T) {
	// Given
	configuration := app.NewConfiguration(":memory:")
	application := app.New(configuration)
	handler := graphql.NewHandler(application)
	mutation := `mutation ($id: String!, $name: String!, $imgUrl: String!) {
		playerCommands {
			create(Id: $id, name: $name, imgUrl: $imgUrl)
		}
	}`
	variables := `{
		"id": "b0282573b90a4f1591888820b63906f6",
		"name": "John Doe",
		"imgUrl": "http://example.com/1.jpg"
	}`
	request := getRequest(mutation, variables)
	response := httptest.NewRecorder()

	// When
	handler.ServeHTTP(response, request)

	// Then
	expectedBody := `{"data":{"playerCommands":{"create":true}}}`
	assert.Equal(t, expectedBody, response.Body.String())
}

func getRequest(query string, variables string) *http.Request {
	body := fmt.Sprintf(
		`{"query": "%s", "variables": %s}`,
		trimSpaces(query),
		trimSpaces(variables),
	)
	bodyReader := strings.NewReader(body)
	request := httptest.NewRequest("POST", "/", bodyReader)
	request.Header.Set("Content-Type", "application/json")
	return request
}

func trimSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
