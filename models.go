package valence

// ---- Common ----------------------------------------------------------------

// RichText holds a value in both plain-text and HTML forms.
type RichText struct {
	Text string `json:"Text"`
	Html string `json:"Html"`
}

// PagingInfo is embedded in paged result sets.
type PagingInfo struct {
	Bookmark    string `json:"Bookmark"`
	HasMoreItems bool  `json:"HasMoreItems"`
}

// PagedResultSet is the generic wrapper returned by list endpoints.
type PagedResultSet[T any] struct {
	PagingInfo PagingInfo `json:"PagingInfo"`
	Items      []T        `json:"Items"`
}

// ---- Versions --------------------------------------------------------------

type ProductVersions struct {
	ProductCode string `json:"ProductCode"`
	LatestVersion string `json:"LatestVersion"`
	SupportedVersions []string `json:"SupportedVersions"`
}

// ---- Organization ----------------------------------------------------------

type OrganizationInfo struct {
	Identifier int64  `json:"Identifier"`
	Name       string `json:"Name"`
	TimeZone   string `json:"TimeZone"`
}

// ---- OrgUnit ---------------------------------------------------------------

type OrgUnitTypeInfo struct {
	Id   int64  `json:"Id"`
	Code string `json:"Code"`
	Name string `json:"Name"`
}

type OrgUnit struct {
	Identifier int64           `json:"Identifier"`
	Name       string          `json:"Name"`
	Code       string          `json:"Code"`
	Type       OrgUnitTypeInfo `json:"Type"`
	IsActive   bool            `json:"IsActive"`
}

type OrgUnitProperties struct {
	Identifier   int64           `json:"Identifier"`
	Name         string          `json:"Name"`
	Code         string          `json:"Code"`
	Type         OrgUnitTypeInfo `json:"Type"`
	Parents      []int64         `json:"Parents"`
	IsActive     bool            `json:"IsActive"`
}

// ---- User ------------------------------------------------------------------

type WhoAmIUser struct {
	Identifier       string `json:"Identifier"`
	FirstName        string `json:"FirstName"`
	LastName         string `json:"LastName"`
	UniqueName       string `json:"UniqueName"`
	ProfileIdentifier string `json:"ProfileIdentifier"`
}

type UserActivationData struct {
	IsActive bool `json:"IsActive"`
}

type UserData struct {
	OrgId            int64              `json:"OrgId"`
	UserId           int64              `json:"UserId"`
	FirstName        string             `json:"FirstName"`
	MiddleName       string             `json:"MiddleName"`
	LastName         string             `json:"LastName"`
	UserName         string             `json:"UserName"`
	ExternalEmail    string             `json:"ExternalEmail"`
	OrgDefinedId     string             `json:"OrgDefinedId"`
	UniqueIdentifier string             `json:"UniqueIdentifier"`
	Activation       UserActivationData `json:"Activation"`
}

// ---- Role ------------------------------------------------------------------

type RoleInfo struct {
	Id   int64  `json:"Id"`
	Code string `json:"Code"`
	Name string `json:"Name"`
}

// ---- Enrollment ------------------------------------------------------------

type OrgUnitInfo struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
	Code string `json:"Code"`
	Type OrgUnitTypeInfo `json:"Type"`
	IsActive bool `json:"IsActive"`
}

type MyOrgUnitInfo struct {
	OrgUnit    OrgUnitInfo `json:"OrgUnit"`
	Access     AccessInfo  `json:"Access"`
	IsPinned   bool        `json:"IsPinned"`
}

type AccessInfo struct {
	IsActive      bool   `json:"IsActive"`
	StartDate     string `json:"StartDate"`
	EndDate       string `json:"EndDate"`
	CanAccess     bool   `json:"CanAccess"`
}

type OrgUnitUser struct {
	Identifier int64    `json:"Identifier"`
	DisplayName string  `json:"DisplayName"`
	UserName   string   `json:"UserName"`
	OrgDefinedId string `json:"OrgDefinedId"`
	Role       RoleInfo `json:"Role"`
}

type UserEnrollmentData struct {
	OrgUnit OrgUnitInfo `json:"OrgUnit"`
	Role    RoleInfo    `json:"Role"`
}

