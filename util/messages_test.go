package util

import (
	"testing"
)

func TestSerializeMessages(t *testing.T) {
	messages := GenerateSerializedMessages("en")

	if len(messages) == 0 {
		t.Errorf("GenerateSerializedMessages() failed.")
	}
}
