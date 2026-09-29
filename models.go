package valence

import (
	"encoding/json"
	"strconv"
)

// ---- Common ----------------------------------------------------------------

// RichText holds a value in both plain-text and HTML forms.
type RichText struct {
	Text string  `json:"Text"`
	Html *string `json:"Html"`
}

// RichTextInput is the input form used by Valence write operations.
type RichTextInput struct {
	Content string `json:"Content"`
	Type    string `json:"Type"`
}

// PagingInfo is embedded in paged result sets.
type PagingInfo struct {
	Bookmark     string `json:"Bookmark"`
	HasMoreItems bool   `json:"HasMoreItems"`
}

// PagedResultSet is the generic wrapper returned by list endpoints that use an "Items" key.
type PagedResultSet[T any] struct {
	PagingInfo PagingInfo `json:"PagingInfo"`
	Items      []T        `json:"Items"`
}

// NumberOrString accepts Valence fields that are documented as numbers but may
// be returned by some tenants as symbolic strings.
type NumberOrString string

func (v *NumberOrString) UnmarshalJSON(data []byte) error {
	var text string
	if err := json.Unmarshal(data, &text); err == nil {
		*v = NumberOrString(text)
		return nil
	}
	var number int
	if err := json.Unmarshal(data, &number); err == nil {
		*v = NumberOrString(strconv.Itoa(number))
		return nil
	}
	return nil
}

// ObjectListPage is the generic paging wrapper for endpoints that return an "Objects" array.
// Next is non-nil when more pages exist and contains the URL to fetch the next page.
type ObjectListPage[T any] struct {
	Objects []T     `json:"Objects"`
	Next    *string `json:"Next"`
}

// ---- Versions --------------------------------------------------------------

type ProductVersions struct {
	ProductCode       string   `json:"ProductCode"`
	LatestVersion     string   `json:"LatestVersion"`
	SupportedVersions []string `json:"SupportedVersions"`
}

// ---- Organization ----------------------------------------------------------

type OrganizationInfo struct {
	Identifier int64  `json:"Identifier,string"`
	Name       string `json:"Name"`
	TimeZone   string `json:"TimeZone"`
}

// ---- OrgUnit ---------------------------------------------------------------

type OrgUnitTypeInfo struct {
	Id   int64  `json:"Id"`
	Code string `json:"Code"`
	Name string `json:"Name"`
}

type OrgUnitType struct {
	OrgUnitTypeInfo
	Description string                 `json:"Description"`
	SortOrder   int                    `json:"SortOrder"`
	Permissions OrgUnitTypePermissions `json:"Permissions"`
}

type OrgUnitTypePermissions struct {
	CanDelete bool `json:"CanDelete"`
	CanEdit   bool `json:"CanEdit"`
}

type OrgUnit struct {
	Identifier int64           `json:"Identifier,string"`
	Name       string          `json:"Name"`
	Code       *string         `json:"Code"`
	Path       string          `json:"Path"`
	Type       OrgUnitTypeInfo `json:"Type"`
}

type OrgUnitProperties struct {
	Identifier int64           `json:"Identifier,string"`
	Name       string          `json:"Name"`
	Code       *string         `json:"Code"`
	Path       string          `json:"Path"`
	Type       OrgUnitTypeInfo `json:"Type"`
}

// ---- User ------------------------------------------------------------------

type WhoAmIUser struct {
	Identifier        string `json:"Identifier"`
	FirstName         string `json:"FirstName"`
	LastName          string `json:"LastName"`
	UniqueName        string `json:"UniqueName"`
	ProfileIdentifier string `json:"ProfileIdentifier"`
	Pronouns          string `json:"Pronouns"`
}

type UserActivationData struct {
	IsActive bool `json:"IsActive"`
}

type UserData struct {
	OrgId            int64              `json:"OrgId"`
	UserId           int64              `json:"UserId"`
	FirstName        string             `json:"FirstName"`
	MiddleName       *string            `json:"MiddleName"`
	LastName         string             `json:"LastName"`
	UserName         string             `json:"UserName"`
	ExternalEmail    *string            `json:"ExternalEmail"`
	OrgDefinedId     *string            `json:"OrgDefinedId"`
	UniqueIdentifier string             `json:"UniqueIdentifier"`
	Activation       UserActivationData `json:"Activation"`
	DisplayName      string             `json:"DisplayName"`
	LastAccessedDate *string            `json:"LastAccessedDate"`
	FirstLoginDate   *string            `json:"FirstLoginDate"`
	Pronouns         string             `json:"Pronouns"`
}

type CreateUserData struct {
	OrgDefinedId      *string           `json:"OrgDefinedId"`
	FirstName         string            `json:"FirstName"`
	MiddleName        *string           `json:"MiddleName"`
	LastName          string            `json:"LastName"`
	ExternalEmail     *string           `json:"ExternalEmail"`
	UserName          string            `json:"UserName"`
	RoleId            int64             `json:"RoleId"`
	IsActive          bool              `json:"IsActive"`
	SendCreationEmail bool              `json:"SendCreationEmail"`
	Pronouns          *string           `json:"Pronouns"`
	PasswordData      *UserPasswordData `json:"PasswordData"`
}

type UserPasswordData struct {
	Password           string `json:"Password"`
	ForcePasswordReset bool   `json:"ForcePasswordReset"`
}

// ---- Role ------------------------------------------------------------------

type RoleInfo struct {
	Id   int64   `json:"Id"`
	Code *string `json:"Code"`
	Name string  `json:"Name"`
}

// Role describes a user role returned by the LP roles service. The fields after
// Code are only populated when the role is retrieved through LP's unstable
// contract.
type Role struct {
	Identifier            string `json:"Identifier"`
	DisplayName           string `json:"DisplayName"`
	Code                  string `json:"Code"`
	Description           string `json:"Description"`
	RoleAlias             string `json:"RoleAlias"`
	IsCascading           bool   `json:"IsCascading"`
	AccessFutureCourses   bool   `json:"AccessFutureCourses"`
	AccessInactiveCourses bool   `json:"AccessInactiveCourses"`
	AccessPastCourses     bool   `json:"AccessPastCourses"`
	ShowInGrades          bool   `json:"ShowInGrades"`
	ShowInUserProgress    bool   `json:"ShowInUserProgress"`
	InClassList           bool   `json:"InClassList"`
}

// ---- Enrollment ------------------------------------------------------------

type OrgUnitInfo struct {
	Id       int64           `json:"Id"`
	Name     string          `json:"Name"`
	Code     *string         `json:"Code"`
	Type     OrgUnitTypeInfo `json:"Type"`
	HomeUrl  *string         `json:"HomeUrl"`
	ImageUrl *string         `json:"ImageUrl"`
}

type MyOrgUnitInfo struct {
	OrgUnit OrgUnitInfo `json:"OrgUnit"`
	Access  AccessInfo  `json:"Access"`
	PinDate *string     `json:"PinDate"`
}

type AccessInfo struct {
	IsActive          bool     `json:"IsActive"`
	StartDate         *string  `json:"StartDate"`
	EndDate           *string  `json:"EndDate"`
	CanAccess         bool     `json:"CanAccess"`
	ClasslistRoleName *string  `json:"ClasslistRoleName"`
	LISRoles          []string `json:"LISRoles"`
	LastAccessed      *string  `json:"LastAccessed"`
}

type OrgUnitUser struct {
	User OrgUnitUserInfo `json:"User"`
	Role RoleInfo        `json:"Role"`
}

