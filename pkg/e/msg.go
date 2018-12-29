package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "Request parameter error",
	ERROR_EXIST_TAG:                 "The tag name already exists,",
	ERROR_EXIST_TAG_FAIL:            "Failed to get existing tags",
	ERROR_NOT_EXIST_TAG:             "The label does not exist",
	ERROR_GET_TAGS_FAIL:             "Failed to get all tags",
	ERROR_COUNT_TAG_FAIL:            "Statistics tag failed",
	ERROR_ADD_TAG_FAIL:              "New label failed",
	ERROR_EDIT_TAG_FAIL:             "Failed to modify the label",
	ERROR_DELETE_TAG_FAIL:           "Failed to delete the label",
	ERROR_EXPORT_TAG_FAIL:           "Export label failed",
	ERROR_IMPORT_TAG_FAIL:           "Import label failed",
	ERROR_NOT_EXIST_ARTICLE:         "The article does not exist",
	ERROR_ADD_ARTICLE_FAIL:          "Added article failed",
	ERROR_DELETE_ARTICLE_FAIL:       "Deleting an article failed",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "Check if the article has failed",
	ERROR_EDIT_ARTICLE_FAIL:         "Modify article failed",
	ERROR_COUNT_ARTICLE_FAIL:        "Statistics article failed",
	ERROR_GET_ARTICLES_FAIL:         "Failed to get multiple articles",
	ERROR_GET_ARTICLE_FAIL:          "Failed to get a single article",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "Failed to generate article posters",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token has timed out",
	ERROR_AUTH_TOKEN:                "Token generation failed",
	ERROR_AUTH:                      "Token error",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "Failed to save the picture",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "Checking the picture failed",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "Check the picture error, the image format or size has a problem",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
