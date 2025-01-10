// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pipe-cd/pipecd/pkg/datastore (interfaces: ProjectStore,PipedStore,ApplicationStore,DeploymentStore,CommandStore)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod --package=datastoretest --destination=pkg/datastore/datastoretest/datastore.mock.go github.com/pipe-cd/pipecd/pkg/datastore ProjectStore,PipedStore,ApplicationStore,DeploymentStore,CommandStore
//

// Package datastoretest is a generated GoMock package.
package datastoretest

import (
	context "context"
	reflect "reflect"
	time "time"

	datastore "github.com/pipe-cd/pipecd/pkg/datastore"
	model "github.com/pipe-cd/pipecd/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectStore is a mock of ProjectStore interface.
type MockProjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockProjectStoreMockRecorder
	isgomock struct{}
}

// MockProjectStoreMockRecorder is the mock recorder for MockProjectStore.
type MockProjectStoreMockRecorder struct {
	mock *MockProjectStore
}

// NewMockProjectStore creates a new mock instance.
func NewMockProjectStore(ctrl *gomock.Controller) *MockProjectStore {
	mock := &MockProjectStore{ctrl: ctrl}
	mock.recorder = &MockProjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectStore) EXPECT() *MockProjectStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockProjectStore) Add(ctx context.Context, proj *model.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, proj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockProjectStoreMockRecorder) Add(ctx, proj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockProjectStore)(nil).Add), ctx, proj)
}

// AddProjectRBACRole mocks base method.
func (m *MockProjectStore) AddProjectRBACRole(ctx context.Context, id, name string, policies []*model.ProjectRBACPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProjectRBACRole", ctx, id, name, policies)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProjectRBACRole indicates an expected call of AddProjectRBACRole.
func (mr *MockProjectStoreMockRecorder) AddProjectRBACRole(ctx, id, name, policies any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjectRBACRole", reflect.TypeOf((*MockProjectStore)(nil).AddProjectRBACRole), ctx, id, name, policies)
}

// AddProjectUserGroup mocks base method.
func (m *MockProjectStore) AddProjectUserGroup(ctx context.Context, id, sso, role string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProjectUserGroup", ctx, id, sso, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProjectUserGroup indicates an expected call of AddProjectUserGroup.
func (mr *MockProjectStoreMockRecorder) AddProjectUserGroup(ctx, id, sso, role any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjectUserGroup", reflect.TypeOf((*MockProjectStore)(nil).AddProjectUserGroup), ctx, id, sso, role)
}

// DeleteProjectRBACRole mocks base method.
func (m *MockProjectStore) DeleteProjectRBACRole(ctx context.Context, id, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectRBACRole", ctx, id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProjectRBACRole indicates an expected call of DeleteProjectRBACRole.
func (mr *MockProjectStoreMockRecorder) DeleteProjectRBACRole(ctx, id, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectRBACRole", reflect.TypeOf((*MockProjectStore)(nil).DeleteProjectRBACRole), ctx, id, name)
}

// DeleteProjectUserGroup mocks base method.
func (m *MockProjectStore) DeleteProjectUserGroup(ctx context.Context, id, sso string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectUserGroup", ctx, id, sso)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProjectUserGroup indicates an expected call of DeleteProjectUserGroup.
func (mr *MockProjectStoreMockRecorder) DeleteProjectUserGroup(ctx, id, sso any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectUserGroup", reflect.TypeOf((*MockProjectStore)(nil).DeleteProjectUserGroup), ctx, id, sso)
}

// DisableStaticAdmin mocks base method.
func (m *MockProjectStore) DisableStaticAdmin(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableStaticAdmin", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableStaticAdmin indicates an expected call of DisableStaticAdmin.
func (mr *MockProjectStoreMockRecorder) DisableStaticAdmin(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableStaticAdmin", reflect.TypeOf((*MockProjectStore)(nil).DisableStaticAdmin), ctx, id)
}

// EnableStaticAdmin mocks base method.
func (m *MockProjectStore) EnableStaticAdmin(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableStaticAdmin", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableStaticAdmin indicates an expected call of EnableStaticAdmin.
func (mr *MockProjectStoreMockRecorder) EnableStaticAdmin(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableStaticAdmin", reflect.TypeOf((*MockProjectStore)(nil).EnableStaticAdmin), ctx, id)
}

// Get mocks base method.
func (m *MockProjectStore) Get(ctx context.Context, id string) (*model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProjectStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProjectStore)(nil).Get), ctx, id)
}

// List mocks base method.
func (m *MockProjectStore) List(ctx context.Context, opts datastore.ListOptions) ([]model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].([]model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProjectStoreMockRecorder) List(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectStore)(nil).List), ctx, opts)
}

