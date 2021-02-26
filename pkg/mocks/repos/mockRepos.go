// Code generated by MockGen. DO NOT EDIT.
// Source: repos.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"
	time "time"

	tarconsumer "github.com/fidelity/kraan/pkg/internal/tarconsumer"
	repos "github.com/fidelity/kraan/pkg/repos"
	v1beta1 "github.com/fluxcd/source-controller/api/v1beta1"
	gomock "github.com/golang/mock/gomock"
)

// MockRepos is a mock of Repos interface.
type MockRepos struct {
	ctrl     *gomock.Controller
	recorder *MockReposMockRecorder
}

// MockReposMockRecorder is the mock recorder for MockRepos.
type MockReposMockRecorder struct {
	mock *MockRepos
}

// NewMockRepos creates a new mock instance.
func NewMockRepos(ctrl *gomock.Controller) *MockRepos {
	mock := &MockRepos{ctrl: ctrl}
	mock.recorder = &MockReposMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepos) EXPECT() *MockReposMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRepos) Add(srcRepo *v1beta1.GitRepository) repos.Repo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", srcRepo)
	ret0, _ := ret[0].(repos.Repo)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockReposMockRecorder) Add(srcRepo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepos)(nil).Add), srcRepo)
}

// Delete mocks base method.
func (m *MockRepos) Delete(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", name)
}

// Delete indicates an expected call of Delete.
func (mr *MockReposMockRecorder) Delete(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepos)(nil).Delete), name)
}

// Get mocks base method.
func (m *MockRepos) Get(name string) repos.Repo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(repos.Repo)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockReposMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepos)(nil).Get), name)
}

// GetRootPath mocks base method.
func (m *MockRepos) GetRootPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRootPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRootPath indicates an expected call of GetRootPath.
func (mr *MockReposMockRecorder) GetRootPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRootPath", reflect.TypeOf((*MockRepos)(nil).GetRootPath))
}

// List mocks base method.
func (m *MockRepos) List() map[string]repos.Repo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(map[string]repos.Repo)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockReposMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepos)(nil).List))
}

// SetHTTPClient mocks base method.
func (m *MockRepos) SetHTTPClient(client *http.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHTTPClient", client)
}

// SetHTTPClient indicates an expected call of SetHTTPClient.
func (mr *MockReposMockRecorder) SetHTTPClient(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHTTPClient", reflect.TypeOf((*MockRepos)(nil).SetHTTPClient), client)
}

// SetHostName mocks base method.
func (m *MockRepos) SetHostName(hostName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHostName", hostName)
}

// SetHostName indicates an expected call of SetHostName.
func (mr *MockReposMockRecorder) SetHostName(hostName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHostName", reflect.TypeOf((*MockRepos)(nil).SetHostName), hostName)
}

// SetRootPath mocks base method.
func (m *MockRepos) SetRootPath(path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRootPath", path)
}

// SetRootPath indicates an expected call of SetRootPath.
func (mr *MockReposMockRecorder) SetRootPath(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRootPath", reflect.TypeOf((*MockRepos)(nil).SetRootPath), path)
}

// SetTimeOut mocks base method.
func (m *MockRepos) SetTimeOut(timeOut time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimeOut", timeOut)
}

// SetTimeOut indicates an expected call of SetTimeOut.
func (mr *MockReposMockRecorder) SetTimeOut(timeOut interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeOut", reflect.TypeOf((*MockRepos)(nil).SetTimeOut), timeOut)
}

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockRepo) AddUser(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUser", name)
}

// AddUser indicates an expected call of AddUser.
func (mr *MockRepoMockRecorder) AddUser(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockRepo)(nil).AddUser), name)
}

// GetDataPath mocks base method.
func (m *MockRepo) GetDataPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDataPath indicates an expected call of GetDataPath.
func (mr *MockRepoMockRecorder) GetDataPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataPath", reflect.TypeOf((*MockRepo)(nil).GetDataPath))
}

// GetGitRepo mocks base method.
func (m *MockRepo) GetGitRepo() *v1beta1.GitRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitRepo")
	ret0, _ := ret[0].(*v1beta1.GitRepository)
	return ret0
}

// GetGitRepo indicates an expected call of GetGitRepo.
func (mr *MockRepoMockRecorder) GetGitRepo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitRepo", reflect.TypeOf((*MockRepo)(nil).GetGitRepo))
}

// GetLoadPath mocks base method.
func (m *MockRepo) GetLoadPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLoadPath indicates an expected call of GetLoadPath.
func (mr *MockRepoMockRecorder) GetLoadPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadPath", reflect.TypeOf((*MockRepo)(nil).GetLoadPath))
}

