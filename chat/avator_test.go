package main

import (
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.AvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.AvatarURL should return ErrNoAvatarURL when no value present")
	}

	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.AvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.AvatarURL should return no error when value present")
	}
	if url != testUrl {
		t.Error("AuthAvatar.AvatarURL should return correct URL")
	}
}
