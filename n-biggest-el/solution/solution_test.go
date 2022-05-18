package solution

import "testing"

func TestFindBiggestEl(t *testing.T) {
	testArr1 := []int{6, 2}
	testArr2 := []int{65, 25, -7, 51, 2, 8}

	want := 51
	if output := FindBiggestEl(testArr1, testArr2); output != want {
		t.Errorf("FindBiggestEl() = %q, want %q", output, want)
	}

    testArr3 := []int{6, 7}
	testArr4 := []int{65, 0, -7, 0, 2, 8}

    want2 := -1
    if output := FindBiggestEl(testArr3, testArr4); output != want2 {
		t.Errorf("FindBiggestEl() = %q, want %q", output, want2)
	}


}

