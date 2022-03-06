package entity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserFirstLevelUp5Days(t *testing.T) {
	u := new(User)
	u.JoinedAt = time.Now()

	for _, n := range []int{1, 2, 3, 4, 5} {
		u.JoinedAt = u.JoinedAt.Add(-time.Hour * 24)
		u.Calculate()
		//fmt.Println(n, "Lv ", u.Level, " EXP ", u.Exp, " REQ ", u.RequiredExp)
		if n == 5 {
			assert.Equal(t, int64(2), u.Level)
		} else {
			assert.Equal(t, int64(1), u.Level)
		}
	}
}

func TestUserFirstLevelUp556Messages(t *testing.T) {
	u := new(User)
	u.JoinedAt = time.Now()

	for n := int64(0); n < 557; n++ {
		u.MessageCount = n
		u.Calculate()
		//fmt.Println(n, "Lv ", u.Level, " EXP ", u.Exp, " REQ ", u.RequiredExp)
		if n == 556 {
			assert.Equal(t, int64(2), u.Level)
		} else {
			assert.Equal(t, int64(1), u.Level)
		}
	}
}

func TestUserIcons(t *testing.T) {
	u := new(User)
	u.Level = 31

	fmt.Println(calculateIcons(float64(u.Level)))
}
