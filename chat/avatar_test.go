package main

import (
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	_, err := authAvatar.AvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.AvatarURL should return ErrNoAvatarURL when no value present")
	}

	testUrl := "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testUrl, nil)
	url, err := authAvatar.AvatarURL(testChatUser)
	if err != nil {
		t.Error("AuthAvatar.AvatarURL should return no error when value present")
	}
	if url != testUrl {
		t.Error("AuthAvatar.AvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.AvatarURL(user)
	if err != nil {
		t.Error("GravatarAvatar.AvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvatar.AvatarURL wrongly returned %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	os.WriteFile(filename, []byte{}, 0644)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.AvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.AvatarURL should not return an error")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.AvatarURL wrongly returned %s", url)
	}
}