type EnrollmentData struct {
	OrgUnitId int64    `json:"OrgUnitId"`
	UserId    int64    `json:"UserId"`
	Role      RoleInfo `json:"Role"`
}

// ---- Course ----------------------------------------------------------------

type CourseOffering struct {
	Identifier       int64    `json:"Identifier"`
	Name             string   `json:"Name"`
	Code             string   `json:"Code"`
	IsActive         bool     `json:"IsActive"`
	CourseTemplate   OrgUnit  `json:"CourseTemplate"`
	Semester         *OrgUnit `json:"Semester"`
	StartDate        *string  `json:"StartDate"`
	EndDate          *string  `json:"EndDate"`
	LocaleId         *int64   `json:"LocaleId"`
	ForceLocale      bool     `json:"ForceLocale"`
	ShowAddressBook  bool     `json:"ShowAddressBook"`
}

type CourseTemplate struct {
	Identifier  int64  `json:"Identifier"`
	Name        string `json:"Name"`
	Code        string `json:"Code"`
	IsActive    bool   `json:"IsActive"`
	Path        string `json:"Path"`
	HomeUrl     string `json:"HomeUrl"`
}

// ---- Group -----------------------------------------------------------------

type GroupEnrollment struct {
	UserId int64 `json:"UserId"`
}

type Group struct {
	GroupId         int64   `json:"GroupId"`
	Name            string  `json:"Name"`
	Code            string  `json:"Code"`
	Description     RichText `json:"Description"`
	Enrollments     []int64 `json:"Enrollments"`
}

type GroupCategory struct {
	GroupCategoryId         int64    `json:"GroupCategoryId"`
	Name                    string   `json:"Name"`
	Description             RichText `json:"Description"`
	EnrollmentStyle         int      `json:"EnrollmentStyle"`
	EnrollmentQuantity      *int     `json:"EnrollmentQuantity"`
	AutoEnroll              bool     `json:"AutoEnroll"`
	RandomizeEnrollments    bool     `json:"RandomizeEnrollments"`
	NumberOfGroups          *int     `json:"NumberOfGroups"`
	MaxUsersPerGroup        *int     `json:"MaxUsersPerGroup"`
	AllocateAfterExpiry     bool     `json:"AllocateAfterExpiry"`
	SelfEnrollmentExpiryDate *string `json:"SelfEnrollmentExpiryDate"`
	Groups                  []Group  `json:"Groups"`
	RestrictedByOrgUnitId   *int64   `json:"RestrictedByOrgUnitId"`
}

// ---- Section ---------------------------------------------------------------

type Section struct {
	SectionId   int64    `json:"SectionId"`
	Name        string   `json:"Name"`
	Code        string   `json:"Code"`
	Description RichText `json:"Description"`
	Enrollments []int64  `json:"Enrollments"`
}

type SectionPropertyData struct {
	EnrollmentStyle    int  `json:"EnrollmentStyle"`
	EnrollmentQuantity *int `json:"EnrollmentQuantity"`
	AutoEnroll         bool `json:"AutoEnroll"`
	RandomizeEnrollments bool `json:"RandomizeEnrollments"`
}

// ---- Grade -----------------------------------------------------------------

type GradeSchemeRange struct {
	Symbol    string  `json:"Symbol"`
	Low       float64 `json:"Low"`
	High      float64 `json:"High"`
	AssignedValue float64 `json:"AssignedValue"`
}

type GradeSchemeInfo struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

type GradeObject struct {
	MaxPoints              *float64        `json:"MaxPoints"`
	CanExceedMaxPoints     bool            `json:"CanExceedMaxPoints"`
	IsBonus                bool            `json:"IsBonus"`
	ExcludeFromFinalGrade  bool            `json:"ExcludeFromFinalGrade"`
	GradeScheme            *GradeSchemeInfo `json:"GradeScheme"`
	Id                     int64           `json:"Id"`
	Name                   string          `json:"Name"`
	ShortName              string          `json:"ShortName"`
	GradeType              string          `json:"GradeType"`
	CategoryId             *int64          `json:"CategoryId"`
	Description            RichText        `json:"Description"`
	GradeObjectTypeId      int             `json:"GradeObjectTypeId"`
	ActivityId             *string         `json:"ActivityId"`
	AssociatedTool         *AssociatedTool `json:"AssociatedTool"`
}

