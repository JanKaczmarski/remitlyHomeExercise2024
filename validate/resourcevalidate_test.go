package validate

import "testing"

// ResourceValidate tests
type resourceValidateTest struct {
	jsonFilePath string 
	expected bool
}

var resourceValidateTests = []resourceValidateTest{
	// Resource: "*"
	{"test_files/asteriskResourceAwsIamRolePolicy.json", false},
	// Resource: ["aws1", "aws2"]
	{"test_files/multipleResourceAwsIamRolePolicy.json", true},
	// no Resrouce specified
	{"test_files/noResourceAwsIamRolePolicy.json", true},
}

func TestResourceValidate(t *testing.T){
	for _, test := range resourceValidateTests{
		if output := ResourceValidate(test.jsonFilePath); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}

