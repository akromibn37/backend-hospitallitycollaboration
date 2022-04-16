package util

import (
	"fmt"
	"strings"
	"testing"

	"github.com/akromibn37/hospitalityCollaboration/constant"
)

func TestGenearteCase1(t *testing.T) {
	x := GenerateID(constant.ASSIGNMENT_ID)
	fmt.Println(x)
	if !strings.Contains(x, "ASS") {
		t.Errorf("length must be 10")
	}
}
func TestGenearteCase2(t *testing.T) {
	x := GenerateID(constant.CLASS_ID)
	fmt.Println(x)
	if !strings.Contains(x, "CLS") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase3(t *testing.T) {
	x := GenerateID(constant.EXERCISE_ID)
	fmt.Println(x)
	if !strings.Contains(x, "EXC") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase4(t *testing.T) {
	x := GenerateID(constant.EXERCISE_QUESTION_ID)
	fmt.Println(x)
	if !strings.Contains(x, "EXD") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase5(t *testing.T) {
	x := GenerateID(constant.HOMEWORK_ID)
	fmt.Println(x)
	if !strings.Contains(x, "HMW") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase6(t *testing.T) {
	x := GenerateID(constant.POINT_ID)
	fmt.Println(x)
	if !strings.Contains(x, "PID") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase7(t *testing.T) {
	x := GenerateID(constant.POINT_DETAIL_ID)
	fmt.Println(x)
	if !strings.Contains(x, "PDI") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase8(t *testing.T) {
	x := GenerateID(constant.SUBTOPIC_ID)
	fmt.Println(x)
	if !strings.Contains(x, "STI") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase9(t *testing.T) {
	x := GenerateID(constant.TEST_ID)
	fmt.Println(x)
	if !strings.Contains(x, "TSI") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase10(t *testing.T) {
	x := GenerateID(constant.TEST_QUESTION_ID)
	fmt.Println(x)
	if !strings.Contains(x, "TQI") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase11(t *testing.T) {
	x := GenerateID(constant.TOPIC_ID)
	fmt.Println(x)
	if !strings.Contains(x, "TPI") {
		t.Errorf("length must be 10")
	}
}

func TestGenearteCase12(t *testing.T) {
	x := GenerateID(constant.USER_ID)
	fmt.Println(x)
	if !strings.Contains(x, "USI") {
		t.Errorf("length must be 10")
	}
}
