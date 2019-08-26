package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"ID", "Username", "Password", "Age"}).
		AddRow("1", "foo", "bar", "21").
		AddRow("2", "foo", "bar", "21")
	mock.ExpectQuery("select \\* from users order by id asc").WillReturnRows(rows)

	req, err := http.NewRequest("GET", "http://localhost:8000/api/user", nil)
	rec := httptest.NewRecorder()

	h := FindAllHandler(db)
	h(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status code to be %d, but got: %d", http.StatusOK, rec.Code)
	}
}
