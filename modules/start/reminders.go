package start

import (
	"fmt"
	"strings"
	"time"

	"github.com/MehraB832/olivia_core/util"

	"github.com/MehraB832/olivia_core/user"
)

func init() {
	RegisterModule(Module{
		Action: CheckReminders,
	})
}

// CheckReminders will check the dates of the user's reminder and if they are outdated, remove them
func CheckReminders(token, locale string) {
	reminders := user.RetrieveUserProfile(token).ImportantDates
	var messages []string

	// Iterate through the reminders to check if they are outdated
	for i, reminder := range reminders {
		date, _ := time.Parse("01/02/2006 03:04", reminder.ReminderDate)

		now := time.Now()
		// If the date is today
		if date.Year() == now.Year() && date.Day() == now.Day() && date.Month() == now.Month() {
			messages = append(messages, fmt.Sprintf("“%s”", reminder.ReminderDetails))

			// Removes the current reminder
			RemoveUserReminder(token, i)
		}
	}

	// Send the startup message
	if len(messages) != 0 {
		// If the message is already filled in return.
		if GetMessage() != "" {
			return
		}

		// Set the message with the current reminders
		listRemindersMessage := util.SelectRandomMessage(locale, "list reminders")
		if listRemindersMessage == "" {
			return
		}

		message := fmt.Sprintf(
			listRemindersMessage,
			user.RetrieveUserProfile(token).FullName,
			strings.Join(messages, ", "),
		)
		SetMessage(message)
	}
}

// RemoveUserReminder removes the reminder at a specific index in the user's information
func RemoveUserReminder(token string, index int) {
	user.UpdateUserProfile(token, func(information user.UserProfile) user.UserProfile {
		reminders := information.ImportantDates

		// Removes the element from the reminders slice
		if len(reminders) == 1 {
			reminders = []user.UserReminder{}
		} else {
			reminders[index] = reminders[len(reminders)-1]
			reminders = reminders[:len(reminders)-1]
		}

		// Set the updated slice
		information.ImportantDates = reminders

		return information
	})
}
