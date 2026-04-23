package valence

import (
	"encoding/json"
	"strconv"
)

// ---- Common ----------------------------------------------------------------

// RichText holds a value in both plain-text and HTML forms.
type RichText struct {
	Text string `json:"Text"`
	Html string `json:"Html"`
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
	Identifier int64           `json:"Identifier,string"`
	Name       string          `json:"Name"`
	Code       string          `json:"Code"`
	Type       OrgUnitTypeInfo `json:"Type"`
	IsActive   bool            `json:"IsActive"`
}

type OrgUnitProperties struct {
	Identifier int64           `json:"Identifier,string"`
	Name       string          `json:"Name"`
	Code       string          `json:"Code"`
	Type       OrgUnitTypeInfo `json:"Type"`
	Parents    []int64         `json:"Parents"`
	IsActive   bool            `json:"IsActive"`
}

// ---- User ------------------------------------------------------------------

type WhoAmIUser struct {
	Identifier        string `json:"Identifier"`
	FirstName         string `json:"FirstName"`
	LastName          string `json:"LastName"`
	UniqueName        string `json:"UniqueName"`
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
	DisplayName      string             `json:"DisplayName"`
	LastAccessedDate *string            `json:"LastAccessedDate"`
	FirstLoginDate   *string            `json:"FirstLoginDate"`
}

// ---- Role ------------------------------------------------------------------

type RoleInfo struct {
	Id   int64  `json:"Id"`
	Code string `json:"Code"`
	Name string `json:"Name"`
}

// ---- Enrollment ------------------------------------------------------------

type OrgUnitInfo struct {
	Id       int64           `json:"Id"`
	Name     string          `json:"Name"`
	Code     string          `json:"Code"`
	Type     OrgUnitTypeInfo `json:"Type"`
	IsActive bool            `json:"IsActive"`
}

type MyOrgUnitInfo struct {
	OrgUnit  OrgUnitInfo `json:"OrgUnit"`
	Access   AccessInfo  `json:"Access"`
	IsPinned bool        `json:"IsPinned"`
}

