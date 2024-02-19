package http

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBorrowingRecordHandler(t *testing.T) {
	dep := NewBorrowingRecordHandler(nil, nil)
	assert.NotNil(t, dep)
}
func TestGetAllBorrowed(t *testing.T) {
	type fields struct {
		borrowingRecordUsecase *mocks.BorrowingRecordUsecase
		bookUsecase            *mocks.BookUsecase
	}

	tests := []struct {
		name        string
		body        io.Reader
		mock        func(ctx *gin.Context, f fields)
		wantCode    int
		wantMessage string
	}{
		{
			name: "success 1",
			body: nil,
			mock: func(ctx *gin.Context, f fields) {
				f.borrowingRecordUsecase.On(
					"GetAllBorrowedRecord",
					ctx,
					mock.Anything,
				).Return(
					[]domain.BorrowingRecord{
						{
							Id:            1,
							UserId:        1,
							BookId:        22,
							Status:        "BORROWED",
							BorrowingDate: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
							ReturningDate: sql.NullTime{},
							CreatedAt:     time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
							UpdatedAt:     time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
							DeletedAt:     sql.NullTime{},
						},
					},
					nil,
				)
			},
			wantCode:    http.StatusOK,
			wantMessage: `{"message":"OK","data":[{"id":1,"user_id":1,"book_id":22,"status":"BORROWED","borrowing_date":"2020-04-11T21:34:01Z","returning_date":{"Time":"0001-01-01T00:00:00Z","Valid":false},"created_at":"2020-04-11T21:34:01Z","updated_at":"2020-04-11T21:34:01Z","deleted_at":{"Time":"0001-01-01T00:00:00Z","Valid":false}}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDep := fields{
				borrowingRecordUsecase: mocks.NewBorrowingRecordUsecase(t),
				bookUsecase:            mocks.NewBookUsecase(t),
			}

			h := &BorrowingRecordHandler{
				BorrowingRecordUsecase: mockDep.borrowingRecordUsecase,
				BookUsecase:            mockDep.bookUsecase,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			r := httptest.NewRequest("GET", "/borrowing-records/borrowed", tt.body)
			c.Request = r

			tt.mock(c, mockDep)

			h.GetAllBorrowed(c)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantMessage, w.Body.String())
		})
	}
}
