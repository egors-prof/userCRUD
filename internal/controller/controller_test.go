package controller

import (
	mock_contracts "CSR/internal/contracts/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_DeleteUserById(t *testing.T) {
	type mockBehaviour func(s *mock_contracts.MockServiceI)
	testCases := []struct {
		name                 string
		paramID              string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "valid id -success",
			paramID: "24",
			mockBehaviour: func(s *mock_contracts.MockServiceI) {
				s.EXPECT().DeleteUserById(24).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"info": "user is successfully deleted"}`,
		},
		{
			name:    "negative id",
			paramID: "-9",
			mockBehaviour: func(s *mock_contracts.MockServiceI) {

			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error": "negative id err. id should positive"}`,
		},
		{
			name:    "incorrect id format",
			paramID: "liwhjr",
			mockBehaviour: func(s *mock_contracts.MockServiceI) {

			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"error": "invalid user id format"}`,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := mock_contracts.NewMockServiceI(ctrl)
			testCase.mockBehaviour(service)
			handler := NewController(service)

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.DELETE("/user/:id", handler.DeleteUserById)

			request := httptest.NewRequest(http.MethodDelete, "/users/"+testCase.paramID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}

}
