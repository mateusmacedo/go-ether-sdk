package handler_test

import (
	"testing"

	"github.com/mateusmacedo/go-ether-sdk/application/err"
	"github.com/mateusmacedo/go-ether-sdk/application/message"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/application/handler"
	blogmsg "github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/application/message"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/repository"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/service"
)

func Test_registerAuthorHandler_Handle(t *testing.T) {
	type fields struct {
		srv  service.RegisterNewAuthorService
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
		errWant error
	}{
		{
			name: "Test should return error when not type of valid message",
			fields: fields{
				srv:  func() service.RegisterNewAuthorService { return nil }(),
				repo: func() repository.AuthorRepository { return nil }(),
			},
			args: args{
				m: func() message.Message { return nil }(),
			},
			want:    nil,
			wantErr: true,
			errWant: err.ErrMessageNotSupported,
		},
		{
			name: "Test should return error when content is empty",
			fields: fields{
				srv:  func() service.RegisterNewAuthorService { return nil }(),
				repo: func() repository.AuthorRepository { return nil }(),
			},
			args: args{
				m: blogmsg.NewRegisterAuthor(),
			},
			want:    nil,
			wantErr: true,
			errWant: err.ErrMessageContentEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(tR *testing.T) {
			sut := handler.NewRegisterAuthorHandler(handler.WithAuthorRepository(tt.fields.repo), handler.WithAuthorService(tt.fields.srv))

			got, err := sut.Handle(tt.args.m)

			if (err != nil) != tt.wantErr {
				tR.Errorf("registerAuthorHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err != tt.errWant {
				tR.Errorf("registerAuthorHandler.Handle() error = %v, wantErr %v", err, tt.errWant)
				return
			}

			if got != nil && got != tt.want {
				tR.Errorf("registerAuthorHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
