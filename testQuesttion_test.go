package testQuestion

import (
	"reflect"
	"testing"
)

func TestStringMustAddStar_WhenNotHavePair(t *testing.T) {

	give := "abcde"
	expect := []string{"ab", "cd", "e*"}
	get := Question1(give)

	if !reflect.DeepEqual(expect, get) {
		t.Errorf("expect %q but get %q\n", expect, get)

	}

}