// GetPath mocks base method.
func (m *MockRepo) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockRepoMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockRepo)(nil).GetPath))
}

// GetSourceName mocks base method.
func (m *MockRepo) GetSourceName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceName indicates an expected call of GetSourceName.
func (mr *MockRepoMockRecorder) GetSourceName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceName", reflect.TypeOf((*MockRepo)(nil).GetSourceName))
}

// GetSourceNameSpace mocks base method.
func (m *MockRepo) GetSourceNameSpace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceNameSpace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceNameSpace indicates an expected call of GetSourceNameSpace.
func (mr *MockRepoMockRecorder) GetSourceNameSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceNameSpace", reflect.TypeOf((*MockRepo)(nil).GetSourceNameSpace))
}

// IsSynced mocks base method.
func (m *MockRepo) IsSynced() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSynced")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSynced indicates an expected call of IsSynced.
func (mr *MockRepoMockRecorder) IsSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSynced", reflect.TypeOf((*MockRepo)(nil).IsSynced))
}

// LinkData mocks base method.
func (m *MockRepo) LinkData(layerPath, sourcePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkData", layerPath, sourcePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkData indicates an expected call of LinkData.
func (mr *MockRepoMockRecorder) LinkData(layerPath, sourcePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkData", reflect.TypeOf((*MockRepo)(nil).LinkData), layerPath, sourcePath)
}

// RemoveUser mocks base method.
func (m *MockRepo) RemoveUser(namer string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", namer)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockRepoMockRecorder) RemoveUser(namer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockRepo)(nil).RemoveUser), namer)
}

// SetGitRepo mocks base method.
func (m *MockRepo) SetGitRepo(src *v1beta1.GitRepository, rootPath string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGitRepo", src, rootPath)
}

// SetGitRepo indicates an expected call of SetGitRepo.
func (mr *MockRepoMockRecorder) SetGitRepo(src, rootPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGitRepo", reflect.TypeOf((*MockRepo)(nil).SetGitRepo), src, rootPath)
}

// SetHTTPClient mocks base method.
func (m *MockRepo) SetHTTPClient(client *http.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHTTPClient", client)
}

// SetHTTPClient indicates an expected call of SetHTTPClient.
func (mr *MockRepoMockRecorder) SetHTTPClient(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHTTPClient", reflect.TypeOf((*MockRepo)(nil).SetHTTPClient), client)
}

// SetHostName mocks base method.
func (m *MockRepo) SetHostName(hostName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHostName", hostName)
}

// SetHostName indicates an expected call of SetHostName.
func (mr *MockRepoMockRecorder) SetHostName(hostName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHostName", reflect.TypeOf((*MockRepo)(nil).SetHostName), hostName)
}

// SetTarConsumer mocks base method.
func (m *MockRepo) SetTarConsumer(tarConsumer tarconsumer.TarConsumer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTarConsumer", tarConsumer)
}

// SetTarConsumer indicates an expected call of SetTarConsumer.
func (mr *MockRepoMockRecorder) SetTarConsumer(tarConsumer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTarConsumer", reflect.TypeOf((*MockRepo)(nil).SetTarConsumer), tarConsumer)
}

// SyncRepo mocks base method.
func (m *MockRepo) SyncRepo() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncRepo")
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncRepo indicates an expected call of SyncRepo.
func (mr *MockRepoMockRecorder) SyncRepo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncRepo", reflect.TypeOf((*MockRepo)(nil).SyncRepo))
}

// TidyAll mocks base method.
func (m *MockRepo) TidyAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TidyAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// TidyAll indicates an expected call of TidyAll.
func (mr *MockRepoMockRecorder) TidyAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TidyAll", reflect.TypeOf((*MockRepo)(nil).TidyAll))
}

// TidyRepo mocks base method.
func (m *MockRepo) TidyRepo() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TidyRepo")
	ret0, _ := ret[0].(error)
	return ret0
}

// TidyRepo indicates an expected call of TidyRepo.
func (mr *MockRepoMockRecorder) TidyRepo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TidyRepo", reflect.TypeOf((*MockRepo)(nil).TidyRepo))
}

// fetchArtifact mocks base method.
func (m *MockRepo) fetchArtifact(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetchArtifact", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// fetchArtifact indicates an expected call of fetchArtifact.
func (mr *MockRepoMockRecorder) fetchArtifact(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetchArtifact", reflect.TypeOf((*MockRepo)(nil).fetchArtifact), ctx)
}
