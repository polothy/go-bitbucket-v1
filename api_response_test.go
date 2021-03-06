/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucketv1

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSSHKey_String(t *testing.T) {
	type fields struct {
		ID    int
		Text  string
		Label string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &SSHKey{
				ID:    tt.fields.ID,
				Text:  tt.fields.Text,
				Label: tt.fields.Label,
			}
			if got := k.String(); got != tt.want {
				t.Errorf("SSHKey.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommitsResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Commit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommitsResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommitsResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommitsResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTagsResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTagsResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagsResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTagsResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBranchesResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Branch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBranchesResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranchesResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBranchesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRepositoriesResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepositoriesResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepositoriesResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepositoriesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRepositoryResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepositoryResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepositoryResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepositoryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDiffResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    Diff
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDiffResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDiffResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiffResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSSHKeysResponse(t *testing.T) {
	type args struct {
		r *APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []SSHKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSSHKeysResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSSHKeysResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSSHKeysResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPIResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name string
		args args
		want *APIResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIResponse(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPIResponseWithError(t *testing.T) {
	type args struct {
		r   *http.Response
		err error
	}
	tests := []struct {
		name    string
		args    args
		want    *APIResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAPIResponseWithError(tt.args.r, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAPIResponseWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIResponseWithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBitbucketAPIResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *APIResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBitbucketAPIResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBitbucketAPIResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBitbucketAPIResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
