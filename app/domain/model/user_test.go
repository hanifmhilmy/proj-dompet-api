package model

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	type args struct {
		ac AccountData
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{
			name: "set the name for account data",
			args: args{
				ac: AccountData{
					Name: "gon",
				},
			},
			want: &Account{
				name: "gon",
			},
		},
		{
			name: "set the account identifier for account data",
			args: args{
				ac: AccountData{
					ID: 1,
				},
			},
			want: &Account{
				id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.ac); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_GetIdentifier(t *testing.T) {
	type fields struct {
		id          int64
		email       string
		name        string
		createdTime time.Time
		createdBy   int64
		updateTime  time.Time
		updateBy    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "get the identifier correct",
			fields: fields{
				id: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:          tt.fields.id,
				email:       tt.fields.email,
				name:        tt.fields.name,
				createdTime: tt.fields.createdTime,
				createdBy:   tt.fields.createdBy,
				updateTime:  tt.fields.updateTime,
				updateBy:    tt.fields.updateBy,
			}
			if got := a.GetIdentifier(); got != tt.want {
				t.Errorf("Account.GetIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