type AssociatedTool struct {
	ToolId         int64  `json:"ToolId"`
	ToolItemId     int64  `json:"ToolItemId"`
}

type GradeCategory struct {
	Id                    int64    `json:"Id"`
	Name                  string   `json:"Name"`
	ShortName             string   `json:"ShortName"`
	CanExceedMax          bool     `json:"CanExceedMax"`
	ExcludeFromFinalGrade bool     `json:"ExcludeFromFinalGrade"`
	StartDate             *string  `json:"StartDate"`
	EndDate               *string  `json:"EndDate"`
	MaxPoints             *float64 `json:"MaxPoints"`
	Weight                *float64 `json:"Weight"`
	BoostUserScore        bool     `json:"BoostUserScore"`
	AutoPoints            bool     `json:"AutoPoints"`
	WeightDistributionType int     `json:"WeightDistributionType"`
	NumberOfHighestToDrop *int     `json:"NumberOfHighestToDrop"`
	NumberOfLowestToDrop  *int     `json:"NumberOfLowestToDrop"`
}

type GradeValue struct {
	UserId             int64    `json:"UserId"`
	OrgUnitId          int64    `json:"OrgUnitId"`
	DisplayedGrade     string   `json:"DisplayedGrade"`
	GradeObjectIdentifier int64 `json:"GradeObjectIdentifier"`
	GradeObjectName    string   `json:"GradeObjectName"`
	GradeObjectType    int      `json:"GradeObjectType"`
	GradeObjectTypeName string  `json:"GradeObjectTypeName"`
	PointsNumerator    *float64 `json:"PointsNumerator"`
	PointsDenominator  *float64 `json:"PointsDenominator"`
	WeightedDenominator *float64 `json:"WeightedDenominator"`
	WeightedNumerator   *float64 `json:"WeightedNumerator"`
	Comments           *RichText `json:"Comments"`
	PrivateComments    *RichText `json:"PrivateComments"`
	LastModified       *string  `json:"LastModified"`
	LastModifiedBy     *UserData `json:"LastModifiedBy"`
	Released           bool     `json:"Released"`
}

type FinalGradeValue struct {
	UserId         int64   `json:"UserId"`
	OrgUnitId      int64   `json:"OrgUnitId"`
	DisplayedGrade string  `json:"DisplayedGrade"`
	GradeObjectIdentifier int64 `json:"GradeObjectIdentifier"`
	GradeObjectName string `json:"GradeObjectName"`
	PointsNumerator *float64 `json:"PointsNumerator"`
	PointsDenominator *float64 `json:"PointsDenominator"`
	WeightedDenominator *float64 `json:"WeightedDenominator"`
	WeightedNumerator *float64 `json:"WeightedNumerator"`
	Comments        *RichText `json:"Comments"`
	PrivateComments *RichText `json:"PrivateComments"`
	Released        bool     `json:"Released"`
}

// ---- Class List ------------------------------------------------------------

