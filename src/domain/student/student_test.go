package student_test

import (
	"testing"

	baseuser "github.com/jeanmolossi/vigilant-waddle/src/domain/base_user"
	"github.com/jeanmolossi/vigilant-waddle/src/domain/student"
	"github.com/stretchr/testify/assert"
)

func TestStudent(t *testing.T) {
	t.Run("should sync received data", func(t *testing.T) {
		wantID := "123"
		wantEmail := "john@doe.com"

		s := student.NewStudent(
			student.WithID(wantID),
			student.WithEmail(wantEmail),
		)

		assert.Equal(t, wantID, s.GetID())
		assert.Equal(t, wantEmail, s.GetEmail())
	})

	t.Run("should return error if options implementation wrong", func(t *testing.T) {
		wrongOption := func(u baseuser.BaseUser) error {
			return baseuser.ErrInvalidBaseUsrImplementation
		}

		have := student.NewStudent().SyncData(wrongOption)
		want := baseuser.ErrInvalidBaseUsrImplementation

		assert.EqualError(t, want, have.Error())
	})

	t.Run("should return an error if SyncData options are empty", func(t *testing.T) {
		have := student.NewStudent().SyncData()
		want := student.ErrNoDataToSync

		assert.EqualError(t, want, have.Error())
	})
}
