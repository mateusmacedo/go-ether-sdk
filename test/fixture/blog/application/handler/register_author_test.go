package handler

import (
	"reflect"
	"testing"

	"github.com/mateusmacedo/go-ether-sdk/application/handler"
	"github.com/mateusmacedo/go-ether-sdk/application/message"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/repository"
)

func TestWithAuthorRepository(t *testing.T) {
	type args struct {
		repo repository.AuthorRepository
	}
	tests := []struct {
		name string
		args args
		want RegisterAuthorHandlerOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithAuthorRepository(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAuthorRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRegisterAuthorHandler(t *testing.T) {
	type args struct {
		opts []RegisterAuthorHandlerOption
	}
	tests := []struct {
		name string
		args args
		want handler.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegisterAuthorHandler(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegisterAuthorHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registerAuthorHandler_Handle(t *testing.T) {
	type fields struct {
		repo repository.AuthorRepository
	}
	type args struct {
		m message.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    message.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := registerAuthorHandler{
				repo: tt.fields.repo,
			}
			got, err := h.Handle(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("registerAuthorHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerAuthorHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