// UpdateProjectRBACConfig mocks base method.
func (m *MockProjectStore) UpdateProjectRBACConfig(ctx context.Context, id string, rbac *model.ProjectRBACConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectRBACConfig", ctx, id, rbac)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectRBACConfig indicates an expected call of UpdateProjectRBACConfig.
func (mr *MockProjectStoreMockRecorder) UpdateProjectRBACConfig(ctx, id, rbac any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectRBACConfig", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectRBACConfig), ctx, id, rbac)
}

// UpdateProjectRBACRole mocks base method.
func (m *MockProjectStore) UpdateProjectRBACRole(ctx context.Context, id, name string, policies []*model.ProjectRBACPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectRBACRole", ctx, id, name, policies)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectRBACRole indicates an expected call of UpdateProjectRBACRole.
func (mr *MockProjectStoreMockRecorder) UpdateProjectRBACRole(ctx, id, name, policies any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectRBACRole", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectRBACRole), ctx, id, name, policies)
}

// UpdateProjectSSOConfig mocks base method.
func (m *MockProjectStore) UpdateProjectSSOConfig(ctx context.Context, id string, sso *model.ProjectSSOConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectSSOConfig", ctx, id, sso)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectSSOConfig indicates an expected call of UpdateProjectSSOConfig.
func (mr *MockProjectStoreMockRecorder) UpdateProjectSSOConfig(ctx, id, sso any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectSSOConfig", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectSSOConfig), ctx, id, sso)
}

// UpdateProjectStaticAdmin mocks base method.
func (m *MockProjectStore) UpdateProjectStaticAdmin(ctx context.Context, id, username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectStaticAdmin", ctx, id, username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectStaticAdmin indicates an expected call of UpdateProjectStaticAdmin.
func (mr *MockProjectStoreMockRecorder) UpdateProjectStaticAdmin(ctx, id, username, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectStaticAdmin", reflect.TypeOf((*MockProjectStore)(nil).UpdateProjectStaticAdmin), ctx, id, username, password)
}

// MockPipedStore is a mock of PipedStore interface.
type MockPipedStore struct {
	ctrl     *gomock.Controller
	recorder *MockPipedStoreMockRecorder
	isgomock struct{}
}

// MockPipedStoreMockRecorder is the mock recorder for MockPipedStore.
type MockPipedStoreMockRecorder struct {
	mock *MockPipedStore
}

// NewMockPipedStore creates a new mock instance.
func NewMockPipedStore(ctrl *gomock.Controller) *MockPipedStore {
	mock := &MockPipedStore{ctrl: ctrl}
	mock.recorder = &MockPipedStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipedStore) EXPECT() *MockPipedStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPipedStore) Add(ctx context.Context, piped *model.Piped) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, piped)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPipedStoreMockRecorder) Add(ctx, piped any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPipedStore)(nil).Add), ctx, piped)
}