type OrgUnitUserInfo struct {
	Identifier        *string `json:"Identifier"`
	DisplayName       *string `json:"DisplayName"`
	EmailAddress      *string `json:"EmailAddress"`
	OrgDefinedId      *string `json:"OrgDefinedId"`
	ProfileBadgeUrl   *string `json:"ProfileBadgeUrl"`
	ProfileIdentifier *string `json:"ProfileIdentifier"`
	UserName          *string `json:"UserName"`
}

type UserEnrollmentData struct {
	OrgUnit            OrgUnitInfo `json:"OrgUnit"`
	Role               RoleInfo    `json:"Role"`
	IsCascading        bool        `json:"IsCascading"`
	EnrolledByUserId   *int64      `json:"EnrolledByUserId"`
	EnrolledByUserDate *string     `json:"EnrolledByUserDate"`
}

type EnrollmentData struct {
	OrgUnitId   int64 `json:"OrgUnitId"`
	UserId      int64 `json:"UserId"`
	RoleId      int64 `json:"RoleId"`
	IsCascading bool  `json:"IsCascading"`
}

type CreateEnrollmentData struct {
	OrgUnitId           int64 `json:"OrgUnitId"`
	UserId              int64 `json:"UserId"`
	RoleId              int64 `json:"RoleId"`
	SendEnrollmentEmail *bool `json:"SendEnrollmentEmail"`
}

// ---- Course ----------------------------------------------------------------

type CourseOffering struct {
	Identifier      int64         `json:"Identifier,string"`
	Name            string        `json:"Name"`
	Code            string        `json:"Code"`
	IsActive        bool          `json:"IsActive"`
	Path            string        `json:"Path"`
	CourseTemplate  *BasicOrgUnit `json:"CourseTemplate"`
	Semester        *BasicOrgUnit `json:"Semester"`
	Department      *BasicOrgUnit `json:"Department"`
	StartDate       *string       `json:"StartDate"`
	EndDate         *string       `json:"EndDate"`
	LocaleId        *int64        `json:"LocaleId"`
	ForceLocale     bool          `json:"ForceLocale"`
	ShowAddressBook bool          `json:"ShowAddressBook"`
	Description     RichText      `json:"Description"`
	CanSelfRegister bool          `json:"CanSelfRegister"`
}

// BasicOrgUnit is the org-unit reference embedded in a course offering.
type BasicOrgUnit struct {
	Identifier int64  `json:"Identifier,string"`
	Name       string `json:"Name"`
	Code       string `json:"Code"`
}

// CourseOfferingInfo is the complete replacement block for updating course offering information.
type CourseOfferingInfo struct {
	Name            string        `json:"Name"`
	Code            string        `json:"Code"`
	StartDate       *string       `json:"StartDate"`
	EndDate         *string       `json:"EndDate"`
	LocaleId        *int64        `json:"LocaleId"`
	ForceLocale     bool          `json:"ForceLocale"`
	ShowAddressBook bool          `json:"ShowAddressBook"`
	IsActive        bool          `json:"IsActive"`
	Description     RichTextInput `json:"Description"`
	CanSelfRegister *bool         `json:"CanSelfRegister"`
}

// CreateCourseOffering is the input block for creating a course offering.
type CreateCourseOffering struct {
	Name             string        `json:"Name"`
	Code             string        `json:"Code"`
	Path             string        `json:"Path"`
	CourseTemplateId int64         `json:"CourseTemplateId"`
	SemesterId       *int64        `json:"SemesterId"`
	StartDate        *string       `json:"StartDate"`
	EndDate          *string       `json:"EndDate"`
	LocaleId         *int64        `json:"LocaleId"`
	ForceLocale      bool          `json:"ForceLocale"`
	ShowAddressBook  bool          `json:"ShowAddressBook"`
	Description      RichTextInput `json:"Description"`
	CanSelfRegister  *bool         `json:"CanSelfRegister"`
	IsActive         bool          `json:"IsActive"`
}

type CourseTemplate struct {
	Identifier int64  `json:"Identifier,string"`
	Name       string `json:"Name"`
	Code       string `json:"Code"`
	Path       string `json:"Path"`
}

type CourseTemplateInfo struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}

type CreateCourseTemplate struct {
	Name             string  `json:"Name"`
	Code             string  `json:"Code"`
	Path             string  `json:"Path"`
	ParentOrgUnitIds []int64 `json:"ParentOrgUnitIds"`
}

type FileSystemObjectType int

const (
	FileSystemObjectTypeFolder FileSystemObjectType = 1
	FileSystemObjectTypeFile   FileSystemObjectType = 2
)

type FileSystemObject struct {
	Name                 string               `json:"Name"`
	FileSystemObjectType FileSystemObjectType `json:"FileSystemObjectType"`
}

// ---- Group -----------------------------------------------------------------

type GroupEnrollment struct {
	UserId int64 `json:"UserId"`
}

type GroupData struct {
	Name        string        `json:"Name"`
	Code        string        `json:"Code"`
	Description RichTextInput `json:"Description"`
}

type Group struct {
	GroupId     int64    `json:"GroupId"`
	Name        string   `json:"Name"`
	Code        string   `json:"Code"`
	Description RichText `json:"Description"`
	Enrollments []int64  `json:"Enrollments"`
}

type GroupCategory struct {
	GroupCategoryId               int64          `json:"GroupCategoryId"`
	Name                          string         `json:"Name"`
	Description                   RichText       `json:"Description"`
	EnrollmentStyle               NumberOrString `json:"EnrollmentStyle"`
	EnrollmentQuantity            *int           `json:"EnrollmentQuantity"`
	AutoEnroll                    bool           `json:"AutoEnroll"`
	RandomizeEnrollments          bool           `json:"RandomizeEnrollments"`
	MaxUsersPerGroup              *int           `json:"MaxUsersPerGroup"`
	AllocateAfterExpiry           bool           `json:"AllocateAfterExpiry"`
	SelfEnrollmentExpiryDate      *string        `json:"SelfEnrollmentExpiryDate"`
	SelfEnrollmentStartDate       *string        `json:"SelfEnrollmentStartDate"`
	GroupPrefix                   *string        `json:"GroupPrefix"`
	Groups                        []int64        `json:"Groups"`
	RestrictedByOrgUnitId         *int64         `json:"RestrictedByOrgUnitId"`
	DescriptionsVisibleToEnrolees bool           `json:"DescriptionsVisibleToEnrolees"`
}

type GroupCategoryData struct {
	Name                          string        `json:"Name"`
	Description                   RichTextInput `json:"Description"`
	EnrollmentStyle               int           `json:"EnrollmentStyle"`
	EnrollmentQuantity            *int          `json:"EnrollmentQuantity"`
	AutoEnroll                    bool          `json:"AutoEnroll"`
	RandomizeEnrollments          bool          `json:"RandomizeEnrollments"`
	NumberOfGroups                *int          `json:"NumberOfGroups"`
	MaxUsersPerGroup              *int          `json:"MaxUsersPerGroup"`
	AllocateAfterExpiry           bool          `json:"AllocateAfterExpiry"`
	SelfEnrollmentStartDate       *string       `json:"SelfEnrollmentStartDate"`
	SelfEnrollmentExpiryDate      *string       `json:"SelfEnrollmentExpiryDate"`
	GroupPrefix                   *string       `json:"GroupPrefix"`
	RestrictedByOrgUnitId         *int64        `json:"RestrictedByOrgUnitId"`
	DescriptionsVisibleToEnrolees bool          `json:"DescriptionsVisibleToEnrolees"`
}

