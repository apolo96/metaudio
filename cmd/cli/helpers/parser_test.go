package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/apolo96/metaudio/cmd/cli/commands"
	"github.com/apolo96/metaudio/internal/interfaces"
)

type MockClient struct {
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.String(), "/upload") {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("123")),
		}, nil
	}
	if strings.Contains(req.URL.String(), "/request") {		
		url := strings.Split(req.URL.Path,"/")
		value := url[len(url) - 1] 		
		if value != "123" {
			return &http.Response{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader("audiofile id does not exist")),
			}, fmt.Errorf("audiofile id does not exist")
		}
		file, err := os.ReadFile("data_test/audio.json")
		if err != nil {
			return nil, err
		}
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(string(file))),
		}, nil
	}
	return nil, nil
}

func TestParser_Parse(t *testing.T) {
	mockClient := &MockClient{}
	type fields struct {
		commands []interfaces.Command
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "upload failure does not exist",
			fields: fields{
				commands: []interfaces.Command{
					commands.NewUploadCommand(mockClient),
				},
			},
			args: args{
				args: []string{"upload", "-filename", "doesNotExist.mp3"},
			},
			wantErr: true,
		},
		{
			name: "upload success uploaded",
			fields: fields{
				commands: []interfaces.Command{
					commands.NewUploadCommand(mockClient),
				},
			},
			args: args{
				args: []string{"upload", "-filename", "data_test/audio.mp3"},
			},
			wantErr: false,
		},
		{
			name: "get failere id doest not exist",
			fields: fields{
				commands: []interfaces.Command{
					commands.NewGetCommand(mockClient),
				},
			},
			args:    args{args: []string{"get", "-id", "567"}},
			wantErr: true,
		},
		{
			name: "get success requested",
			fields: fields{
				commands: []interfaces.Command{
					commands.NewGetCommand(mockClient),
				},
			},
			args: args{
				args: []string{"get", "-id", "123"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				commands: tt.fields.commands,
			}
			if err := p.Parse(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
