package util_test

import (
	"strconv"
	"testing"

	"github.com/adnpa/IM/internal/utils"
)

func TestGenerateToken(t *testing.T) {
	token, e, err := utils.GenerateToken(strconv.FormatInt(0, 10))
	t.Log(token)
	t.Log(e)
	t.Log(err)
}