type GroupsJobData struct {
	OrgUnitId  int64  `json:"OrgUnitId"`
	CategoryId int64  `json:"CategoryId"`
	SubmitDate string `json:"SubmitDate"`
	Status     int    `json:"Status"`
}

type GroupCategoryJobStatus struct {
	Status int `json:"Status"`
}

// ---- Section ---------------------------------------------------------------

type Section struct {
	SectionId   int64    `json:"SectionId"`
	Name        string   `json:"Name"`
	Code        string   `json:"Code"`
	Description RichText `json:"Description"`
	Enrollments []int64  `json:"Enrollments"`
}

type SectionData struct {
	Name        string        `json:"Name"`
	Code        string        `json:"Code"`
	Description RichTextInput `json:"Description"`
}

type CreateSectionSettingsData struct {
	EnrollmentStyle                int64 `json:"EnrollmentStyle"`
	EnrollmentQuantity             int64 `json:"EnrollmentQuantity"`
	AutoEnroll                     bool  `json:"AutoEnroll"`
	RandomizeEnrollments           bool  `json:"RandomizeEnrollments"`
	DescriptionsVisibleToEnrollees bool  `json:"DescriptionsVisibleToEnrollees"`
}

type UpdateSectionSettingsData struct {
	Name                           string        `json:"Name"`
	Description                    RichTextInput `json:"Description"`
	AutoEnroll                     bool          `json:"AutoEnroll"`
	RandomizeEnrollments           bool          `json:"RandomizeEnrollments"`
	DescriptionsVisibleToEnrollees bool          `json:"DescriptionsVisibleToEnrollees"`
}

type SectionSettingsData struct {
	Name                           string        `json:"Name"`
	Description                    RichTextInput `json:"Description"`
	EnrollmentStyle                int64         `json:"EnrollmentStyle"`
	EnrollmentQuantity             int64         `json:"EnrollmentQuantity"`
	AutoEnroll                     bool          `json:"AutoEnroll"`
	RandomizeEnrollments           bool          `json:"RandomizeEnrollments"`
	DescriptionsVisibleToEnrollees bool          `json:"DescriptionsVisibleToEnrollees"`
}

// ---- Grade -----------------------------------------------------------------

type GradeSchemeRange struct {
	PercentStart  float64  `json:"PercentStart"`
	Symbol        string   `json:"Symbol"`
	AssignedValue *float64 `json:"AssignedValue"`
	Colour        string   `json:"Colour"`
}

type GradeScheme struct {
	Id        int64              `json:"Id"`
	Name      string             `json:"Name"`
	ShortName string             `json:"ShortName"`
	Ranges    []GradeSchemeEntry `json:"Ranges"`
}

type GradeSchemeEntry = GradeSchemeRange

type GradeSetupInfo struct {
	GradingSystem        string `json:"GradingSystem"`
	IsNullGradeZero      bool   `json:"IsNullGradeZero"`
	DefaultGradeSchemeId int64  `json:"DefaultGradeSchemeId"`
}

type GradeObject struct {
	MaxPoints                        *float64        `json:"MaxPoints"`
	CanExceedMaxPoints               bool            `json:"CanExceedMaxPoints"`
	IsBonus                          bool            `json:"IsBonus"`
	ExcludeFromFinalGradeCalculation bool            `json:"ExcludeFromFinalGradeCalculation"`
	GradeSchemeId                    *int64          `json:"GradeSchemeId"`
	GradeSchemeUrl                   string          `json:"GradeSchemeUrl"`
	Id                               int64           `json:"Id"`
	Name                             string          `json:"Name"`
	ShortName                        string          `json:"ShortName"`
	GradeType                        string          `json:"GradeType"`
	CategoryId                       *int64          `json:"CategoryId"`
	Description                      RichText        `json:"Description"`
	AssociatedTool                   *AssociatedTool `json:"AssociatedTool"`
	IsHidden                         bool            `json:"IsHidden"`
	Weight                           *float64        `json:"Weight"`
}

type AssociatedTool struct {
	ToolId     int64 `json:"ToolId"`
	ToolItemId int64 `json:"ToolItemId"`
}

type GradeCategory struct {
	Id                              int64          `json:"Id"`
	Grades                          []GradeObject  `json:"Grades"`
	Name                            string         `json:"Name"`
	ShortName                       string         `json:"ShortName"`
	CanExceedMax                    bool           `json:"CanExceedMax"`
	ExcludeFromFinalGrade           bool           `json:"ExcludeFromFinalGrade"`
	StartDate                       *string        `json:"StartDate"`
	EndDate                         *string        `json:"EndDate"`
	MaxPoints                       *float64       `json:"MaxPoints"`
	Weight                          *float64       `json:"Weight"`
	AutoPoints                      *bool          `json:"AutoPoints"`
	WeightDistributionType          *int           `json:"WeightDistributionType"`
	NumberOfHighestToDrop           *int           `json:"NumberOfHighestToDrop"`
	NumberOfLowestToDrop            *int           `json:"NumberOfLowestToDrop"`
	Description                     *RichTextInput `json:"Description"`
	ShowDescription                 *bool          `json:"ShowDescription"`
	DisplayClassAverageToUsers      *bool          `json:"DisplayClassAverageToUsers"`
	DisplayGradeDistributionToUsers *bool          `json:"DisplayGradeDistributionToUsers"`
	OverrideDisplayOptions          *bool          `json:"OverrideDisplayOptions"`
	DisplayPointsToUsers            *bool          `json:"DisplayPointsToUsers"`
	DisplayWeightToUsers            *bool          `json:"DisplayWeightToUsers"`
	DisplayGradeSchemeSymbolToUsers *bool          `json:"DisplayGradeSchemeSymbolToUsers"`
	DisplayGradeSchemeColorToUsers  *bool          `json:"DisplayGradeSchemeColorToUsers"`
}

type GradeValue struct {
	UserId                string    `json:"UserId"`
	OrgUnitId             string    `json:"OrgUnitId"`
	DisplayedGrade        string    `json:"DisplayedGrade"`
	GradeObjectIdentifier string    `json:"GradeObjectIdentifier"`
	GradeObjectName       string    `json:"GradeObjectName"`
	GradeObjectType       int       `json:"GradeObjectType"`
	GradeObjectTypeName   *string   `json:"GradeObjectTypeName"`
	PointsNumerator       *float64  `json:"PointsNumerator"`
	PointsDenominator     *float64  `json:"PointsDenominator"`
	WeightedDenominator   *float64  `json:"WeightedDenominator"`
	WeightedNumerator     *float64  `json:"WeightedNumerator"`
	Comments              *RichText `json:"Comments"`
	PrivateComments       *RichText `json:"PrivateComments"`
	LastModified          *string   `json:"LastModified"`
	LastModifiedBy        *string   `json:"LastModifiedBy"`
	ReleasedDate          *string   `json:"ReleasedDate"`
}

// GradeUserRef is the user reference embedded in a GradeValueEntry.
type GradeUserRef = OrgUnitUserInfo

// GradeValueData is the grade value data embedded in a GradeValueEntry.
type GradeValueData = GradeValue

// GradeValueEntry is one item in the ObjectListPage returned by the grade values endpoint.
// GradeValue is nil for users who have not been graded.
type GradeValueEntry struct {
	User       GradeUserRef    `json:"User"`
	GradeValue *GradeValueData `json:"GradeValue"`
}

