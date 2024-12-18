package test

import (
	"github.com/adnpa/IM/pkg/common/db/mysql"
	"testing"
)

func TestMysql(t *testing.T) {
	mysql.Close()
}
