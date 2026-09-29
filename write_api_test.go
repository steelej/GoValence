package valence

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type noopAuth struct{}

func (noopAuth) AuthenticateRequest(*http.Request) error { return nil }

func newTestClient(handler http.HandlerFunc) (*Client, func()) {
	server := httptest.NewServer(handler)
	client := New(Config{
		BaseURL:   server.URL,
		Auth:      noopAuth{},
		LPVersion: "1.50",
		LEVersion: "1.75",
	})
	return client, server.Close
}

func TestCreateCourseUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/courses/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body CreateCourseOffering
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.Name != "Example Course" || body.CourseTemplateId != 77 {
			t.Fatalf("unexpected body: %+v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Identifier":"12345","Name":"Example Course","Code":"EXAMPLE","IsActive":true}`))
	})
	defer closeServer()

	canSelfRegister := false
	got, err := client.CreateCourse(CreateCourseOffering{
		Name:             "Example Course",
		Code:             "EXAMPLE",
		CourseTemplateId: 77,
		Description:      RichTextInput{Content: "", Type: "Text"},
		CanSelfRegister:  &canSelfRegister,
		IsActive:         true,
	})
	if err != nil {
		t.Fatalf("CreateCourse returned error: %v", err)
	}
	if got.Identifier != 12345 {
		t.Fatalf("Identifier = %d, want 12345", got.Identifier)
	}
}

func TestCreateCourseCopyJobUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/le/1.75/import/12345/copy/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body CreateCopyJobRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.SourceOrgUnitId != 77 {
			t.Fatalf("SourceOrgUnitId = %d, want 77", body.SourceOrgUnitId)
		}
		if body.CallbackUrl == nil || *body.CallbackUrl != "https://example.test/ccjr/12345" {
			t.Fatalf("CallbackUrl = %v, want callback", body.CallbackUrl)
		}
		if body.Components != nil {
			t.Fatalf("Components = %#v, want nil for all components", body.Components)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"JobToken":"copy-token"}`))
	})
	defer closeServer()

	callbackURL := "https://example.test/ccjr/12345"
	got, err := client.CreateCourseCopyJob(12345, CreateCopyJobRequest{
		SourceOrgUnitId: 77,
		CallbackUrl:     &callbackURL,
	})
	if err != nil {
		t.Fatalf("CreateCourseCopyJob returned error: %v", err)
	}
	if got.JobToken != "copy-token" {
		t.Fatalf("JobToken = %q, want copy-token", got.JobToken)
	}
}

func TestCreateUserUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/users/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body CreateUserData
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.UserName != "testUser_123_1" || body.RoleId != 88 || !body.IsActive || body.SendCreationEmail {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body.OrgDefinedId == nil || *body.OrgDefinedId != body.UserName {
			t.Fatalf("OrgDefinedId = %v, want username", body.OrgDefinedId)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"OrgId":1,"UserId":456,"FirstName":"Test Student","LastName":"Learner","UserName":"testUser_123_1","ExternalEmail":"noemail@example.com","OrgDefinedId":"testUser_123_1","Activation":{"IsActive":true}}`))
	})
	defer closeServer()

	orgDefinedID := "testUser_123_1"
	email := "noemail@example.com"
	got, err := client.CreateUser(CreateUserData{
		OrgDefinedId:      &orgDefinedID,
		FirstName:         "Test Student",
		LastName:          "Learner",
		ExternalEmail:     &email,
		UserName:          "testUser_123_1",
		RoleId:            88,
		IsActive:          true,
		SendCreationEmail: false,
	})
	if err != nil {
		t.Fatalf("CreateUser returned error: %v", err)
	}
	if got.UserId != 456 || got.UserName != "testUser_123_1" {
		t.Fatalf("unexpected user: %+v", got)
	}
}

func TestCreateEnrollmentUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/enrollments/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body CreateEnrollmentData
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.OrgUnitId != 123 || body.UserId != 456 || body.RoleId != 88 {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body.SendEnrollmentEmail != nil {
			t.Fatalf("SendEnrollmentEmail = %v, want omitted nil", body.SendEnrollmentEmail)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"OrgUnitId":123,"UserId":456,"RoleId":88,"IsCascading":false}`))
	})
	defer closeServer()

	got, err := client.CreateEnrollment(CreateEnrollmentData{
		OrgUnitId: 123,
		UserId:    456,
		RoleId:    88,
	})
	if err != nil {
		t.Fatalf("CreateEnrollment returned error: %v", err)
	}
	if got.OrgUnitId != 123 || got.UserId != 456 || got.RoleId != 88 || got.IsCascading {
		t.Fatalf("unexpected enrollment: %+v", got)
	}
}

