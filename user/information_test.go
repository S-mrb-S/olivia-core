package user

import "testing"

func TestUserInformation(t *testing.T) {
	StoreUserProfile("1", UserProfile{
		FullName: "Hugo",
	})

	if RetrieveUserProfile("1").FullName != "Hugo" {
		t.Errorf("SetUserInformation and/or RetrieveUserProfile failed.")
	}

	UpdateUserProfile("1", func(information UserProfile) UserProfile {
		information.FullName = "Steve"
		return information
	})

	if RetrieveUserProfile("1").FullName != "Steve" {
		t.Errorf("ChangeUserInformation failed.")
	}
}