type ClasslistUser struct {
	Identifier    int64  `json:"Identifier"`
	ProfileIdentifier string `json:"ProfileIdentifier"`
	DisplayName   string `json:"DisplayName"`
	UserName      string `json:"UserName"`
	OrgDefinedId  string `json:"OrgDefinedId"`
	Email         string `json:"Email"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
}

// ---- News ------------------------------------------------------------------

type NewsItem struct {
	Id               int64    `json:"Id"`
	Title            string   `json:"Title"`
	Body             RichText `json:"Body"`
	StartDate        *string  `json:"StartDate"`
	EndDate          *string  `json:"EndDate"`
	IsGlobal         bool     `json:"IsGlobal"`
	IsPublished      bool     `json:"IsPublished"`
	ShowOnlyInCourseOfferings bool `json:"ShowOnlyInCourseOfferings"`
}

// ---- Quiz ------------------------------------------------------------------

type QuizReadData struct {
	QuizId           int64    `json:"QuizId"`
	Name             string   `json:"Name"`
	IsActive         bool     `json:"IsActive"`
	SortOrder        int      `json:"SortOrder"`
	SubmissionTimeLimit TimeLimit `json:"SubmissionTimeLimit"`
	StartDate        *string  `json:"StartDate"`
	EndDate          *string  `json:"EndDate"`
	DueDate          *string  `json:"DueDate"`
	DisplayInCalendar bool    `json:"DisplayInCalendar"`
	Instructions     Instructions `json:"Instructions"`
	Description      Description  `json:"Description"`
}

type TimeLimit struct {
	IsEnforced     bool `json:"IsEnforced"`
	ShowClock      bool `json:"ShowClock"`
	TimeLimitValue int  `json:"TimeLimitValue"`
}

type Instructions struct {
	IsDisplayed bool     `json:"IsDisplayed"`
	Text        RichText `json:"Text"`
}

type Description struct {
	IsDisplayed bool     `json:"IsDisplayed"`
	Text        RichText `json:"Text"`
}

type QuizAttemptData struct {
	AttemptId         int64   `json:"AttemptId"`
	UserId            int64   `json:"UserId"`
	AttemptNumber     int     `json:"AttemptNumber"`
	TimeStarted       string  `json:"TimeStarted"`
	TimeCompleted     *string `json:"TimeCompleted"`
	Score             *float64 `json:"Score"`
	IsInProgress      bool    `json:"IsInProgress"`
}

type QuizQuestion struct {
	QuestionId   int64    `json:"QuestionId"`
	Name         string   `json:"Name"`
	QuestionText RichText `json:"QuestionText"`
	Points       float64  `json:"Points"`
	Difficulty   int      `json:"Difficulty"`
	Bonus        bool     `json:"Bonus"`
	Mandatory    bool     `json:"Mandatory"`
	QuestionTypeId int    `json:"QuestionTypeId"`
}

type QuizSpecialAccessData struct {
	UserId          int64   `json:"UserId"`
	StartDate       *string `json:"StartDate"`
	EndDate         *string `json:"EndDate"`
	DueDate         *string `json:"DueDate"`
	IsActive        bool    `json:"IsActive"`
}

// ---- Discussion ------------------------------------------------------------

type Forum struct {
	ForumId          int64    `json:"ForumId"`
	Name             string   `json:"Name"`
	Description      RichText `json:"Description"`
	AllowAnonymous   bool     `json:"AllowAnonymous"`
	IsLocked         bool     `json:"IsLocked"`
	IsHidden         bool     `json:"IsHidden"`
	RequiresApproval bool     `json:"RequiresApproval"`
	StartDate        *string  `json:"StartDate"`
	EndDate          *string  `json:"EndDate"`
}

type Topic struct {
	ForumId          int64    `json:"ForumId"`
	TopicId          int64    `json:"TopicId"`
	Name             string   `json:"Name"`
	Description      RichText `json:"Description"`
	AllowAnonymous   bool     `json:"AllowAnonymous"`
	IsLocked         bool     `json:"IsLocked"`
	IsHidden         bool     `json:"IsHidden"`
	RequiresApproval bool     `json:"RequiresApproval"`
	StartDate        *string  `json:"StartDate"`
	EndDate          *string  `json:"EndDate"`
	TopicType        int      `json:"TopicType"`
}

type Post struct {
	PostId      int64    `json:"PostId"`
	TopicId     int64    `json:"TopicId"`
	ParentPostId *int64  `json:"ParentPostId"`
	Subject     string   `json:"Subject"`
	Message     RichText `json:"Message"`
	IsAnonymous bool     `json:"IsAnonymous"`
	IsApproved  bool     `json:"IsApproved"`
	IsDeleted   bool     `json:"IsDeleted"`
	ThreadId    int64    `json:"ThreadId"`
	UserId      int64    `json:"UserId"`
	DatePosted  string   `json:"DatePosted"`
	LastModified *string `json:"LastModified"`
}

// ---- Dropbox ---------------------------------------------------------------

type DropboxFolder struct {
	Id                 int64    `json:"Id"`
	CategoryId         *int64   `json:"CategoryId"`
	Name               string   `json:"Name"`
	CustomInstructions RichText `json:"CustomInstructions"`
	TotalFiles         int      `json:"TotalFiles"`
	UnreadFiles        int      `json:"UnreadFiles"`
	FlaggedFiles       int      `json:"FlaggedFiles"`
	TotalUsers         int      `json:"TotalUsers"`
	IsHidden           bool     `json:"IsHidden"`
	DueDate            *string  `json:"DueDate"`
	DisplayInCalendar  bool     `json:"DisplayInCalendar"`
}

type DropboxCategory struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

type SubmissionFile struct {
	FileId    int64  `json:"FileId"`
	FileName  string `json:"FileName"`
	Size      int64  `json:"Size"`
}

type Submission struct {
	Id              int64           `json:"Id"`
	SubmittedBy     UserData        `json:"SubmittedBy"`
	Files           []SubmissionFile `json:"Files"`
	SubmissionDate  string          `json:"SubmissionDate"`
	Comment         string          `json:"Comment"`
}

// ---- Survey ----------------------------------------------------------------

type Survey struct {
	SurveyId          int64    `json:"SurveyId"`
	Name              string   `json:"Name"`
	IsActive          bool     `json:"IsActive"`
	IsAnonymous       bool     `json:"IsAnonymous"`
	StartDate         *string  `json:"StartDate"`
	EndDate           *string  `json:"EndDate"`
	DueDate           *string  `json:"DueDate"`
	DisplayInCalendar bool     `json:"DisplayInCalendar"`
	Instructions      Instructions `json:"Instructions"`
}

type SurveyAttempt struct {
	AttemptId     int64   `json:"AttemptId"`
	UserId        int64   `json:"UserId"`
	AttemptNumber int     `json:"AttemptNumber"`
	TimeStarted   string  `json:"TimeStarted"`
	TimeCompleted *string `json:"TimeCompleted"`
	IsInProgress  bool    `json:"IsInProgress"`
}

// ---- Self Assessment -------------------------------------------------------

type SelfAssessment struct {
	SelfAssessmentId  int64    `json:"SelfAssessmentId"`
	Name              string   `json:"Name"`
	IsActive          bool     `json:"IsActive"`
	StartDate         *string  `json:"StartDate"`
	EndDate           *string  `json:"EndDate"`
	Instructions      Instructions `json:"Instructions"`
}

type SelfAssessmentAttempt struct {
	AttemptId     int64   `json:"AttemptId"`
	UserId        int64   `json:"UserId"`
	AttemptNumber int     `json:"AttemptNumber"`
	TimeStarted   string  `json:"TimeStarted"`
	TimeCompleted *string `json:"TimeCompleted"`
}

// ---- LTI -------------------------------------------------------------------

type LTILink struct {
	LtiLinkId   int64  `json:"LtiLinkId"`
	Title       string `json:"Title"`
	Url         string `json:"Url"`
	Description string `json:"Description"`
	IsHidden    bool   `json:"IsHidden"`
}

type LTIAdvantageLink struct {
	LinkId      int64  `json:"LinkId"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Url         string `json:"Url"`
}