type AccessInfo struct {
	IsActive  bool   `json:"IsActive"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	CanAccess bool   `json:"CanAccess"`
}

type OrgUnitUser struct {
	User         OrgUnitUserInfo `json:"User"`
	Identifier   int64           `json:"Identifier"`
	DisplayName  string          `json:"DisplayName"`
	UserName     string          `json:"UserName"`
	OrgDefinedId string          `json:"OrgDefinedId"`
	Role         RoleInfo        `json:"Role"`
}

type OrgUnitUserInfo struct {
	Identifier   string `json:"Identifier"`
	DisplayName  string `json:"DisplayName"`
	UserName     string `json:"UserName"`
	OrgDefinedId string `json:"OrgDefinedId"`
}

func (o *OrgUnitUser) UnmarshalJSON(data []byte) error {
	type alias OrgUnitUser
	var raw struct {
		User *OrgUnitUserInfo `json:"User"`
		*alias
	}
	raw.alias = (*alias)(o)

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.User == nil {
		return nil
	}

	o.User = *raw.User
	o.DisplayName = raw.User.DisplayName
	o.UserName = raw.User.UserName
	o.OrgDefinedId = raw.User.OrgDefinedId

	if raw.User.Identifier == "" {
		return nil
	}

	id, err := strconv.ParseInt(raw.User.Identifier, 10, 64)
	if err != nil {
		return nil
	}

	o.Identifier = id
	return nil
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
	Identifier      int64    `json:"Identifier,string"`
	Name            string   `json:"Name"`
	Code            string   `json:"Code"`
	IsActive        bool     `json:"IsActive"`
	Path            string   `json:"Path"`
	CourseTemplate  OrgUnit  `json:"CourseTemplate"`
	Semester        *OrgUnit `json:"Semester"`
	Department      *OrgUnit `json:"Department"`
	StartDate       *string  `json:"StartDate"`
	EndDate         *string  `json:"EndDate"`
	LocaleId        *int64   `json:"LocaleId"`
	ForceLocale     bool     `json:"ForceLocale"`
	ShowAddressBook bool     `json:"ShowAddressBook"`
	Description     RichText `json:"Description"`
	CanSelfRegister bool     `json:"CanSelfRegister"`
}

type CourseTemplate struct {
	Identifier int64  `json:"Identifier"`
	Name       string `json:"Name"`
	Code       string `json:"Code"`
	IsActive   bool   `json:"IsActive"`
	Path       string `json:"Path"`
	HomeUrl    string `json:"HomeUrl"`
}

// ---- Group -----------------------------------------------------------------

type GroupEnrollment struct {
	UserId int64 `json:"UserId"`
}

type Group struct {
	GroupId     int64    `json:"GroupId"`
	Name        string   `json:"Name"`
	Code        string   `json:"Code"`
	Description RichText `json:"Description"`
	Enrollments []int64  `json:"Enrollments"`
}

type GroupCategory struct {
	GroupCategoryId          int64    `json:"GroupCategoryId"`
	Name                     string   `json:"Name"`
	Description              RichText `json:"Description"`
	EnrollmentStyle          int      `json:"EnrollmentStyle"`
	EnrollmentQuantity       *int     `json:"EnrollmentQuantity"`
	AutoEnroll               bool     `json:"AutoEnroll"`
	RandomizeEnrollments     bool     `json:"RandomizeEnrollments"`
	NumberOfGroups           *int     `json:"NumberOfGroups"`
	MaxUsersPerGroup         *int     `json:"MaxUsersPerGroup"`
	AllocateAfterExpiry      bool     `json:"AllocateAfterExpiry"`
	SelfEnrollmentExpiryDate *string  `json:"SelfEnrollmentExpiryDate"`
	Groups                   []int64  `json:"Groups"`
	RestrictedByOrgUnitId    *int64   `json:"RestrictedByOrgUnitId"`
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
	EnrollmentStyle      string `json:"EnrollmentStyle"`
	EnrollmentQuantity   *int   `json:"EnrollmentQuantity"`
	AutoEnroll           bool   `json:"AutoEnroll"`
	RandomizeEnrollments bool   `json:"RandomizeEnrollments"`
}

// ---- Grade -----------------------------------------------------------------

type GradeSchemeRange struct {
	Symbol        string  `json:"Symbol"`
	Low           float64 `json:"Low"`
	High          float64 `json:"High"`
	AssignedValue float64 `json:"AssignedValue"`
}

type GradeScheme struct {
	Id        int64              `json:"Id"`
	Name      string             `json:"Name"`
	ShortName string             `json:"ShortName"`
	Ranges    []GradeSchemeEntry `json:"Ranges"`
}

type GradeSchemeEntry struct {
	PercentStart  float64  `json:"PercentStart"`
	Symbol        string   `json:"Symbol"`
	AssignedValue *float64 `json:"AssignedValue"`
	Colour        string   `json:"Colour"`
}

type GradeSetupInfo struct {
	GradingSystem        string `json:"GradingSystem"`
	IsNullGradeZero      bool   `json:"IsNullGradeZero"`
	DefaultGradeSchemeId int64  `json:"DefaultGradeSchemeId"`
}

type GradeSchemeInfo struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

type GradeObject struct {
	MaxPoints             *float64         `json:"MaxPoints"`
	CanExceedMaxPoints    bool             `json:"CanExceedMaxPoints"`
	IsBonus               bool             `json:"IsBonus"`
	ExcludeFromFinalGrade bool             `json:"ExcludeFromFinalGrade"`
	GradeSchemeId         *int64           `json:"GradeSchemeId"`
	GradeSchemeUrl        string           `json:"GradeSchemeUrl"`
	GradeScheme           *GradeSchemeInfo `json:"GradeScheme"`
	Id                    int64            `json:"Id"`
	Name                  string           `json:"Name"`
	ShortName             string           `json:"ShortName"`
	GradeType             string           `json:"GradeType"`
	CategoryId            *int64           `json:"CategoryId"`
	Description           RichText         `json:"Description"`
	GradeObjectTypeId     int              `json:"GradeObjectTypeId"`
	ActivityId            *string          `json:"ActivityId"`
	AssociatedTool        *AssociatedTool  `json:"AssociatedTool"`
	IsHidden              bool             `json:"IsHidden"`
}

type AssociatedTool struct {
	ToolId     int64 `json:"ToolId"`
	ToolItemId int64 `json:"ToolItemId"`
}

type GradeCategory struct {
	Id                     int64    `json:"Id"`
	Name                   string   `json:"Name"`
	ShortName              string   `json:"ShortName"`
	CanExceedMax           bool     `json:"CanExceedMax"`
	ExcludeFromFinalGrade  bool     `json:"ExcludeFromFinalGrade"`
	StartDate              *string  `json:"StartDate"`
	EndDate                *string  `json:"EndDate"`
	MaxPoints              *float64 `json:"MaxPoints"`
	Weight                 *float64 `json:"Weight"`
	BoostUserScore         bool     `json:"BoostUserScore"`
	AutoPoints             bool     `json:"AutoPoints"`
	WeightDistributionType int      `json:"WeightDistributionType"`
	NumberOfHighestToDrop  *int     `json:"NumberOfHighestToDrop"`
	NumberOfLowestToDrop   *int     `json:"NumberOfLowestToDrop"`
}

type GradeValue struct {
	UserId                int64     `json:"UserId"`
	OrgUnitId             int64     `json:"OrgUnitId"`
	DisplayedGrade        string    `json:"DisplayedGrade"`
	GradeObjectIdentifier int64     `json:"GradeObjectIdentifier"`
	GradeObjectName       string    `json:"GradeObjectName"`
	GradeObjectType       int       `json:"GradeObjectType"`
	GradeObjectTypeName   string    `json:"GradeObjectTypeName"`
	PointsNumerator       *float64  `json:"PointsNumerator"`
	PointsDenominator     *float64  `json:"PointsDenominator"`
	WeightedDenominator   *float64  `json:"WeightedDenominator"`
	WeightedNumerator     *float64  `json:"WeightedNumerator"`
	Comments              *RichText `json:"Comments"`
	PrivateComments       *RichText `json:"PrivateComments"`
	LastModified          *string   `json:"LastModified"`
	LastModifiedBy        *UserData `json:"LastModifiedBy"`
	Released              bool      `json:"Released"`
}

// GradeUserRef is the user reference embedded in a GradeValueEntry.
type GradeUserRef struct {
	Identifier  string `json:"Identifier"`
	DisplayName string `json:"DisplayName"`
}

// GradeValueData is the grade value data embedded in a GradeValueEntry.
type GradeValueData struct {
	PointsNumerator       *float64 `json:"PointsNumerator"`
	PointsDenominator     *float64 `json:"PointsDenominator"`
	WeightedNumerator     *float64 `json:"WeightedNumerator"`
	WeightedDenominator   *float64 `json:"WeightedDenominator"`
	DisplayedGrade        string   `json:"DisplayedGrade"`
	GradeObjectIdentifier string   `json:"GradeObjectIdentifier"`
	GradeObjectName       string   `json:"GradeObjectName"`
	GradeObjectType       int      `json:"GradeObjectType"`
}

// GradeValueEntry is one item in the ObjectListPage returned by the grade values endpoint.
type GradeValueEntry struct {
	User       GradeUserRef   `json:"User"`
	GradeValue GradeValueData `json:"GradeValue"`
}

type FinalGradeValue struct {
	UserId                int64     `json:"UserId"`
	OrgUnitId             int64     `json:"OrgUnitId"`
	DisplayedGrade        string    `json:"DisplayedGrade"`
	GradeObjectIdentifier string    `json:"GradeObjectIdentifier"`
	GradeObjectName       string    `json:"GradeObjectName"`
	PointsNumerator       *float64  `json:"PointsNumerator"`
	PointsDenominator     *float64  `json:"PointsDenominator"`
	WeightedDenominator   *float64  `json:"WeightedDenominator"`
	WeightedNumerator     *float64  `json:"WeightedNumerator"`
	Comments              *RichText `json:"Comments"`
	PrivateComments       *RichText `json:"PrivateComments"`
	Released              bool      `json:"Released"`
}

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
	OrgDefinedId             string  `json:"OrgDefinedId"`
	Email                    string  `json:"Email"`
	FirstName                string  `json:"FirstName"`
	LastName                 string  `json:"LastName"`
	RoleId                   int64   `json:"RoleId"`
	ClasslistRoleDisplayName string  `json:"ClasslistRoleDisplayName"`
	LastAccessed             *string `json:"LastAccessed"`
	IsOnline                 bool    `json:"IsOnline"`
	Pronouns                 *string `json:"Pronouns"`
}

// ---- News ------------------------------------------------------------------

type NewsItem struct {
	Id                        int64    `json:"Id"`
	Title                     string   `json:"Title"`
	Body                      RichText `json:"Body"`
	StartDate                 *string  `json:"StartDate"`
	EndDate                   *string  `json:"EndDate"`
	IsGlobal                  bool     `json:"IsGlobal"`
	IsPublished               bool     `json:"IsPublished"`
	ShowOnlyInCourseOfferings bool     `json:"ShowOnlyInCourseOfferings"`
}

// ---- Quiz ------------------------------------------------------------------

type QuizReadData struct {
	QuizId              int64        `json:"QuizId"`
	Name                string       `json:"Name"`
	IsActive            bool         `json:"IsActive"`
	SortOrder           int          `json:"SortOrder"`
	AutoExportToGrades  *bool        `json:"AutoExportToGrades"`
	GradeItemId         *int64       `json:"GradeItemId"`
	IsAutoSetGraded     bool         `json:"IsAutoSetGraded"`
	SubmissionTimeLimit TimeLimit    `json:"SubmissionTimeLimit"`
	StartDate           *string      `json:"StartDate"`
	EndDate             *string      `json:"EndDate"`
	DueDate             *string      `json:"DueDate"`
	DisplayInCalendar   bool         `json:"DisplayInCalendar"`
	Instructions        Instructions `json:"Instructions"`
	Description         Description  `json:"Description"`
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
	AttemptId     int64    `json:"AttemptId"`
	UserId        int64    `json:"UserId"`
	AttemptNumber int      `json:"AttemptNumber"`
	TimeStarted   string   `json:"TimeStarted"`
	TimeCompleted *string  `json:"TimeCompleted"`
	Score         *float64 `json:"Score"`
	IsInProgress  bool     `json:"IsInProgress"`
}

type QuizQuestion struct {
	QuestionId     int64           `json:"QuestionId"`
	Name           string          `json:"Name"`
	QuestionText   RichText        `json:"QuestionText"`
	Points         float64         `json:"Points"`
	Difficulty     int             `json:"Difficulty"`
	Bonus          bool            `json:"Bonus"`
	Mandatory      bool            `json:"Mandatory"`
	QuestionTypeId int             `json:"QuestionTypeId"`
	Hint           RichText        `json:"Hint"`
	Feedback       RichText        `json:"Feedback"`
	LastModified   *string         `json:"LastModified"`
	LastModifiedBy *int64          `json:"LastModifiedBy"`
	SectionId      *int64          `json:"SectionId"`
	QuestionInfo   json.RawMessage `json:"QuestionInfo"`
}

type QuizSpecialAccessData struct {
	UserId    int64   `json:"UserId"`
	StartDate *string `json:"StartDate"`
	EndDate   *string `json:"EndDate"`
	DueDate   *string `json:"DueDate"`
	IsActive  bool    `json:"IsActive"`
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
	PostId       int64    `json:"PostId"`
	TopicId      int64    `json:"TopicId"`
	ForumId      int64    `json:"ForumId"`
	ParentPostId *int64   `json:"ParentPostId"`
	Subject      string   `json:"Subject"`
	Message      RichText `json:"Message"`
	IsAnonymous  bool     `json:"IsAnonymous"`
	IsApproved   bool     `json:"IsApproved"`
	IsDeleted    bool     `json:"IsDeleted"`
	ThreadId     int64    `json:"ThreadId"`
	UserId       int64    `json:"PostingUserId"`
	DatePosted   string   `json:"DatePosted"`
	LastModified *string  `json:"LastModified"`
}

// ---- Dropbox ---------------------------------------------------------------

type RubricLevel struct {
	Id     int64   `json:"Id"`
	Name   string  `json:"Name"`
	Points float64 `json:"Points"`
}

type RubricCell struct {
	Description RichText `json:"Description"`
	Feedback    RichText `json:"Feedback"`
	Points      *float64 `json:"Points"`
	LevelId     int64    `json:"LevelId"`
}

type RubricCriterion struct {
	Id    int64        `json:"Id"`
	Name  string       `json:"Name"`
	Cells []RubricCell `json:"Cells"`
}

type RubricCriteriaGroup struct {
	Name       string            `json:"Name"`
	Levels     []RubricLevel     `json:"Levels"`
	Criteria   []RubricCriterion `json:"Criteria"`
	LevelSetId int64             `json:"LevelSetId"`
}

type RubricOverallLevel struct {
	Id          int64    `json:"Id"`
	Name        string   `json:"Name"`
	RangeStart  float64  `json:"RangeStart"`
	Description RichText `json:"Description"`
	Feedback    RichText `json:"Feedback"`
}

type DropboxRubric struct {
	RubricId       int64                 `json:"RubricId"`
	Name           string                `json:"Name"`
	Description    RichText              `json:"Description"`
	RubricType     int                   `json:"RubricType"`
	ScoringMethod  int                   `json:"ScoringMethod"`
	CriteriaGroups []RubricCriteriaGroup `json:"CriteriaGroups"`
	OverallLevels  []RubricOverallLevel  `json:"OverallLevels"`
}

type DropboxAssessment struct {
	ScoreDenominator *float64        `json:"ScoreDenominator"`
	Rubrics          []DropboxRubric `json:"Rubrics"`
}

type DropboxFolder struct {
	Id                        int64             `json:"Id"`
	CategoryId                *int64            `json:"CategoryId"`
	Name                      string            `json:"Name"`
	CustomInstructions        RichText          `json:"CustomInstructions"`
	Attachments               []SubmissionFile  `json:"Attachments"`
	TotalFiles                int               `json:"TotalFiles"`
	UnreadFiles               int               `json:"UnreadFiles"`
	FlaggedFiles              int               `json:"FlaggedFiles"`
	TotalUsers                int               `json:"TotalUsers"`
	TotalUsersWithSubmissions int               `json:"TotalUsersWithSubmissions"`
	TotalUsersWithFeedback    int               `json:"TotalUsersWithFeedback"`
	IsHidden                  bool              `json:"IsHidden"`
	IsAnonymous               bool              `json:"IsAnonymous"`
	DueDate                   *string           `json:"DueDate"`
	DisplayInCalendar         bool              `json:"DisplayInCalendar"`
	DropboxType               int               `json:"DropboxType"`
	SubmissionType            *int              `json:"SubmissionType"`
	CompletionType            *int              `json:"CompletionType"`
	GroupTypeId               *int64            `json:"GroupTypeId"`
	GradeItemId               *int64            `json:"GradeItemId"`
	ActivityId                *string           `json:"ActivityId"`
	Assessment                DropboxAssessment `json:"Assessment"`
}

type DropboxCategory struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

type SubmissionFile struct {
	FileId   int64  `json:"FileId"`
	FileName string `json:"FileName"`
	Size     int64  `json:"Size"`
}

// DropboxEntity is the user entity at the top level of a submission group.
// User entities use "DisplayName"; Group entities use "Name".
type DropboxEntity struct {
	DisplayName string `json:"DisplayName"`
	Name        string `json:"Name"`
	EntityId    int64  `json:"EntityId"`
	EntityType  string `json:"EntityType"`
	Active      bool   `json:"Active"`
}

// DropboxFeedback holds instructor feedback for a user's folder submission.
type DropboxFeedback struct {
	Score        *float64                `json:"Score"`
	Feedback     *RichText               `json:"Feedback"`
	IsGraded     bool                    `json:"IsGraded"`
	GradedSymbol *string                 `json:"GradedSymbol"`
	Files        []DropboxSubmissionFile `json:"Files"`
}

// DropboxSubmitter is the per-submission submitter reference (Identifier is a string user ID).
type DropboxSubmitter struct {
	Identifier  string `json:"Identifier"`
	DisplayName string `json:"DisplayName"`
}

// DropboxSubmissionFile is a file attached to a dropbox submission entry.
type DropboxSubmissionFile struct {
	FileId    int64  `json:"FileId"`
	FileName  string `json:"FileName"`
	Size      int64  `json:"Size"`
	IsRead    bool   `json:"IsRead"`
	IsFlagged bool   `json:"IsFlagged"`
	IsDeleted bool   `json:"IsDeleted"`
}

// DropboxSubmissionEntry is a single submission within a UserSubmissions group.
type DropboxSubmissionEntry struct {
	Id             int64                   `json:"Id"`
	SubmittedBy    DropboxSubmitter        `json:"SubmittedBy"`
	SubmissionDate string                  `json:"SubmissionDate"`
	Comment        RichText                `json:"Comment"`
	Files          []DropboxSubmissionFile `json:"Files"`
}

// UserSubmissions groups all submissions (and feedback) for one user in a dropbox folder.
// This is the element type returned by GetDropboxSubmissions.
type UserSubmissions struct {
	Entity         DropboxEntity            `json:"Entity"`
	Status         int                      `json:"Status"`
	Feedback       DropboxFeedback          `json:"Feedback"`
	Submissions    []DropboxSubmissionEntry `json:"Submissions"`
	CompletionDate string                   `json:"CompletionDate"`
}

// ---- Survey ----------------------------------------------------------------

type Survey struct {
	SurveyId          int64        `json:"SurveyId"`
	Name              string       `json:"Name"`
	IsActive          bool         `json:"IsActive"`
	IsAnonymous       bool         `json:"IsAnonymous"`
	StartDate         *string      `json:"StartDate"`
	EndDate           *string      `json:"EndDate"`
	DueDate           *string      `json:"DueDate"`
	DisplayInCalendar bool         `json:"DisplayInCalendar"`
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
	SelfAssessmentId int64        `json:"SelfAssessmentId"`
	Name             string       `json:"Name"`
	IsActive         bool         `json:"IsActive"`
	StartDate        *string      `json:"StartDate"`
	EndDate          *string      `json:"EndDate"`
	Instructions     Instructions `json:"Instructions"`
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
	LinkId      int64  `json:"LinkId"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Url         string `json:"Url"`
}

