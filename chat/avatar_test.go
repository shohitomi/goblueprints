package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	_, err := authAvatar.AvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.AvatarURL should return ErrNoAvatarURL when no value present")
	}

	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err := authAvatar.AvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.AvatarURL should return no error when value present")
	}
	if url != testUrl {
		t.Error("AuthAvatar.AvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346",
	}
	url, err := gravatarAvatar.AvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.AvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvatar.AvatarURL wrongly returned %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	os.WriteFile(filename, []byte{}, 0644)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "abc"}
	url, err := fileSystemAvatar.AvatarURL(client)
	if err != nil {
		t.Error("FileSystemAvatar.AvatarURL should not return an error")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.AvatarURL wrongly returned %s", url)
	}
}