func TestGetUserOrgUnitEnrollmentUsesDocumentedResponse(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/d2l/api/lp/1.50/enrollments/orgUnits/123/users/456" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"OrgUnitId":123,"UserId":456,"RoleId":88,"IsCascading":true}`))
	})
	defer closeServer()

	got, err := client.GetUserOrgUnitEnrollment(123, 456)
	if err != nil {
		t.Fatalf("GetUserOrgUnitEnrollment returned error: %v", err)
	}
	if got.OrgUnitId != 123 || got.UserId != 456 || got.RoleId != 88 || !got.IsCascading {
		t.Fatalf("unexpected enrollment: %+v", got)
	}
}

func TestEnrollUserInGroupUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/10/groupcategories/20/groups/30/enrollments/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body GroupEnrollment
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.UserId != 40 {
			t.Fatalf("UserId = %d, want 40", body.UserId)
		}

		w.WriteHeader(http.StatusOK)
	})
	defer closeServer()

	if err := client.EnrollUserInGroup(10, 20, 30, GroupEnrollment{UserId: 40}); err != nil {
		t.Fatalf("EnrollUserInGroup returned error: %v", err)
	}
}

func TestDeleteEnrollmentUsesDocumentedRoute(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("method = %s, want DELETE", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/enrollments/orgUnits/123/users/456" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"OrgUnitId":123,"UserId":456,"RoleId":88,"IsCascading":true}`))
	})
	defer closeServer()

	got, err := client.DeleteEnrollment(123, 456)
	if err != nil {
		t.Fatalf("DeleteEnrollment returned error: %v", err)
	}
	if got.OrgUnitId != 123 || got.UserId != 456 || got.RoleId != 88 || !got.IsCascading {
		t.Fatalf("unexpected deleted enrollment: %+v", got)
	}
}

func TestCreateSectionUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/10/sections/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body SectionData
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.Name != "CISC 101 Section 001" || body.Code != "123.1.2261.1.001" {
			t.Fatalf("unexpected body: %+v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"SectionId":99,"Name":"CISC 101 Section 001","Code":"123.1.2261.1.001","Description":{"Text":"","Html":""},"Enrollments":[]}`))
	})
	defer closeServer()

	got, err := client.CreateSection(10, SectionData{
		Name:        "CISC 101 Section 001",
		Code:        "123.1.2261.1.001",
		Description: RichTextInput{Content: "", Type: "Text"},
	})
	if err != nil {
		t.Fatalf("CreateSection returned error: %v", err)
	}
	if got.SectionId != 99 {
		t.Fatalf("SectionId = %d, want 99", got.SectionId)
	}
}

func TestCreateGroupCategoryUsesAsyncJobShape(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/10/groupcategories/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body GroupCategoryData
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.Name != "Tutorial Groups" || body.NumberOfGroups == nil || *body.NumberOfGroups != 1 {
			t.Fatalf("unexpected body: %+v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"OrgUnitId":10,"CategoryId":20,"SubmitDate":"2026-07-03T00:00:00.000Z","Status":0}`))
	})
	defer closeServer()

	numberOfGroups := 1
	got, err := client.CreateGroupCategory(10, GroupCategoryData{
		Name:                 "Tutorial Groups",
		Description:          RichTextInput{Content: "", Type: "Text"},
		EnrollmentStyle:      0,
		NumberOfGroups:       &numberOfGroups,
		RandomizeEnrollments: false,
	})
	if err != nil {
		t.Fatalf("CreateGroupCategory returned error: %v", err)
	}
	if got.CategoryId != 20 || got.Status != 0 {
		t.Fatalf("unexpected job: %+v", got)
	}
}

func TestGetGroupCategoryStatusUsesDocumentedRoute(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/10/groupcategories/20/status" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":1}`))
	})
	defer closeServer()

	got, err := client.GetGroupCategoryStatus(10, 20)
	if err != nil {
		t.Fatalf("GetGroupCategoryStatus returned error: %v", err)
	}
	if got.Status != 1 {
		t.Fatalf("Status = %d, want 1", got.Status)
	}
}

func TestGroupCategoryEnrollmentStyleAcceptsStringOrNumber(t *testing.T) {
	for name, payload := range map[string]string{
		"string": `{"GroupCategoryId":5,"Name":"Groups","EnrollmentStyle":"NumberOfGroupsNoEnrollment"}`,
		"number": `{"GroupCategoryId":5,"Name":"Groups","EnrollmentStyle":0}`,
	} {
		t.Run(name, func(t *testing.T) {
			var category GroupCategory
			if err := json.Unmarshal([]byte(payload), &category); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if category.EnrollmentStyle == "" {
				t.Fatalf("EnrollmentStyle was empty for payload %s", payload)
			}
		})
	}
}

func TestDeleteGroupUsesDocumentedRoute(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("method = %s, want DELETE", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/10/groupcategories/20/groups/30" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	})
	defer closeServer()

	if err := client.DeleteGroup(10, 20, 30); err != nil {
		t.Fatalf("DeleteGroup returned error: %v", err)
	}
}

func TestGetUserByUserNameUsesDocumentedQuery(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/users/" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("userName"); got != "57js" {
			t.Fatalf("userName = %s, want 57js", got)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"UserId":123,"UserName":"57js","DisplayName":"Test Student","OrgDefinedId":"07367455"}`))
	})
	defer closeServer()

	user, err := client.GetUserByUserName("57js")
	if err != nil {
		t.Fatalf("GetUserByUserName returned error: %v", err)
	}
	if user.UserId != 123 || user.UserName != "57js" {
		t.Fatalf("unexpected user: %+v", user)
	}
}