// AddKey mocks base method.
func (m *MockPipedStore) AddKey(ctx context.Context, id, keyHash, creator string, createdAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKey", ctx, id, keyHash, creator, createdAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddKey indicates an expected call of AddKey.
func (mr *MockPipedStoreMockRecorder) AddKey(ctx, id, keyHash, creator, createdAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKey", reflect.TypeOf((*MockPipedStore)(nil).AddKey), ctx, id, keyHash, creator, createdAt)
}

// DeleteOldKeys mocks base method.
func (m *MockPipedStore) DeleteOldKeys(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOldKeys", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOldKeys indicates an expected call of DeleteOldKeys.
func (mr *MockPipedStoreMockRecorder) DeleteOldKeys(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOldKeys", reflect.TypeOf((*MockPipedStore)(nil).DeleteOldKeys), ctx, id)
}

// DisablePiped mocks base method.
func (m *MockPipedStore) DisablePiped(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisablePiped", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisablePiped indicates an expected call of DisablePiped.
func (mr *MockPipedStoreMockRecorder) DisablePiped(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisablePiped", reflect.TypeOf((*MockPipedStore)(nil).DisablePiped), ctx, id)
}

// EnablePiped mocks base method.
func (m *MockPipedStore) EnablePiped(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePiped", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnablePiped indicates an expected call of EnablePiped.
func (mr *MockPipedStoreMockRecorder) EnablePiped(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePiped", reflect.TypeOf((*MockPipedStore)(nil).EnablePiped), ctx, id)
}

// Get mocks base method.
func (m *MockPipedStore) Get(ctx context.Context, id string) (*model.Piped, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.Piped)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPipedStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPipedStore)(nil).Get), ctx, id)
}

// List mocks base method.
func (m *MockPipedStore) List(ctx context.Context, opts datastore.ListOptions) ([]*model.Piped, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].([]*model.Piped)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPipedStoreMockRecorder) List(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPipedStore)(nil).List), ctx, opts)
}

// UpdateDesiredVersion mocks base method.
func (m *MockPipedStore) UpdateDesiredVersion(ctx context.Context, id, version string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDesiredVersion", ctx, id, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDesiredVersion indicates an expected call of UpdateDesiredVersion.
func (mr *MockPipedStoreMockRecorder) UpdateDesiredVersion(ctx, id, version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDesiredVersion", reflect.TypeOf((*MockPipedStore)(nil).UpdateDesiredVersion), ctx, id, version)
}

// UpdateInfo mocks base method.
func (m *MockPipedStore) UpdateInfo(ctx context.Context, id, name, desc string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfo", ctx, id, name, desc)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfo indicates an expected call of UpdateInfo.
func (mr *MockPipedStoreMockRecorder) UpdateInfo(ctx, id, name, desc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfo", reflect.TypeOf((*MockPipedStore)(nil).UpdateInfo), ctx, id, name, desc)
}

// UpdateMetadata mocks base method.
func (m *MockPipedStore) UpdateMetadata(ctx context.Context, id, version, config string, pps []*model.Piped_PlatformProvider, repos []*model.ApplicationGitRepository, se *model.Piped_SecretEncryption, startedAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetadata", ctx, id, version, config, pps, repos, se, startedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetadata indicates an expected call of UpdateMetadata.
func (mr *MockPipedStoreMockRecorder) UpdateMetadata(ctx, id, version, config, pps, repos, se, startedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetadata", reflect.TypeOf((*MockPipedStore)(nil).UpdateMetadata), ctx, id, version, config, pps, repos, se, startedAt)
}

// MockApplicationStore is a mock of ApplicationStore interface.
type MockApplicationStore struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationStoreMockRecorder
	isgomock struct{}
}

// MockApplicationStoreMockRecorder is the mock recorder for MockApplicationStore.
type MockApplicationStoreMockRecorder struct {
	mock *MockApplicationStore
}

// NewMockApplicationStore creates a new mock instance.
func NewMockApplicationStore(ctrl *gomock.Controller) *MockApplicationStore {
	mock := &MockApplicationStore{ctrl: ctrl}
	mock.recorder = &MockApplicationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationStore) EXPECT() *MockApplicationStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockApplicationStore) Add(ctx context.Context, app *model.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, app)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockApplicationStoreMockRecorder) Add(ctx, app any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockApplicationStore)(nil).Add), ctx, app)
}