type LTISharingData struct {
	OrgUnitId int64  `json:"OrgUnitId"`
	Name      string `json:"Name"`
}

type LTIDeploymentSharingData struct {
	OrgUnitId int64  `json:"OrgUnitId"`
	Name      string `json:"Name"`
	IsShared  bool   `json:"IsShared"`
}

// ---- Tools -----------------------------------------------------------------

type ToolInfo struct {
	ToolId    int64  `json:"ToolId"`
	ToolName  string `json:"ToolName"`
	IsEnabled bool   `json:"IsEnabled"`
}

// ---- Config Variables ------------------------------------------------------

type ConfigVariableValue struct {
	VariableUUID string `json:"VariableUUID"`
	OrgUnitId    int64  `json:"OrgUnitId"`
	Value        string `json:"Value"`
}

// ---- Course Import ---------------------------------------------------------

type CourseImportJobData struct {
	JobToken string `json:"JobToken"`
}

// ---- Badges ----------------------------------------------------------------

type IssuedBadge struct {
	IssuedId    int64  `json:"IssuedId"`
	UserId      int64  `json:"UserId"`
	BadgeId     int64  `json:"BadgeId"`
	BadgeName   string `json:"BadgeName"`
	IssueDate   string `json:"IssueDate"`
	Expiry      *string `json:"Expiry"`
	Evidence    string `json:"Evidence"`
	Narrative   string `json:"Narrative"`
}
