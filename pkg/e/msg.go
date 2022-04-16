package e

var MsgFlags = map[int]string{
	SUCCESS:                      "ok",
	ERROR:                        "fail",
	INVALID_PARAMS:               "Invalid request parameters",
	GET_REQUEST_FAIL:             "Get Request Parameters Failed",
	JSON_UNMARSHAL_FAIL:          "JSON unmarshal failed",
	ERROR_CREATE_USER:            "Error Insert new user to database",
	ERROR_GET_USER:               "Error Get user from database",
	ERROR_UPDATE_USER:            "Error Update user to database",
	ERROR_ADD_CLASS:              "Error Add Class to database",
	ERROR_GET_CLASS:              "Error Get Class from database",
	ERROR_ADD_CLASS_MEMBER:       "Error Add Class Member to Database",
	ERROR_GET_CLASS_MEMBER:       "Error Get Class Member from Database",
	ERROR_UPDATE_CLASS:           "Error Update Class to database",
	ERROR_UPDATE_CLASS_MEMBER:    "Error Update Class Member to database",
	ERROR_GET_SUBJECT_AND_TOPIC:  "Error Get Subject And Topic",
	ERROR_GET_SUBTOPIC:           "Error Get SubTopic",
	ERROR_GET_EXERCISE:           "Error Get Exercise",
	ERROR_GET_EXERCISE_DATA:      "Error Get Exercise Data",
	ERROR_GET_TEST:               "Error Get Test",
	ERROR_GET_TEST_DATA:          "Error Get Test Data",
	ERROR_UPDATE_SUBTOPIC:        "Error Update SubTopic to database",
	ERROR_UPDATE_EXERCISE:        "Error Update Exercise to database",
	ERROR_UPDATE_TEST:            "Error Update Test to database",
	ERROR_UPDATE_TOPIC:           "Error Update Topic to database",
	ERROR_ADD_HOMEWORK:           "Error Add Homework to database",
	ERROR_GET_HOMEWORK:           "Error Get Homework from database",
	ERROR_GET_HOMEWORK_DETAIL:    "Error Get Homework Detail from database",
	ERROR_UPDATE_HOMEWORK:        "Error Update Homework to database",
	ERROR_UPDATE_HOMEWORK_DETAIL: "Error Update Homework Detail to database",
	ERROR_ADD_ASSIGNMENT:         "Error Add Assignment to Database",
	ERROR_GET_ASSIGNMENT:         "Error Get Assignment from database",
	ERROR_GET_POINT_MASTER:       "Error Get Point Master from database",
	ERROR_GET_CLASS_BY_UID:       "Error Get Class By UserId from database",

	ERROR_EXIST_TAG:                 "Tag name already exist",
	ERROR_EXIST_TAG_FAIL:            "failed to get existing tag",
	ERROR_NOT_EXIST_TAG:             "Tag name does not exist",
	ERROR_GET_TAGS_FAIL:             "Get tags error",
	ERROR_COUNT_TAG_FAIL:            "统计标签失败",
	ERROR_ADD_TAG_FAIL:              "新增标签失败",
	ERROR_EDIT_TAG_FAIL:             "修改标签失败",
	ERROR_DELETE_TAG_FAIL:           "删除标签失败",
	ERROR_EXPORT_TAG_FAIL:           "导出标签失败",
	ERROR_IMPORT_TAG_FAIL:           "导入标签失败",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:          "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
	ERROR_EDIT_ARTICLE_FAIL:         "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:        "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:         "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:          "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海报失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