type LTISharingData struct {
	OrgUnitId int64  `json:"OrgUnitId"`
	Name      string `json:"Name"`
}

type LTICustomParameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type LTIToolProvider struct {
	TpId                int64                `json:"TpId"`
	Name                string               `json:"Name"`
	Description         string               `json:"Description"`
	Domain              string               `json:"Domain"`
	Url                 string               `json:"Url"`
	Key                 string               `json:"Key"`
	Secret              string               `json:"Secret"`
	SendTcInfo          bool                 `json:"SendTcInfo"`
	SendContextInfo     bool                 `json:"SendContextInfo"`
	SendUserId          bool                 `json:"SendUserId"`
	SendUserName        bool                 `json:"SendUserName"`
	SendUserEmail       bool                 `json:"SendUserEmail"`
	SendLinkTitle       bool                 `json:"SendLinkTitle"`
	SendLinkDescription bool                 `json:"SendLinkDescription"`
	SendD2LUserName     bool                 `json:"SendD2LUserName"`
	SendD2LOrgDefinedId bool                 `json:"SendD2LOrgDefinedId"`
	SendD2LOrgRoleId    bool                 `json:"SendD2LOrgRoleId"`
	SendSectionCode     bool                 `json:"SendSectionCode"`
	CustomParameters    []LTICustomParameter `json:"CustomParameters"`
}