func TestUpdateConfigVariableValueUsesDocumentedRouteAndBody(t *testing.T) {
	const uuid = "c47f0d66-bfa6-4d37-b30e-5925d2e21a72"

	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("method = %s, want PUT", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/configVariables/"+uuid+"/values/orgUnits/123" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body SpecifiedOrgUnitValue
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.OrgUnitValue == nil || *body.OrgUnitValue != "on" {
			t.Fatalf("OrgUnitValue = %v, want on", body.OrgUnitValue)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"VariableUUID":"` + uuid + `","OrgUnitId":123,"Value":"on"}`))
	})
	defer closeServer()

	value := "on"
	got, err := client.UpdateConfigVariableValue(uuid, 123, SpecifiedOrgUnitValue{OrgUnitValue: &value})
	if err != nil {
		t.Fatalf("UpdateConfigVariableValue returned error: %v", err)
	}
	if got.Value == nil || *got.Value != "on" {
		t.Fatalf("Value = %v, want on", got.Value)
	}
}

func TestGetEffectiveConfigVariableValueUsesDocumentedRoute(t *testing.T) {
	const uuid = "c47f0d66-bfa6-4d37-b30e-5925d2e21a72"

	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/configVariables/"+uuid+"/effectiveValues/orgUnits/123" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"VariableUUID":"` + uuid + `","OrgUnitId":123,"Value":"on"}`))
	})
	defer closeServer()

	got, err := client.GetEffectiveConfigVariableValue(uuid, 123)
	if err != nil {
		t.Fatalf("GetEffectiveConfigVariableValue returned error: %v", err)
	}
	if got.Value == nil || *got.Value != "on" {
		t.Fatalf("Value = %v, want on", got.Value)
	}
}

func TestUpdateOrgUnitToolStatusUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("method = %s, want PUT", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/tools/orgUnits/123/tool/456" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body UpdateStatus
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if !body.Status {
			t.Fatal("Status = false, want true")
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ToolId":"456","Status":true}`))
	})
	defer closeServer()

	got, err := client.UpdateOrgUnitToolStatus(123, 456, UpdateStatus{Status: true})
	if err != nil {
		t.Fatalf("UpdateOrgUnitToolStatus returned error: %v", err)
	}
	if !got.Status {
		t.Fatal("Status = false, want true")
	}
}

func TestAddLTIAdvantageDeploymentSharingRuleUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/le/1.75/ltiadvantage/deployment/7bb6d98c-d27b-45d1-9f5a-2e3371115d35/sharing/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body LTIAdvantageCreateSharingRuleData
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body.SharingOrgUnitId != 456 || !body.ShareWithOrgUnit || body.ShareWithDescendants {
			t.Fatalf("unexpected body: %+v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"SharingOrgUnitId":456,"ShareWithOrgUnit":true,"ShareWithDescendants":false}`))
	})
	defer closeServer()

	got, err := client.AddLTIAdvantageDeploymentSharingRule("7bb6d98c-d27b-45d1-9f5a-2e3371115d35", LTIAdvantageCreateSharingRuleData{
		SharingOrgUnitId:     456,
		ShareWithOrgUnit:     true,
		ShareWithDescendants: false,
	})
	if err != nil {
		t.Fatalf("AddLTIAdvantageDeploymentSharingRule returned error: %v", err)
	}
	if got.SharingOrgUnitId != 456 || !got.ShareWithOrgUnit {
		t.Fatalf("got = %+v", got)
	}
}

func TestAddOrgUnitParentUsesDocumentedRouteAndBody(t *testing.T) {
	client, closeServer := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/d2l/api/lp/1.50/orgstructure/123/parents/" {
			t.Fatalf("path = %s", r.URL.Path)
		}

		var body int64
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body != 456 {
			t.Fatalf("body = %d, want 456", body)
		}

		w.WriteHeader(http.StatusOK)
	})
	defer closeServer()

	if err := client.AddOrgUnitParent(123, 456); err != nil {
		t.Fatalf("AddOrgUnitParent returned error: %v", err)
	}
}