// Delete mocks base method.
func (m *MockApplicationStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockApplicationStoreMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApplicationStore)(nil).Delete), ctx, id)
}

// Disable mocks base method.
func (m *MockApplicationStore) Disable(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockApplicationStoreMockRecorder) Disable(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockApplicationStore)(nil).Disable), ctx, id)
}

// Enable mocks base method.
func (m *MockApplicationStore) Enable(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockApplicationStoreMockRecorder) Enable(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockApplicationStore)(nil).Enable), ctx, id)
}

// Get mocks base method.
func (m *MockApplicationStore) Get(ctx context.Context, id string) (*model.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockApplicationStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationStore)(nil).Get), ctx, id)
}

// List mocks base method.
func (m *MockApplicationStore) List(ctx context.Context, opts datastore.ListOptions) ([]*model.Application, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].([]*model.Application)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockApplicationStoreMockRecorder) List(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApplicationStore)(nil).List), ctx, opts)
}

// UpdateBasicInfo mocks base method.
func (m *MockApplicationStore) UpdateBasicInfo(ctx context.Context, id, name, description string, labels map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBasicInfo", ctx, id, name, description, labels)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBasicInfo indicates an expected call of UpdateBasicInfo.
func (mr *MockApplicationStoreMockRecorder) UpdateBasicInfo(ctx, id, name, description, labels any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBasicInfo", reflect.TypeOf((*MockApplicationStore)(nil).UpdateBasicInfo), ctx, id, name, description, labels)
}

// UpdateConfigFilename mocks base method.
func (m *MockApplicationStore) UpdateConfigFilename(ctx context.Context, id, configFilename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigFilename", ctx, id, configFilename)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfigFilename indicates an expected call of UpdateConfigFilename.
func (mr *MockApplicationStoreMockRecorder) UpdateConfigFilename(ctx, id, configFilename any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigFilename", reflect.TypeOf((*MockApplicationStore)(nil).UpdateConfigFilename), ctx, id, configFilename)
}

// UpdateConfiguration mocks base method.
func (m *MockApplicationStore) UpdateConfiguration(ctx context.Context, id, pipedID, platformProvider, configFilename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfiguration", ctx, id, pipedID, platformProvider, configFilename)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfiguration indicates an expected call of UpdateConfiguration.
func (mr *MockApplicationStoreMockRecorder) UpdateConfiguration(ctx, id, pipedID, platformProvider, configFilename any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfiguration", reflect.TypeOf((*MockApplicationStore)(nil).UpdateConfiguration), ctx, id, pipedID, platformProvider, configFilename)
}

// UpdateDeployingStatus mocks base method.
func (m *MockApplicationStore) UpdateDeployingStatus(ctx context.Context, id string, deploying bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeployingStatus", ctx, id, deploying)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeployingStatus indicates an expected call of UpdateDeployingStatus.
func (mr *MockApplicationStoreMockRecorder) UpdateDeployingStatus(ctx, id, deploying any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployingStatus", reflect.TypeOf((*MockApplicationStore)(nil).UpdateDeployingStatus), ctx, id, deploying)
}

// UpdateMostRecentDeployment mocks base method.
func (m *MockApplicationStore) UpdateMostRecentDeployment(ctx context.Context, id string, status model.DeploymentStatus, deployment *model.ApplicationDeploymentReference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMostRecentDeployment", ctx, id, status, deployment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMostRecentDeployment indicates an expected call of UpdateMostRecentDeployment.
func (mr *MockApplicationStoreMockRecorder) UpdateMostRecentDeployment(ctx, id, status, deployment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMostRecentDeployment", reflect.TypeOf((*MockApplicationStore)(nil).UpdateMostRecentDeployment), ctx, id, status, deployment)
}

// UpdatePlatformProvider mocks base method.
func (m *MockApplicationStore) UpdatePlatformProvider(ctx context.Context, id, provider string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlatformProvider", ctx, id, provider)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlatformProvider indicates an expected call of UpdatePlatformProvider.
func (mr *MockApplicationStoreMockRecorder) UpdatePlatformProvider(ctx, id, provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlatformProvider", reflect.TypeOf((*MockApplicationStore)(nil).UpdatePlatformProvider), ctx, id, provider)
}

// UpdateSyncState mocks base method.
func (m *MockApplicationStore) UpdateSyncState(ctx context.Context, id string, syncState *model.ApplicationSyncState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSyncState", ctx, id, syncState)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSyncState indicates an expected call of UpdateSyncState.
func (mr *MockApplicationStoreMockRecorder) UpdateSyncState(ctx, id, syncState any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSyncState", reflect.TypeOf((*MockApplicationStore)(nil).UpdateSyncState), ctx, id, syncState)
}

// MockDeploymentStore is a mock of DeploymentStore interface.
type MockDeploymentStore struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentStoreMockRecorder
	isgomock struct{}
}

// MockDeploymentStoreMockRecorder is the mock recorder for MockDeploymentStore.
type MockDeploymentStoreMockRecorder struct {
	mock *MockDeploymentStore
}

// NewMockDeploymentStore creates a new mock instance.
func NewMockDeploymentStore(ctrl *gomock.Controller) *MockDeploymentStore {
	mock := &MockDeploymentStore{ctrl: ctrl}
	mock.recorder = &MockDeploymentStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentStore) EXPECT() *MockDeploymentStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockDeploymentStore) Add(ctx context.Context, d *model.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, d)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockDeploymentStoreMockRecorder) Add(ctx, d any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDeploymentStore)(nil).Add), ctx, d)
}

// Get mocks base method.
func (m *MockDeploymentStore) Get(ctx context.Context, id string) (*model.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDeploymentStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeploymentStore)(nil).Get), ctx, id)
}

