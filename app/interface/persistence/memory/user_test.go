package memory

import (
	"database/sql"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	mockDB "github.com/hanifmhilmy/proj-dompet-api/mock/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func Test_userRepository_FindAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	q := `select`
	mockErr := errors.New("err")
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		uname    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		expect  func()
		want    int64
		wantErr bool
	}{
		{
			name: "fail to query",
			args: args{
				uname:    "test",
				password: "ing",
			},
			fields: fields{
				db: sqlxDB,
			},
			expect: func() {
				mock.ExpectQuery(q).WithArgs("test", "ing").WillReturnError(mockErr)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "fail to query - invalid param",
			args: args{
				uname:    "test",
				password: "ing and 1=1",
			},
			fields: fields{
				db: sqlxDB,
			},
			expect: func() {
				mock.ExpectQuery(q).WithArgs("test", "ing and 1=1").WillReturnError(mockErr)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				uname:    "test",
				password: "ing",
			},
			fields: fields{
				db: sqlxDB,
			},
			expect: func() {
				rows := mock.NewRows([]string{"user_id"}).AddRow(1)
				mock.ExpectQuery(q).WithArgs("test", "ing").WillReturnRows(rows)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &userRepository{
				db: database.NewClient(tt.fields.db),
			}
			tt.expect()
			got, err := r.FindAccount(tt.args.uname, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("userRepository.FindAccount() there were unfulfilled expectations: %s", err)
			}
			if got != tt.want {
				t.Errorf("userRepository.FindAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindAccountDetail(t *testing.T) {
	q := `select user_id, name, email, create_time, create_by, update_time, update_by from account_detail where user_id=?`
	mockErr := errors.New("err")
	mockedAccountResult := model.AccountData{
		ID: 1,
	}
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		expect  func(arg args, mockClient *mockDB.MockClient)
		want    *model.AccountData
		wantErr bool
	}{
		{
			name: "fail to query",
			args: args{
				userID: 1,
			},
			expect: func(arg args, mockClient *mockDB.MockClient) {
				mockClient.EXPECT().Select(&model.AccountData{}, q, arg.userID).Return(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail to query - not found",
			args: args{
				userID: 1,
			},
			expect: func(arg args, mockClient *mockDB.MockClient) {
				mockClient.EXPECT().Select(&model.AccountData{}, q, arg.userID).Return(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail to query - invalid param",
			args: args{
				userID: 0,
			},
			expect: func(arg args, mockClient *mockDB.MockClient) {
				mockClient.EXPECT().Select(&model.AccountData{}, q, arg.userID).Return(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				userID: 1,
			},
			expect: func(arg args, mockClient *mockDB.MockClient) {
				ac := model.AccountData{}
				mockClient.EXPECT().Select(&ac, q, arg.userID).SetArg(0, model.AccountData{
					ID: 1,
				}).Return(nil)
			},
			want:    &mockedAccountResult,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockedClient := mockDB.NewMockClient(ctrl)
			r := &userRepository{
				db: mockedClient,
			}
			tt.expect(tt.args, mockedClient)
			got, err := r.FindAccountDetail(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindAccountDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindAccountDetail() = %v, want %v", got, tt.want)
			}

			ctrl.Finish()
		})
	}
}
