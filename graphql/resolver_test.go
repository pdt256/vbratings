package graphql_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pdt256/vbratings/app"
	"github.com/pdt256/vbratings/graphql"
	"github.com/stretchr/testify/assert"
)

func Test_GetTopPlayerRatings(t *testing.T) {
	// Given
	configuration := app.NewConfiguration(":memory:")
	application := app.New(configuration)
	handler := graphql.NewGraphQLHandler(application)
	query := `{"query": "{ getTopPlayerRatings(year: 2018, gender: \"male\", limit: 10) { playerName } }"}`
	body := strings.NewReader(query)
	request := httptest.NewRequest("POST", "/", body)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	// When
	handler.ServeHTTP(response, request)

	// Then
	expectedBody := `{"data":{"getTopPlayerRatings":[]}}`
	assert.Equal(t, expectedBody, response.Body.String())
}