type FinalGradeValue = GradeValue

// FinalGradeValueEntry is one item returned by the paginated final grade values list endpoint.
// GradeValue is nil when no final grade has been assigned.
type FinalGradeValueEntry struct {
	User       GradeUserRef    `json:"User"`
	GradeValue *GradeValueData `json:"GradeValue"`
}

// ---- Class List ------------------------------------------------------------

type ClasslistUser struct {
	Identifier               int64   `json:"Identifier,string"`
	ProfileIdentifier        string  `json:"ProfileIdentifier"`
	DisplayName              string  `json:"DisplayName"`
	UserName                 *string `json:"Username"`
	OrgDefinedId             *string `json:"OrgDefinedId"`
	Email                    *string `json:"Email"`
	FirstName                *string `json:"FirstName"`
	LastName                 *string `json:"LastName"`
	RoleId                   *int64  `json:"RoleId"`
	ClasslistRoleDisplayName string  `json:"ClasslistRoleDisplayName"`
	LastAccessed             *string `json:"LastAccessed"`
	IsOnline                 bool    `json:"IsOnline"`
	Pronouns                 *string `json:"Pronouns"`
}

// ---- News ------------------------------------------------------------------

type NewsItem struct {
	Id                        int64      `json:"Id"`
	Title                     string     `json:"Title"`
	Body                      RichText   `json:"Body"`
	StartDate                 *string    `json:"StartDate"`
	EndDate                   *string    `json:"EndDate"`
	IsGlobal                  bool       `json:"IsGlobal"`
	IsPublished               bool       `json:"IsPublished"`
	ShowOnlyInCourseOfferings bool       `json:"ShowOnlyInCourseOfferings"`
	IsHidden                  bool       `json:"IsHidden"`
	Attachments               []NewsFile `json:"Attachments"`
	CreatedBy                 *int64     `json:"CreatedBy"`
	CreatedDate               *string    `json:"CreatedDate"`
	LastModifiedBy            *int64     `json:"LastModifiedBy"`
	LastModifiedDate          *string    `json:"LastModifiedDate"`
	IsAuthorInfoShown         bool       `json:"IsAuthorInfoShown"`
	IsPinned                  bool       `json:"IsPinned"`
	PinnedDate                *string    `json:"PinnedDate"`
	IsStartDateShown          bool       `json:"IsStartDateShown"`
	SortOrder                 int        `json:"SortOrder"`
}

type NewsFile struct {
	FileId   int64  `json:"FileId"`
	FileName string `json:"FileName"`
	FileSize int64  `json:"FileSize"`
}

// ---- Quiz ------------------------------------------------------------------

type QuizReadData struct {
	QuizId                          int64                  `json:"QuizId"`
	Name                            string                 `json:"Name"`
	IsActive                        bool                   `json:"IsActive"`
	SortOrder                       int                    `json:"SortOrder"`
	AutoExportToGrades              *bool                  `json:"AutoExportToGrades"`
	GradeItemId                     *int64                 `json:"GradeItemId"`
	IsAutoSetGraded                 bool                   `json:"IsAutoSetGraded"`
	SubmissionTimeLimit             TimeLimit              `json:"SubmissionTimeLimit"`
	StartDate                       *string                `json:"StartDate"`
	EndDate                         *string                `json:"EndDate"`
	DueDate                         *string                `json:"DueDate"`
	DisplayInCalendar               bool                   `json:"DisplayInCalendar"`
	Instructions                    Instructions           `json:"Instructions"`
	Description                     Description            `json:"Description"`
	AttemptsAllowed                 QuizAttemptsAllowed    `json:"AttemptsAllowed"`
	LateSubmissionInfo              QuizLateSubmissionInfo `json:"LateSubmissionInfo"`
	SubmissionGracePeriod           *int                   `json:"SubmissionGracePeriod"`
	Password                        *string                `json:"Password"`
	Header                          Instructions           `json:"Header"`
	Footer                          Instructions           `json:"Footer"`
	AllowHints                      bool                   `json:"AllowHints"`
	DisableRightClick               bool                   `json:"DisableRightClick"`
	DisablePagerAndAlerts           bool                   `json:"DisablePagerAndAlerts"`
	NotificationEmail               *string                `json:"NotificationEmail"`
	CalcTypeId                      int                    `json:"CalcTypeId"`
	CategoryId                      *int64                 `json:"CategoryId"`
	PreventMovingBackwards          bool                   `json:"PreventMovingBackwards"`
	Shuffle                         bool                   `json:"Shuffle"`
	ActivityId                      *string                `json:"ActivityId"`
	AllowOnlyUsersWithSpecialAccess bool                   `json:"AllowOnlyUsersWithSpecialAccess"`
	IsRetakeIncorrectOnly           bool                   `json:"IsRetakeIncorrectOnly"`
	PagingTypeId                    *int                   `json:"PagingTypeId"`
	IsSynchronous                   bool                   `json:"IsSynchronous"`
	DeductionPercentage             *float64               `json:"DeductionPercentage"`
	HideQuestionPoints              bool                   `json:"HideQuestionPoints"`
	IsSingleSession                 bool                   `json:"IsSingleSession"`
	AnnotationToolsEnabled          bool                   `json:"AnnotationToolsEnabled"`
	RestrictIPAddressRange          []QuizIPRange          `json:"RestrictIPAddressRange"`
}

type QuizAttemptsAllowed struct {
	IsUnlimited             bool `json:"IsUnlimited"`
	NumberOfAttemptsAllowed *int `json:"NumberOfAttemptsAllowed"`
}

type QuizLateSubmissionInfo struct {
	LateSubmissionOption int  `json:"LateSubmissionOption"`
	LateLimitMinutes     *int `json:"LateLimitMinutes"`
}

type QuizIPRange struct {
	IPRangeStart string  `json:"IPRangeStart"`
	IPRangeEnd   *string `json:"IPRangeEnd"`
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
	AttemptId                   int64    `json:"AttemptId"`
	QuizId                      int64    `json:"QuizId"`
	UserId                      int64    `json:"UserId"`
	AttemptNumber               int      `json:"AttemptNumber"`
	Started                     string   `json:"Started"`
	Completed                   *string  `json:"Completed"`
	Score                       *float64 `json:"Score"`
	AttemptFeedback             RichText `json:"AttemptFeedback"`
	FeedbackLastModified        *string  `json:"FeedbackLastModified"`
	IsPublished                 bool     `json:"IsPublished"`
	IsRetakeIncorrectOnly       bool     `json:"IsRetakeIncorrectOnly"`
	AttemptDueDate              *string  `json:"AttemptDueDate"`
	AttemptEnforceTimeLimit     bool     `json:"AttemptEnforceTimeLimit"`
	AttemptSubmissionTimeLimit  int      `json:"AttemptSubmissionTimeLimit"`
	AttemptSubmissionGraceLimit int      `json:"AttemptSubmissionGraceLimit"`
	AttemptSubmissionLateTypeId int      `json:"AttemptSubmissionLateTypeId"`
	AttemptSubmissionLateData   int      `json:"AttemptSubmissionLateData"`
	AttemptIsSynchronous        bool     `json:"AttemptIsSynchronous"`
	DeductionPercentage         *float64 `json:"DeductionPercentage"`
}

