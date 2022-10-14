package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildPathVars(t *testing.T) {
	newstr := func(s string) *string {
		str := new(string)
		*str = s
		return str
	}
	tests := []struct {
		name    string
		path    string
		wantRes map[string]*string
	}{
		{
			name:    "1",
			path:    "/test/noparams",
			wantRes: map[string]*string{},
		},
		{
			name: "2",
			path: "/test/{message.id}",
			wantRes: map[string]*string{
				"message.id": nil,
			},
		},
		{
			name: "3",
			path: "/test/{message.id}/{message.name=messages/*}",
			wantRes: map[string]*string{
				"message.id":   nil,
				"message.name": newstr("messages/*"),
			},
		},
		{
			name: "4",
			path: "/test/{message.name=messages/*}/books",
			wantRes: map[string]*string{
				"message.name": newstr("messages/*"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildPathVars(tt.path)
			assert.Equal(t, tt.wantRes, got)
		})
	}
}

func Test_replacePath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantRes string
	}{
		{
			name:    "1",
			path:    "/test/noparams",
			wantRes: "/test/noparams",
		},
		{
			name:    "2",
			path:    "/test/{message.id}",
			wantRes: "/test/{message.id}",
		},
		{
			name:    "3",
			path:    "/test/{message.id=test}",
			wantRes: "/test/{message.id:test}",
		},
		{
			name:    "4",
			path:    "/test/{message.name=messages/*}/books",
			wantRes: "/test/{message.name:messages/.*}/books",
		},
		{
			name:    "5",
			path:    "/test/{message.id}/{message.name=messages/*}",
			wantRes: "/test/{message.id}/{message.name:messages/.*}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vars := buildPathVars(tt.path)
			for v, s := range vars {
				if s == nil {
					continue
				}
				tt.path = replacePath(v, *s, tt.path)
			}
			assert.Equal(t, tt.wantRes, tt.path)
		})
	}
}
