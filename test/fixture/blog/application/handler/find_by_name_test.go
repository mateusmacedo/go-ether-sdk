package handler_test

import (
	"reflect"
	"testing"

	apperr "github.com/mateusmacedo/go-ether-sdk/application/err"
	appmsg "github.com/mateusmacedo/go-ether-sdk/application/message"
	"github.com/mateusmacedo/go-ether-sdk/application/test/helper"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/application/handler"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/application/message"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/model"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/repository"
	mockRep "github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/mocks/domain/repository"
)

func Test_findByNameHandler_Handle(t *testing.T) {
	type fields struct {
		repo repository.FindByName
	}
	type args struct {
		m appmsg.Message
	}
	var findResult = repository.FindByNameResult{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    appmsg.Message
		wantErr bool
		errWant error
	}{
		{
			name: "Test should return error when message is nil",
			fields: fields{
				repo: func() repository.FindByName {
					repoImpl := mockRep.NewFindByName(t)
					return repoImpl
				}(),
			},
			args: args{
				m: nil,
			},
			want:    nil,
			wantErr: true,
			errWant: apperr.ErrMessageContentEmpty,
		},
		{
			name: "Test should return error when message is empty",
			fields: fields{
				repo: func() repository.FindByName {
					repoImpl := mockRep.NewFindByName(t)
					return repoImpl
				}(),
			},
			args: args{
				m: message.NewFindByNameMessage(""),
			},
			want:    nil,
			wantErr: true,
			errWant: apperr.ErrMessageContentEmpty,
		},
		{
			name: "Test should return error when not type of valid message",
			fields: fields{
				repo: func() repository.FindByName {
					repoImpl := mockRep.NewFindByName(t)
					return repoImpl
				}(),
			},
			args: args{
				m: &helper.InvalidMessage{},
			},
			want:    nil,
			wantErr: true,
			errWant: apperr.ErrMessageNotSupported,
		},
		{
			name: "Test should return error when call find by name on repository and return error",
			fields: fields{
				repo: func() repository.FindByName {
					repoImpl := mockRep.NewFindByName(t)
					repoImpl.On("FindByName", "name").Return(nil, apperr.ErrInternalHandler)
					return repoImpl
				}(),
			},
			args: args{
				m: message.NewFindByNameMessage("name"),
			},
			want:    nil,
			wantErr: true,
			errWant: apperr.ErrInternalHandler,
		},
		{
			name: "Test should return message record when call find by name on repository err is nil",
			fields: fields{
				repo: func() repository.FindByName {
					findResult = append(findResult, model.NewAuthor(
						model.WithID(1),
						model.WithVersion(1),
						model.WithName("name"),
					), model.NewAuthor(
						model.WithID(2),
						model.WithVersion(1),
						model.WithName("name"),
					))
					repoImpl := mockRep.NewFindByName(t)
					repoImpl.On("FindByName", "name").Return(findResult, nil)
					return repoImpl
				}(),
			},
			args: args{
				m: message.NewFindByNameMessage("name"),
			},
			want:    message.NewFindByNameMessageResult(findResult),
			wantErr: false,
			errWant: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(tR *testing.T) {
			sut := handler.NewFindByNameHandler(handler.WithFindByNameRepository(tt.fields.repo))

			got, err := sut.Handle(tt.args.m)

			if (err != nil) != tt.wantErr {
				tR.Errorf("findByNameHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err != tt.errWant {
				tR.Errorf("registerAuthorHandler.Handle() error = %v, wantErr %v", err, tt.errWant)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				tR.Errorf("findByNameHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