// List mocks base method.
func (m *MockDeploymentStore) List(ctx context.Context, opts datastore.ListOptions) ([]*model.Deployment, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].([]*model.Deployment)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockDeploymentStoreMockRecorder) List(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeploymentStore)(nil).List), ctx, opts)
}

// UpdateMetadata mocks base method.
func (m *MockDeploymentStore) UpdateMetadata(ctx context.Context, id string, metadata map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetadata", ctx, id, metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetadata indicates an expected call of UpdateMetadata.
func (mr *MockDeploymentStoreMockRecorder) UpdateMetadata(ctx, id, metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetadata", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateMetadata), ctx, id, metadata)
}

// UpdatePluginMetadata mocks base method.
func (m *MockDeploymentStore) UpdatePluginMetadata(ctx context.Context, deploymentID, pluginName string, metadata map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePluginMetadata", ctx, deploymentID, pluginName, metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePluginMetadata indicates an expected call of UpdatePluginMetadata.
func (mr *MockDeploymentStoreMockRecorder) UpdatePluginMetadata(ctx, deploymentID, pluginName, metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePluginMetadata", reflect.TypeOf((*MockDeploymentStore)(nil).UpdatePluginMetadata), ctx, deploymentID, pluginName, metadata)
}

// UpdateStageMetadata mocks base method.
func (m *MockDeploymentStore) UpdateStageMetadata(ctx context.Context, deploymentID, stageID string, metadata map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStageMetadata", ctx, deploymentID, stageID, metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStageMetadata indicates an expected call of UpdateStageMetadata.
func (mr *MockDeploymentStoreMockRecorder) UpdateStageMetadata(ctx, deploymentID, stageID, metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStageMetadata", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateStageMetadata), ctx, deploymentID, stageID, metadata)
}

// UpdateStageStatus mocks base method.
func (m *MockDeploymentStore) UpdateStageStatus(ctx context.Context, id, stageID string, status model.StageStatus, reason string, requires []string, visible bool, retriedCount int32, completedAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStageStatus", ctx, id, stageID, status, reason, requires, visible, retriedCount, completedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStageStatus indicates an expected call of UpdateStageStatus.
func (mr *MockDeploymentStoreMockRecorder) UpdateStageStatus(ctx, id, stageID, status, reason, requires, visible, retriedCount, completedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStageStatus", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateStageStatus), ctx, id, stageID, status, reason, requires, visible, retriedCount, completedAt)
}

// UpdateStatus mocks base method.
func (m *MockDeploymentStore) UpdateStatus(ctx context.Context, id string, status model.DeploymentStatus, reason string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, id, status, reason)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockDeploymentStoreMockRecorder) UpdateStatus(ctx, id, status, reason any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateStatus), ctx, id, status, reason)
}

// UpdateToCompleted mocks base method.
func (m *MockDeploymentStore) UpdateToCompleted(ctx context.Context, id string, status model.DeploymentStatus, stageStatuses map[string]model.StageStatus, reason string, completedAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToCompleted", ctx, id, status, stageStatuses, reason, completedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToCompleted indicates an expected call of UpdateToCompleted.
func (mr *MockDeploymentStoreMockRecorder) UpdateToCompleted(ctx, id, status, stageStatuses, reason, completedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToCompleted", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateToCompleted), ctx, id, status, stageStatuses, reason, completedAt)
}

// UpdateToPlanned mocks base method.
func (m *MockDeploymentStore) UpdateToPlanned(ctx context.Context, id, summary, reason, runningCommitHash, runningConfigFilename, version string, versions []*model.ArtifactVersion, stages []*model.PipelineStage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToPlanned", ctx, id, summary, reason, runningCommitHash, runningConfigFilename, version, versions, stages)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToPlanned indicates an expected call of UpdateToPlanned.
func (mr *MockDeploymentStoreMockRecorder) UpdateToPlanned(ctx, id, summary, reason, runningCommitHash, runningConfigFilename, version, versions, stages any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToPlanned", reflect.TypeOf((*MockDeploymentStore)(nil).UpdateToPlanned), ctx, id, summary, reason, runningCommitHash, runningConfigFilename, version, versions, stages)
}

// MockCommandStore is a mock of CommandStore interface.
type MockCommandStore struct {
	ctrl     *gomock.Controller
	recorder *MockCommandStoreMockRecorder
	isgomock struct{}
}

// MockCommandStoreMockRecorder is the mock recorder for MockCommandStore.
type MockCommandStoreMockRecorder struct {
	mock *MockCommandStore
}

// NewMockCommandStore creates a new mock instance.
func NewMockCommandStore(ctrl *gomock.Controller) *MockCommandStore {
	mock := &MockCommandStore{ctrl: ctrl}
	mock.recorder = &MockCommandStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandStore) EXPECT() *MockCommandStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCommandStore) Add(ctx context.Context, cmd *model.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCommandStoreMockRecorder) Add(ctx, cmd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCommandStore)(nil).Add), ctx, cmd)
}

// Get mocks base method.
func (m *MockCommandStore) Get(ctx context.Context, id string) (*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCommandStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCommandStore)(nil).Get), ctx, id)
}

// List mocks base method.
func (m *MockCommandStore) List(ctx context.Context, opts datastore.ListOptions) ([]*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].([]*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCommandStoreMockRecorder) List(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCommandStore)(nil).List), ctx, opts)
}

// UpdateStatus mocks base method.
func (m *MockCommandStore) UpdateStatus(ctx context.Context, id string, status model.CommandStatus, metadata map[string]string, handledAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, id, status, metadata, handledAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockCommandStoreMockRecorder) UpdateStatus(ctx, id, status, metadata, handledAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockCommandStore)(nil).UpdateStatus), ctx, id, status, metadata, handledAt)
}