type QuizQuestion struct {
	QuestionId                int64           `json:"QuestionId"`
	Name                      *string         `json:"Name"`
	QuestionText              RichText        `json:"QuestionText"`
	Points                    float64         `json:"Points"`
	Difficulty                int             `json:"Difficulty"`
	Bonus                     bool            `json:"Bonus"`
	Mandatory                 bool            `json:"Mandatory"`
	QuestionTypeId            int             `json:"QuestionTypeId"`
	Hint                      RichText        `json:"Hint"`
	Feedback                  RichText        `json:"Feedback"`
	LastModified              *string         `json:"LastModified"`
	LastModifiedBy            *int64          `json:"LastModifiedBy"`
	SectionId                 int64           `json:"SectionId"`
	QuestionTemplateId        int64           `json:"QuestionTemplateId"`
	QuestionTemplateVersionId int64           `json:"QuestionTemplateVersionId"`
	QuestionInfo              json.RawMessage `json:"QuestionInfo"`
}

type QuizSpecialAccessData struct {
	StartDate           *string                     `json:"StartDate"`
	EndDate             *string                     `json:"EndDate"`
	DueDate             *string                     `json:"DueDate"`
	SubmissionTimeLimit *QuizSpecialAccessTimeLimit `json:"SubmissionTimeLimit"`
	AttemptsAllowed     *QuizAttemptsAllowed        `json:"AttemptsAllowed"`
}

type QuizSpecialAccessTimeLimit struct {
	IsEnforced     bool `json:"IsEnforced"`
	TimeLimitValue int  `json:"TimeLimitValue"`
}

// ---- Discussion ------------------------------------------------------------

type Forum struct {
	ForumId                    int64    `json:"ForumId"`
	Name                       string   `json:"Name"`
	Description                RichText `json:"Description"`
	AllowAnonymous             bool     `json:"AllowAnonymous"`
	IsLocked                   bool     `json:"IsLocked"`
	IsHidden                   bool     `json:"IsHidden"`
	RequiresApproval           bool     `json:"RequiresApproval"`
	StartDate                  *string  `json:"StartDate"`
	EndDate                    *string  `json:"EndDate"`
	PostStartDate              *string  `json:"PostStartDate"`
	PostEndDate                *string  `json:"PostEndDate"`
	ShowDescriptionInTopics    *bool    `json:"ShowDescriptionInTopics"`
	DisplayInCalendar          bool     `json:"DisplayInCalendar"`
	DisplayPostDatesInCalendar bool     `json:"DisplayPostDatesInCalendar"`
	StartDateAvailabilityType  *string  `json:"StartDateAvailabilityType"`
	EndDateAvailabilityType    *string  `json:"EndDateAvailabilityType"`
}

type Topic struct {
	ForumId                   int64    `json:"ForumId"`
	TopicId                   int64    `json:"TopicId"`
	Name                      string   `json:"Name"`
	Description               RichText `json:"Description"`
	AllowAnonymousPosts       bool     `json:"AllowAnonymousPosts"`
	IsLocked                  bool     `json:"IsLocked"`
	IsHidden                  bool     `json:"IsHidden"`
	RequiresApproval          bool     `json:"RequiresApproval"`
	StartDate                 *string  `json:"StartDate"`
	EndDate                   *string  `json:"EndDate"`
	UnlockStartDate           *string  `json:"UnlockStartDate"`
	UnlockEndDate             *string  `json:"UnlockEndDate"`
	UnApprovedPostCount       int      `json:"UnApprovedPostCount"`
	PinnedPostCount           int      `json:"PinnedPostCount"`
	ScoringType               string   `json:"ScoringType"`
	IsAutoScore               bool     `json:"IsAutoScore"`
	ScoreOutOf                *float64 `json:"ScoreOutOf"`
	IncludeNonScoredValues    bool     `json:"IncludeNonScoredValues"`
	ScoredCount               int      `json:"ScoredCount"`
	RatingsSum                float64  `json:"RatingsSum"`
	RatingsCount              int      `json:"RatingsCount"`
	MustPostToParticipate     bool     `json:"MustPostToParticipate"`
	RatingType                string   `json:"RatingType"`
	ActivityId                *string  `json:"ActivityId"`
	GroupTypeId               *int64   `json:"GroupTypeId"`
	StartDateAvailabilityType *string  `json:"StartDateAvailabilityType"`
	EndDateAvailabilityType   *string  `json:"EndDateAvailabilityType"`
	DueDate                   *string  `json:"DueDate"`
}

// ---- Content ---------------------------------------------------------------

type TableOfContents struct {
	Modules []ContentModule `json:"Modules"`
}

type ContentModule struct {
	ModuleId         int64           `json:"ModuleId"`
	Title            string          `json:"Title"`
	SortOrder        int             `json:"SortOrder"`
	StartDateTime    *string         `json:"StartDateTime"`
	EndDateTime      *string         `json:"EndDateTime"`
	Modules          []ContentModule `json:"Modules"`
	Topics           []ContentTopic  `json:"Topics"`
	IsHidden         bool            `json:"IsHidden"`
	IsLocked         bool            `json:"IsLocked"`
	PacingStartDate  *string         `json:"PacingStartDate"`
	PacingEndDate    *string         `json:"PacingEndDate"`
	DefaultPath      string          `json:"DefaultPath"`
	LastModifiedDate *string         `json:"LastModifiedDate"`
}

type ContentTopic struct {
	TopicId                   int64   `json:"TopicId"`
	Identifier                int64   `json:"Identifier,string"`
	TypeIdentifier            string  `json:"TypeIdentifier"`
	Title                     string  `json:"Title"`
	Bookmarked                bool    `json:"Bookmarked"`
	Unread                    bool    `json:"Unread"`
	Url                       string  `json:"Url"`
	SortOrder                 int     `json:"SortOrder"`
	StartDateTime             *string `json:"StartDateTime"`
	EndDateTime               *string `json:"EndDateTime"`
	ActivityId                *string `json:"ActivityId"`
	CompletionType            int     `json:"CompletionType"`
	IsExempt                  bool    `json:"IsExempt"`
	IsHidden                  bool    `json:"IsHidden"`
	IsLocked                  bool    `json:"IsLocked"`
	IsBroken                  bool    `json:"IsBroken"`
	ToolId                    *int64  `json:"ToolId"`
	ToolItemId                *int64  `json:"ToolItemId"`
	ActivityType              int     `json:"ActivityType"`
	GradeItemId               *int64  `json:"GradeItemId"`
	LastModifiedDate          *string `json:"LastModifiedDate"`
	StartDateAvailabilityType *int    `json:"StartDateAvailabilityType"`
	EndDateAvailabilityType   *int    `json:"EndDateAvailabilityType"`
}

type Post struct {
	PostId                 int64            `json:"PostId"`
	TopicId                int64            `json:"TopicId"`
	ForumId                int64            `json:"ForumId"`
	ParentPostId           *int64           `json:"ParentPostId"`
	Subject                string           `json:"Subject"`
	Message                RichText         `json:"Message"`
	IsAnonymous            bool             `json:"IsAnonymous"`
	RequiresApproval       bool             `json:"RequiresApproval"`
	IsDeleted              bool             `json:"IsDeleted"`
	ThreadId               int64            `json:"ThreadId"`
	UserId                 *int64           `json:"PostingUserId"`
	PostingUserDisplayName string           `json:"PostingUserDisplayName"`
	DatePosted             string           `json:"DatePosted"`
	LastEditedDate         *string          `json:"LastEditedDate"`
	LastEditedBy           *int64           `json:"LastEditedBy"`
	CanRate                bool             `json:"CanRate"`
	ReplyPostIds           []int64          `json:"ReplyPostIds"`
	WordCount              int              `json:"WordCount"`
	AttachmentCount        int              `json:"AttachmentCount"`
	IsRead                 bool             `json:"IsRead"`
	Attachments            []SubmissionFile `json:"Attachments"`
	ThreadIsPinned         bool             `json:"ThreadIsPinned"`
}

