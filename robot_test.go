package robot

import (
	"reflect"
	"testing"
)

func TestRobot_Send(t *testing.T) {
	type fields struct {
		URI   string
		Token string
	}
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "case 1",
			fields: fields{
				URI:   "https://oapi.dingtalk.com/robot/send",
				Token: "your token here",
			},
			args: args{
				msg: &MsgText{Msgtype: "text", Text: MsgTextPayload{Content: "case 1"}},
			},
			wantErr: false,
			want:    &Result{Errcode: 0, Errmsg: "ok"},
		},
		{
			name: "case 2",
			fields: fields{
				URI:   "https://oapi.dingtalk.com/robot/send",
				Token: "your token here",
			},
			args: args{
				msg: &MsgMarkdown{Msgtype: "markdown", Markdown: MsgMarkdownPayload{Title: "测试2", Text: "```<?php echo 'ok';```"}},
			},
			wantErr: false,
			want:    &Result{Errcode: 0, Errmsg: "ok"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				URI:   tt.fields.URI,
				Token: tt.fields.Token,
			}
			got, err := r.Send(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Robot.Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Robot.Send() = %v, want %v", got, tt.want)
			}
		})
	}
}