type LTIDeploymentSharingData struct {
	OrgUnitId int64  `json:"OrgUnitId"`
	Name      string `json:"Name"`
	IsShared  bool   `json:"IsShared"`
}

// ---- Tools -----------------------------------------------------------------

// OrgUnitInformation block returned by GET /d2l/api/lp/(version)/tools/orgUnits/(orgUnitId)
type ToolInfo struct {
	ToolId           string `json:"ToolId"`
	DisplayName      string `json:"DisplayName"`
	OrgUnitId        int64  `json:"OrgUnitId"`
	Status           bool   `json:"Status"`
	CustomNavbarName string `json:"CustomNavbarName"`
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
	IsEnabled      bool     `json:"IsEnabled"`
	Type           *int     `json:"Type"`
	StartDate      *string  `json:"StartDate"`
	EndDate        *string  `json:"EndDate"`
	RepeatsEvery   *int     `json:"RepeatsEvery"`
	RepeatsOnDay   *int     `json:"RepeatsOnDay"`
	RepeatsOnDays  []string `json:"RepeatsOnDays"`
	RepeatsOnMonth *int     `json:"RepeatsOnMonth"`
}

type IntelligentAgentAction struct {
	RepeatType       int                           `json:"RepeatType"`
	EmailAction      *IntelligentAgentEmailAction  `json:"EmailAction"`
	EnrollmentAction *IntelligentAgentEnrollAction `json:"EnrollmentAction"`
}

type IntelligentAgentEmailAction struct {
	IsEnabled bool   `json:"IsEnabled"`
	To        string `json:"To"`
	Cc        string `json:"Cc"`
	Bcc       string `json:"Bcc"`
	Subject   string `json:"Subject"`
	Message   string `json:"Message"`
	IsHtml    bool   `json:"IsHtml"`
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
	VariableUUID string `json:"VariableUUID"`
	OrgUnitId    int64  `json:"OrgUnitId"`
	Value        string `json:"Value"`
}

// ---- Course Import ---------------------------------------------------------

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
	IssuedId  int64   `json:"IssuedId"`
	UserId    int64   `json:"UserId"`
	BadgeId   int64   `json:"BadgeId"`
	BadgeName string  `json:"BadgeName"`
	IssueDate string  `json:"IssueDate"`
	Expiry    *string `json:"Expiry"`
	Evidence  string  `json:"Evidence"`
	Narrative string  `json:"Narrative"`
}
