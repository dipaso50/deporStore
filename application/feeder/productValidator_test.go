package feeder

import "testing"

func TestValidFormats(t *testing.T) {

	tt := []string{"KASL-3423", "LPOS-32411"}

	for _, tc := range tt {
		res := validFormat(tc)
		if !res {
			t.Errorf("Expected positive validation for %s", tc)
		}
	}
}

func TestInvalidFormats(t *testing.T) {

	tt := []string{"KASL3423", "LPOS-321", "23", "·2345-rdf", "wsde-3453"}

	for _, tc := range tt {
		res := validFormat(tc)

		if res {
			t.Errorf("Expected negative validation for %s", tc)
		}
	}
}
