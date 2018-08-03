package helpers

import (
	"reflect"
	"testing"

	"github.com/getlunaform/lunaform/models"
)

func Test_newHalRscLinks(t *testing.T) {
	tests := []struct {
		name string
		want *models.HalRscLinks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHalRscLinks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHalRscLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHalRootRscLinks(t *testing.T) {
	type args struct {
		ch ContextHelper
	}
	tests := []struct {
		name      string
		args      args
		wantLinks *models.HalRscLinks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLinks := HalRootRscLinks(tt.args.ch); !reflect.DeepEqual(gotLinks, tt.wantLinks) {
				t.Errorf("HalRootRscLinks() = %v, want %v", gotLinks, tt.wantLinks)
			}
		})
	}
}

func TestHalSelfLink(t *testing.T) {
	type args struct {
		links *models.HalRscLinks
		href  string
	}
	tests := []struct {
		name string
		args args
		want *models.HalRscLinks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HalSelfLink(tt.args.links, tt.args.href); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HalSelfLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHalDocLink(t *testing.T) {
	type args struct {
		links       *models.HalRscLinks
		operationId string
	}
	tests := []struct {
		name string
		args args
		want *models.HalRscLinks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HalDocLink(tt.args.links, tt.args.operationId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HalDocLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHalAddCuries(t *testing.T) {
	type args struct {
		ch    ContextHelper
		links *models.HalRscLinks
	}
	tests := []struct {
		name string
		args args
		want *models.HalRscLinks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HalAddCuries(tt.args.ch, tt.args.links); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HalAddCuries() = %v, want %v", got, tt.want)
			}
		})
	}
}