// ---- Dropbox ---------------------------------------------------------------

type RubricLevel struct {
	Id     int64    `json:"Id"`
	Name   string   `json:"Name"`
	Points *float64 `json:"Points"`
}

type RubricCell struct {
	Description RichText `json:"Description"`
	Feedback    RichText `json:"Feedback"`
	Points      *float64 `json:"Points"`
}

type RubricCriterion struct {
	Id    int64        `json:"Id"`
	Name  string       `json:"Name"`
	Cells []RubricCell `json:"Cells"`
}

type RubricCriteriaGroup struct {
	Name     string            `json:"Name"`
	Levels   []RubricLevel     `json:"Levels"`
	Criteria []RubricCriterion `json:"Criteria"`
}

type RubricOverallLevel struct {
	Id          int64    `json:"Id"`
	Name        string   `json:"Name"`
	RangeStart  *float64 `json:"RangeStart"`
	Description RichText `json:"Description"`
	Feedback    RichText `json:"Feedback"`
}

type DropboxAssessment struct {
	ScoreDenominator *float64 `json:"ScoreDenominator"`
	Rubrics          []Rubric `json:"Rubrics"`
}

type DropboxFolder struct {
	Id                              int64                `json:"Id"`
	CategoryId                      *int64               `json:"CategoryId"`
	Name                            string               `json:"Name"`
	CustomInstructions              RichText             `json:"CustomInstructions"`
	Attachments                     []SubmissionFile     `json:"Attachments"`
	TotalFiles                      int                  `json:"TotalFiles"`
	UnreadFiles                     int                  `json:"UnreadFiles"`
	FlaggedFiles                    int                  `json:"FlaggedFiles"`
	TotalUsers                      int                  `json:"TotalUsers"`
	TotalUsersWithSubmissions       int                  `json:"TotalUsersWithSubmissions"`
	TotalUsersWithFeedback          int                  `json:"TotalUsersWithFeedback"`
	IsHidden                        bool                 `json:"IsHidden"`
	IsAnonymous                     bool                 `json:"IsAnonymous"`
	DueDate                         *string              `json:"DueDate"`
	DisplayInCalendar               bool                 `json:"DisplayInCalendar"`
	DropboxType                     string               `json:"DropboxType"`
	SubmissionType                  string               `json:"SubmissionType"`
	CompletionType                  string               `json:"CompletionType"`
	GroupTypeId                     *int64               `json:"GroupTypeId"`
	GradeItemId                     *int64               `json:"GradeItemId"`
	ActivityId                      *string              `json:"ActivityId"`
	Assessment                      DropboxAssessment    `json:"Assessment"`
	Availability                    *DropboxAvailability `json:"Availability"`
	NotificationEmail               *string              `json:"NotificationEmail"`
	LinkAttachments                 []DropboxLink        `json:"LinkAttachments"`
	SubmissionRule                  string               `json:"SubmissionRule"`
	AllowOnlyUsersWithSpecialAccess *bool                `json:"AllowOnlyUsersWithSpecialAccess"`
}

type DropboxAvailability struct {
	StartDate                 *string `json:"StartDate"`
	EndDate                   *string `json:"EndDate"`
	StartDateAvailabilityType *string `json:"StartDateAvailabilityType"`
	EndDateAvailabilityType   *string `json:"EndDateAvailabilityType"`
}

type DropboxCategory struct {
	Id                   int64   `json:"Id"`
	Name                 string  `json:"Name"`
	LastModifiedByUserId *int64  `json:"LastModifiedByUserId"`
	LastModifiedDate     *string `json:"LastModifiedDate"`
}

type SubmissionFile struct {
	FileId   int64  `json:"FileId"`
	FileName string `json:"FileName"`
	Size     int64  `json:"Size"`
}

// DropboxEntity identifies the user or group at the top level of a submission group.
type DropboxEntity struct {
	DisplayName string `json:"DisplayName"`
	Name        string `json:"Name"`
	EntityId    int64  `json:"EntityId"`
	EntityType  string `json:"EntityType"`
}

// DropboxFeedback holds instructor feedback for a user's folder submission.
type DropboxFeedback struct {
	Score             *float64          `json:"Score"`
	Feedback          *RichText         `json:"Feedback"`
	IsGraded          bool              `json:"IsGraded"`
	GradedSymbol      *string           `json:"GradedSymbol"`
	Files             []SubmissionFile  `json:"Files"`
	Links             []DropboxLink     `json:"Links"`
	RubricAssessments []json.RawMessage `json:"RubricAssessments"`
}

type DropboxLink struct {
	Type     string  `json:"Type"`
	LinkId   int64   `json:"LinkId"`
	LinkName string  `json:"LinkName"`
	Href     *string `json:"Href"`
}

// DropboxSubmitter is the per-submission submitter reference (Id is a string user ID).
type DropboxSubmitter struct {
	Id          string `json:"Id"`
	DisplayName string `json:"DisplayName"`
}

// DropboxSubmissionFile is a file attached to a dropbox submission entry.
type DropboxSubmissionFile struct {
	FileId    int64  `json:"FileId"`
	FileName  string `json:"FileName"`
	Size      int64  `json:"Size"`
	IsRead    bool   `json:"isRead"`
	IsFlagged bool   `json:"isFlagged"`
}

// DropboxSubmissionEntry is a single submission within a UserSubmissions group.
type DropboxSubmissionEntry struct {
	Id             int64                   `json:"Id"`
	SubmittedBy    DropboxSubmitter        `json:"SubmittedBy"`
	SubmissionDate *string                 `json:"SubmissionDate"`
	Comment        RichText                `json:"Comment"`
	Files          []DropboxSubmissionFile `json:"Files"`
}

// UserSubmissions groups all submissions (and feedback) for one user in a dropbox folder.
// This is the element type returned by GetDropboxSubmissions.
type UserSubmissions struct {
	Entity         DropboxEntity            `json:"Entity"`
	Status         string                   `json:"Status"`
	Feedback       DropboxFeedback          `json:"Feedback"`
	Submissions    []DropboxSubmissionEntry `json:"Submissions"`
	CompletionDate *string                  `json:"CompletionDate"`
}

// ---- Survey ----------------------------------------------------------------

type Survey struct {
	SurveyId                        int64               `json:"SurveyId"`
	Name                            string              `json:"Name"`
	IsActive                        bool                `json:"IsActive"`
	IsAnonymous                     bool                `json:"IsAnonymous"`
	StartDate                       *string             `json:"StartDate"`
	EndDate                         *string             `json:"EndDate"`
	DisplayInCalendar               bool                `json:"DisplayInCalendar"`
	SortOrder                       int                 `json:"SortOrder"`
	HasInstantFeedback              bool                `json:"HasInstantFeedback"`
	Description                     Description         `json:"Description"`
	Submission                      RichText            `json:"Submission"`
	Footer                          Description         `json:"Footer"`
	UserResponses                   SurveyUserResponses `json:"UserResponses"`
	CategoryId                      *int64              `json:"CategoryId"`
	PreventMovingBackwards          bool                `json:"PreventMovingBackwards"`
	Shuffle                         bool                `json:"Shuffle"`
	ActivityId                      *string             `json:"ActivityId"`
	AllowOnlyUsersWithSpecialAccess bool                `json:"AllowOnlyUsersWithSpecialAccess"`
}

