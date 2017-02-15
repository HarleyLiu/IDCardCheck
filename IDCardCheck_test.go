package IDCardCheck

import "testing"

func Test_NumberCheck(t *testing.T) {
	//check	321081196404077524
	err := NumberCheck("321081196404077520")
	if err == nil {
		println("321081196404077520 : check pass")
	}
	if err = NumberCheck("321081196404077521"); err == nil {
		println("321081196404077521 : check pass")
	}
	if err = NumberCheck("321081196404077522"); err == nil {
		println("321081196404077522 : check pass")
	}
	if err = NumberCheck("321081196404077523"); err == nil {
		println("321081196404077523 : check pass")
	}
	if err = NumberCheck("321081196404077524"); err == nil {
		println("321081196404077524 : check pass")
	}
	if err = NumberCheck("321081196404077525"); err == nil {
		println("321081196404077525 : check pass")
	}
	if err = NumberCheck("321081196404077526"); err == nil {
		println("321081196404077526 : check pass")
	}
	if err = NumberCheck("321081196404077527"); err == nil {
		println("321081196404077527 : check pass")
	}
	if err = NumberCheck("321081196404077528"); err == nil {
		println("321081196404077528 : check pass")
	}
	if err = NumberCheck("321081196404077529"); err == nil {
		println("321081196404077529 : check pass")
	}
	if err = NumberCheck("32108119640407752x"); err == nil {
		println("32108119640407752x : check pass")
	}

	//check 32082119760829410x
	if err = NumberCheck("320821197608294100"); err == nil {
		println("320821197608294100 : check pass")
	}
	if err = NumberCheck("320821197608294101"); err == nil {
		println("320821197608294101 : check pass")
	}
	if err = NumberCheck("320821197608294102"); err == nil {
		println("320821197608294102 : check pass")
	}
	if err = NumberCheck("320821197608294103"); err == nil {
		println("320821197608294103 : check pass")
	}
	if err = NumberCheck("320821197608294104"); err == nil {
		println("320821197608294104 : check pass")
	}
	if err = NumberCheck("320821197608294105"); err == nil {
		println("320821197608294105 : check pass")
	}
	if err = NumberCheck("320821197608294106"); err == nil {
		println("320821197608294106 : check pass")
	}
	if err = NumberCheck("320821197608294107"); err == nil {
		println("320821197608294107 : check pass")
	}
	if err = NumberCheck("320821197608294108"); err == nil {
		println("320821197608294108 : check pass")
	}
	if err = NumberCheck("320821197608294109"); err == nil {
		println("320821197608294109 : check pass")
	}
	if err = NumberCheck("32082119760829410x"); err == nil {
		println("32082119760829410x : check pass")
	}
}
