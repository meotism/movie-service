package errmsg

type ErrMsg struct {
	AuthMsg     *AuthMsg
	FileMsg     *FileMsg
	DatabaseMsg *DatabaseMsg
	QueryMsg    *QueryMsg
}

type AuthMsg struct {
	InvalidEmail        string
	InvalidPassword     string
	AgeNotTrue          string
	ExistedEmail        string
	WrongMailFormat     string
	NotEnoughInfo       string
	WrongFullName       string
	EmptyValue          string
	NewPasswordNotMatch string
	ChangePassFail      string
	ChangePassOK        string
	IsHaveBadWord       string
}

type FileMsg struct {
	FileOver5MB  string
	FileNotOpen  string
	ReadBuffFail string
	NotImageFile string
}

type DatabaseMsg struct {
	SelectFail    string
	UpdateFail    string
	InsertFail    string
	DeleteFail    string
	DeleteSuccess string
}

type QueryMsg struct {
	PageNotNumber      string
	IsAdultWrong       string
	MinRatingWrong     string
	MustBeNumber       string
	ResourceNotFound   string
	StatusWrong        string
	WrongFomat         string
	IsNotCommentAuthor string
}

func InitErrMsg() *ErrMsg {
	return &ErrMsg{
		AuthMsg:     NewAuthMsg(),
		FileMsg:     NewFileMsg(),
		DatabaseMsg: NewDatabaseMsg(),
		QueryMsg:    NewQueryMsg(),
	}
}

func NewAuthMsg() *AuthMsg {
	return &AuthMsg{
		InvalidEmail:        "Invalid mail.",
		InvalidPassword:     "Invalid password.",
		AgeNotTrue:          "Age not true.",
		ExistedEmail:        "Email is existed.",
		WrongMailFormat:     "Wrong email Format.",
		NotEnoughInfo:       "Not enough infomation.",
		WrongFullName:       "Full Name is wrong format.",
		EmptyValue:          "Empty Value: ",
		NewPasswordNotMatch: "New Password not match.",
		ChangePassFail:      "Change password fail.",
		ChangePassOK:        "Password change success.",
		IsHaveBadWord:       " is contain bad word.",
	}
}

func NewFileMsg() *FileMsg {
	return &FileMsg{
		FileOver5MB:  "File over 5MB.",
		FileNotOpen:  "File open fail.",
		ReadBuffFail: "Read Buffer Fail.",
		NotImageFile: "File is not an image.",
	}
}

func NewDatabaseMsg() *DatabaseMsg {
	return &DatabaseMsg{
		SelectFail:    "Select Fails.",
		UpdateFail:    "Update Fail.",
		InsertFail:    "Insert Fail.",
		DeleteFail:    "Delete Fail.",
		DeleteSuccess: "Delete Success.",
	}
}

func NewQueryMsg() *QueryMsg {
	return &QueryMsg{
		PageNotNumber:    "Page param must a number.",
		IsAdultWrong:     "is_adult value must be 0 - non adult/ 1 - adult.",
		MinRatingWrong:   "min_rating value must be greater than 0 and lower than 10.",
		MustBeNumber:     " must be a number.",
		ResourceNotFound: "The resource you requested could not be found.",
		StatusWrong:      "Status value must be 0-Comming soon / 1-Release / 2-prohibit.",
		WrongFomat:       " is wrong format.",
		IsNotCommentAuthor: "Comment author is not current user.",
	}
}