type SurveyUserResponses struct {
	AttemptsAllowedTypeId int  `json:"AttemptsAllowedTypeId"`
	NumberOfAttempts      *int `json:"NumberOfAttempts"`
}

type SurveyAttempt struct {
	AttemptId     int64   `json:"AttemptId"`
	SurveyId      int64   `json:"SurveyId"`
	UserId        *int64  `json:"UserId"`
	AttemptNumber int     `json:"AttemptNumber"`
	Started       string  `json:"Started"`
	Completed     *string `json:"Completed"`
}

// ---- LTI -------------------------------------------------------------------

type LTILink struct {
	LtiLinkId                       int64                `json:"LtiLinkId"`
	Title                           string               `json:"Title"`
	Url                             string               `json:"Url"`
	Description                     string               `json:"Description"`
	Key                             string               `json:"Key"`
	IsVisible                       bool                 `json:"IsVisible"`
	SignMessage                     bool                 `json:"SignMessage"`
	SignWithTc                      bool                 `json:"SignWithTc"`
	SendTcInfo                      bool                 `json:"SendTcInfo"`
	SendContextInfo                 bool                 `json:"SendContextInfo"`
	SendUserId                      bool                 `json:"SendUserId"`
	SendUserName                    bool                 `json:"SendUserName"`
	SendUserEmail                   bool                 `json:"SendUserEmail"`
	SendLinkTitle                   bool                 `json:"SendLinkTitle"`
	SendLinkDescription             bool                 `json:"SendLinkDescription"`
	SendD2LUserName                 bool                 `json:"SendD2LUserName"`
	SendD2LOrgDefinedId             bool                 `json:"SendD2LOrgDefinedId"`
	SendD2LOrgRoleId                bool                 `json:"SendD2LOrgRoleId"`
	SendSectionCode                 bool                 `json:"SendSectionCode"`
	UseToolProviderSecuritySettings bool                 `json:"UseToolProviderSecuritySettings"`
	CustomParameters                []LTICustomParameter `json:"CustomParameters"`
}

type LTIAdvantageLink struct {
	LinkId           int64                `json:"LinkId"`
	DeploymentId     string               `json:"DeploymentId"`
	IsEnabled        bool                 `json:"IsEnabled"`
	Name             string               `json:"Name"`
	Description      *string              `json:"Description"`
	URL              string               `json:"URL"`
	Type             int                  `json:"Type"`
	Height           *int                 `json:"Height"`
	Width            *int                 `json:"Width"`
	CustomParameters []LTICustomParameter `json:"CustomParameters"`
	IsAvailable      bool                 `json:"IsAvailable"`
	OwnerOrgUnitId   int64                `json:"OwnerOrgUnitId"`
}

type LTIAdvantageQuicklink struct {
	LtiLinkId int64  `json:"LtiLinkId"`
	PublicUrl string `json:"PublicUrl"`
}

type LTISharingData struct {
	SharingOrgUnitId     int64   `json:"SharingOrgUnitId"`
	ShareWithOrgUnit     bool    `json:"ShareWithOrgUnit"`
	ShareWithDescendants bool    `json:"ShareWithDescendants"`
	DescendantsTypes     []int64 `json:"DescendantsTypes"`
}

type LTICustomParameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type LTIToolProvider struct {
	LtiToolProviderId    int64  `json:"LtiToolProviderId"`
	OrgUnitId            int64  `json:"OrgUnitId"`
	LaunchPoint          string `json:"LaunchPoint"`
	UseDefaultTcInfo     bool   `json:"UseDefaultTcInfo"`
	Name                 string `json:"Name"`
	Description          string `json:"Description"`
	ContactEmail         string `json:"ContactEmail"`
	IsVisible            bool   `json:"IsVisible"`
	Key                  string `json:"Key"`
	SendTcInfo           bool   `json:"SendTcInfo"`
	SendContextInfo      bool   `json:"SendContextInfo"`
	SendUserId           bool   `json:"SendUserId"`
	SendUserName         bool   `json:"SendUserName"`
	SendUserEmail        bool   `json:"SendUserEmail"`
	SendLinkTitle        bool   `json:"SendLinkTitle"`
	SendLinkDescription  bool   `json:"SendLinkDescription"`
	SendD2LUserName      bool   `json:"SendD2LUserName"`
	SendD2LOrgDefinedId  bool   `json:"SendD2LOrgDefinedId"`
	SendD2LOrgRoleId     bool   `json:"SendD2LOrgRoleId"`
	SendSectionCode      bool   `json:"SendSectionCode"`
	OAuthSignatureMethod int    `json:"OAuthSignatureMethod"`
	Version              int    `json:"Version"`
}

type LTIDeploymentSharingData struct {
	SharingOrgUnitId     int64 `json:"SharingOrgUnitId"`
	ShareWithOrgUnit     bool  `json:"ShareWithOrgUnit"`
	ShareWithDescendants bool  `json:"ShareWithDescendants"`
	Inherited            *bool `json:"Inherited"`
}

type LTIAdvantageCreateSharingRuleData struct {
	SharingOrgUnitId     int64 `json:"SharingOrgUnitId"`
	ShareWithOrgUnit     bool  `json:"ShareWithOrgUnit"`
	ShareWithDescendants bool  `json:"ShareWithDescendants"`
}

// ---- Tools -----------------------------------------------------------------

// OrgToolInfo is the organization-level information returned by the tools/org/ route.
type OrgToolInfo struct {
	ToolId      string `json:"ToolId"`
	DisplayName string `json:"DisplayName"`
	OrgId       int64  `json:"OrgId"`
	Status      bool   `json:"Status"`
	OUDefault   bool   `json:"OUDefault"`
}

// OrgUnitInformation block returned by GET /d2l/api/lp/(version)/tools/orgUnits/(orgUnitId)
type ToolInfo struct {
	ToolId           string `json:"ToolId"`
	DisplayName      string `json:"DisplayName"`
	OrgUnitId        int64  `json:"OrgUnitId"`
	Status           bool   `json:"Status"`
	CustomNavbarName string `json:"CustomNavbarName"`
}

// ToolWithName block returned by tool listing routes when namesOnly is true.
type ToolWithName struct {
	ToolId      string `json:"ToolId"`
	DisplayName string `json:"DisplayName"`
}

// ---- Rubrics ---------------------------------------------------------------

type Rubric struct {
	RubricId                      int64                 `json:"RubricId"`
	Name                          string                `json:"Name"`
	Description                   RichText              `json:"Description"`
	RubricType                    int                   `json:"RubricType"`
	RubricStateId                 int                   `json:"RubricStateId"`
	ScoringMethod                 int                   `json:"ScoringMethod"`
	Visibility                    *int                  `json:"Visibility"`
	IsScoreVisibleToAssessedUsers bool                  `json:"IsScoreVisibleToAssessedUsers"`
	ReverseLevelDisplayOrder      bool                  `json:"ReverseLevelDisplayOrder"`
	CriteriaGroups                []RubricCriteriaGroup `json:"CriteriaGroups"`
	OverallLevels                 []RubricOverallLevel  `json:"OverallLevels"`
}

// ---- Release Conditions ----------------------------------------------------

type ReleaseConditionsData struct {
	Expression ExpressionData `json:"Expression"`
}

