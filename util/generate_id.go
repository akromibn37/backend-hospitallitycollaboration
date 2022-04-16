package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/akromibn37/hospitalityCollaboration/constant"
)

func GenerateID(typ string) string {
	preFix := ""
	midFix := ""
	postFix := ""

	switch typ {
	case constant.ASSIGNMENT_ID:
		preFix = "ASS"
	case constant.CLASS_ID:
		preFix = "CLS"
	case constant.EXERCISE_ID:
		preFix = "EXC"
	case constant.EXERCISE_QUESTION_ID:
		preFix = "EXD"
	case constant.HOMEWORK_ID:
		preFix = "HMW"
	case constant.POINT_ID:
		preFix = "PID"
	case constant.POINT_DETAIL_ID:
		preFix = "PDI"
	case constant.SUBTOPIC_ID:
		preFix = "STI"
	case constant.TEST_ID:
		preFix = "TSI"
	case constant.TEST_QUESTION_ID:
		preFix = "TQI"
	case constant.TOPIC_ID:
		preFix = "TPI"
	case constant.SHIP_ID:
		preFix = "SHP"
	case constant.TRANS_ID:
		preFix = "TRN"
	case constant.DOC_ID:
		preFix = "DOC"
	case constant.CAN_ID:
		preFix = "CAN"
	case constant.SERVICE_ID:
		preFix = "SVC"
	case constant.HOSPITALITY_ID:
		preFix = "HOS"
	case constant.CUSTOMER_PROFILE_ID:
		preFix = "CUS"
	case constant.CUSTOMER_SERVICE_ID:
		preFix = "CSI"
	default:
		preFix = "USI"
	}

	midFix = String(5)
	postFix = String(2)

	return fmt.Sprintf("%s%s%s", preFix, midFix, postFix)
}

const charset = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
