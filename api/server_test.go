package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	mockdb "github.com/husky_dusky/simplebank/db/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestStart(t *testing.T) {
	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the Gin router's Run method
	mockRouter := gin.New()
	mockRouter.Handle("GET", "/mock", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "mock"})
	})

	// Create a new server with the mock router
	store := mockdb.NewMockStore(ctrl)
	server := &Server{
		store:  store,
		router: mockRouter,
	}

	// Run the server in a goroutine to avoid blocking
	go func() {
		err := server.Start(":8080")
		require.NoError(t, err)
	}()

	// Perform a simple HTTP request to verify the server is running
	resp, err := http.Get("http://localhost:8080/mock")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