type ExpressionData struct {
	Type             string           `json:"Type"`
	State            *string          `json:"State"`
	ExpressionParams ExpressionParams `json:"ExpressionParams"`
	Text             RichText         `json:"Text"`
}

type ExpressionParams struct {
	Operator string            `json:"Operator"`
	Operands []json.RawMessage `json:"Operands"`
}

// ---- Intelligent Agents ----------------------------------------------------

type IntelligentAgent struct {
	AgentId     *int64                     `json:"AgentId"`
	Name        string                     `json:"Name"`
	Description string                     `json:"Description"`
	IsEnabled   bool                       `json:"IsEnabled"`
	Schedule    *IntelligentAgentSchedule  `json:"Schedule"`
	Action      *IntelligentAgentAction    `json:"Action"`
	Condition   *IntelligentAgentCondition `json:"Condition"`
	LastRunDate *string                    `json:"LastRunDate"`
	NextRunDate *string                    `json:"NextRunDate"`
	CategoryId  *int64                     `json:"CategoryId"`
}

type IntelligentAgentSchedule struct {
	IsEnabled           bool     `json:"IsEnabled"`
	Type                *int     `json:"Type"`
	StartDate           *string  `json:"StartDate"`
	EndDate             *string  `json:"EndDate"`
	RepeatsEvery        *int     `json:"RepeatsEvery"`
	RepeatsOnDay        *int     `json:"RepeatsOnDay"`
	RepeatsOnDays       []string `json:"RepeatsOnDays"`
	RepeatsOnMonth      *int     `json:"RepeatsOnMonth"`
	ScheduledTimeHour   *int     `json:"ScheduledTimeHour"`
	ScheduledTimeMinute *int     `json:"ScheduledTimeMinute"`
}

type IntelligentAgentAction struct {
	RepeatType       int                           `json:"RepeatType"`
	EmailAction      *IntelligentAgentEmailAction  `json:"EmailAction"`
	EnrollmentAction *IntelligentAgentEnrollAction `json:"EnrollmentAction"`
}

type IntelligentAgentEmailAction struct {
	IsEnabled bool    `json:"IsEnabled"`
	To        *string `json:"To"`
	Cc        *string `json:"Cc"`
	Bcc       *string `json:"Bcc"`
	Subject   *string `json:"Subject"`
	Message   *string `json:"Message"`
	IsHtml    bool    `json:"IsHtml"`
}

type IntelligentAgentEnrollAction struct {
	IsEnabled      bool   `json:"IsEnabled"`
	EnrollmentType *int   `json:"EnrollmentType"`
	OrgUnitId      *int64 `json:"OrgUnitId"`
	RoleId         *int64 `json:"RoleId"`
}

type IntelligentAgentCondition struct {
	LoginActivity    *IntelligentAgentDateCondition    `json:"LoginActivity"`
	CourseActivity   *IntelligentAgentDateCondition    `json:"CourseActivity"`
	ReleaseCondition *IntelligentAgentReleaseCondition `json:"ReleaseCondition"`
	RoleIds          []int64                           `json:"RoleIds"`
}

type IntelligentAgentDateCondition struct {
	Type int `json:"Type"`
	Days int `json:"Days"`
}

type IntelligentAgentReleaseCondition struct {
	ConditionSetId *int64 `json:"ConditionSetId"`
}

// ---- Config Variables ------------------------------------------------------

type ConfigVariableValue struct {
	OrgUnitId int64   `json:"OrgUnitId"`
	Value     *string `json:"Value"`
}

type SpecifiedOrgUnitValue struct {
	OrgUnitValue *string `json:"OrgUnitValue"`
}

type UpdateStatus struct {
	Status bool `json:"Status"`
}

// ---- Course Import / Copy --------------------------------------------------

type CreateCopyJobRequest struct {
	SourceOrgUnitId             int64    `json:"SourceOrgUnitId"`
	Components                  []string `json:"Components"`
	CallbackUrl                 *string  `json:"CallbackUrl,omitempty"`
	DaysToOffsetDates           *int     `json:"DaysToOffsetDates,omitempty"`
	HoursToOffsetDates          *float64 `json:"HoursToOffsetDates,omitempty"`
	OffsetByStartDateDifference *bool    `json:"OffsetByStartDateDifference,omitempty"`
}

type CreateCopyJobResponse struct {
	JobToken string `json:"JobToken"`
}

type GetCopyJobResponse struct {
	Status string `json:"Status"`
}

type CourseImportJobData struct {
	JobToken string `json:"JobToken"`
}

type CourseImportJobStatus struct {
	JobToken        string `json:"JobToken"`
	TargetOrgUnitId int64  `json:"TargetOrgUnitId"`
	Status          string `json:"Status"`
}

// ---- Badges ----------------------------------------------------------------

type IssuedBadge struct {
	IssuedId       int64            `json:"IssuedId"`
	OrgUnitId      int64            `json:"OrgUnitId"`
	Criteria       string           `json:"Criteria"`
	Evidence       string           `json:"Evidence"`
	IssuedDate     string           `json:"IssuedDate"`
	ExpiryDate     *string          `json:"ExpiryDate"`
	IssuedByUserId int64            `json:"IssuedByUserId"`
	IssuedToUserId int64            `json:"IssuedToUserId"`
	Credit         *float64         `json:"Credit"`
	Share          IssuedAwardShare `json:"Share"`
	Award          Award            `json:"Award"`
	CertificateId  *string          `json:"CertificateId"`
}

type IssuedAwardShare struct {
	SentToProfile    bool   `json:"SentToProfile"`
	SentToMozilla    bool   `json:"SentToMozilla"`
	SentToEportfolio bool   `json:"SentToEportfolio"`
	SharedObjectId   *int64 `json:"SharedObjectId"`
	WithProfile      bool   `json:"WithProfile"`
	WithMozilla      bool   `json:"WithMozilla"`
	WithEportfolio   bool   `json:"WithEportfolio"`
}

type Award struct {
	AwardId            int64                   `json:"AwardId"`
	CreatedBy          int64                   `json:"CreatedBy"`
	Title              string                  `json:"Title"`
	Description        string                  `json:"Description"`
	ExpiryCalculation  AwardExpiryCalculation  `json:"ExpiryCalculation"`
	ExpiryNotification AwardExpiryNotification `json:"ExpiryNotification"`
	IssuerName         string                  `json:"IssuerName"`
	IssuerUrl          string                  `json:"IssuerUrl"`
	IssuerContact      string                  `json:"IssuerContact"`
	AwardType          int                     `json:"AwardType"`
	CertificateData    AwardFileData           `json:"CertificateData"`
	ImageData          AwardFileData           `json:"ImageData"`
	IsDeleted          bool                    `json:"IsDeleted"`
	Criteria           *string                 `json:"Criteria"`
}

type AwardExpiryCalculation struct {
	ExpiryCalculationType int  `json:"ExpiryCalculationType"`
	Minute                *int `json:"Minute"`
	Hour                  *int `json:"Hour"`
	Day                   *int `json:"Day"`
	Week                  *int `json:"Week"`
	DayOfWeek             *int `json:"DayOfWeek"`
	Month                 *int `json:"Month"`
	Year                  *int `json:"Year"`
}

type AwardExpiryNotification struct {
	ExpiryNotifyValue *float64 `json:"ExpiryNotifyValue"`
	ExpiryNotifyType  *int     `json:"ExpiryNotifyType"`
}

type AwardFileData struct {
	Name string `json:"Name"`
	Path string `json:"Path"`
}
