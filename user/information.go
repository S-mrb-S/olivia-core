package user

import (
	"golang.org/x/oauth2"
)

// Information -> UserProfile
// UserProfile is the user's information retrieved from the client
type UserProfile struct {
	FullName       string        `json:"name"`           // Name -> FullName
	GenrePreferences []string      `json:"movie_genres"`   // MovieGenres -> GenrePreferences
	DislikedMovies []string      `json:"movie_blacklist"` // MovieBlacklist -> DislikedMovies
	ImportantDates []UserReminder `json:"reminders"`      // Reminders -> ImportantDates
	StreamingToken *oauth2.Token `json:"spotify_token"`  // SpotifyToken -> StreamingToken
	StreamingID    string        `json:"spotify_id"`     // SpotifyID -> StreamingID
	StreamingSecret string        `json:"spotify_secret"` // SpotifySecret -> StreamingSecret
}

// Reminder -> UserReminder
// A UserReminder is something the user asked to be remembered
type UserReminder struct {
	ReminderDetails string `json:"reason"` // Reason -> ReminderDetails
	ReminderDate    string `json:"date"`   // Date -> ReminderDate
}

// userInformation -> cachedUserData
var cachedUserData = map[string]UserProfile{}

// ChangeUserInformation -> UpdateUserProfile
// UpdateUserProfile requires the token of the user and a function to update the profile,
// and returns the updated profile.
func UpdateUserProfile(authToken string, profileUpdater func(UserProfile) UserProfile) { // token -> authToken, changer -> profileUpdater
	cachedUserData[authToken] = profileUpdater(cachedUserData[authToken])
}

// SetUserInformation -> StoreUserProfile
// StoreUserProfile sets the user's profile using their authentication token.
func StoreUserProfile(authToken string, profile UserProfile) { // token -> authToken, information -> profile
	cachedUserData[authToken] = profile
}

// GetUserInformation -> RetrieveUserProfile
// RetrieveUserProfile returns the user's profile using their authentication token.
func RetrieveUserProfile(authToken string) UserProfile { // token -> authToken
	return cachedUserData[authToken]
}
