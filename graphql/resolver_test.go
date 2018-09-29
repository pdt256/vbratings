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
	query := `query { playerRatingQueries { getTopPlayerRatings(year: 2018, gender: \"male\", limit: 10) { player { Name } } } }`
	request := getRequestWithQuery(query)
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
	mutation := `mutation { playerCommands { create(bvbId: 1, name: \"John Doe\", imgUrl: \"\") } }`
	request := getRequestWithQuery(mutation)
	response := httptest.NewRecorder()

	// When
	handler.ServeHTTP(response, request)

	// Then
	expectedBody := `{"data":{"playerCommands":{"create":true}}}`
	assert.Equal(t, expectedBody, response.Body.String())
}

func getRequestWithQuery(query string) *http.Request {
	body := fmt.Sprintf(`{"query": "%s"}`, query)
	bodyReader := strings.NewReader(body)
	request := httptest.NewRequest("POST", "/", bodyReader)
	request.Header.Set("Content-Type", "application/json")
	return request
}
