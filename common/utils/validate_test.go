package utils_test

import (
	"rest-demo/common/utils"
	"rest-demo/models"
	"testing"
)

type ValidateUUIDCondition struct {
	ua     models.UserAuditHistory
	fields []string
}

func Test_ValidateUUID(t *testing.T) {
	var tests = []Test{
		{ValidateUUIDCondition{
			ua: models.UserAuditHistory{
				ID: "asd",
			},
			fields: []string{"ID"},
		}, []utils.ValidateError{
			utils.ValidateError{PropertyPath: "ID",
				Message: "invalid UUID length: 3"},
		}},
		{ValidateUUIDCondition{
			ua: models.UserAuditHistory{
				ID: "514a6eea-f5ba-4a2c-8218-fd342e4bb798",
			},
			fields: []string{"ID"},
		}, []utils.ValidateError{}},
	}
	for _, test := range tests {
		b := utils.ValidateUUID(test.in.(ValidateUUIDCondition).ua, test.in.(ValidateUUIDCondition).fields)
		for i := range b {
			if test.out.([]utils.ValidateError)[i] != b[i] {
				t.Errorf("#%d: Contains(%s,%v)=%v; want %v", i, test.in.(ValidateUUIDCondition).ua, test.in.(ValidateUUIDCondition).fields, b, test.out)
			}
		}
	}
}
