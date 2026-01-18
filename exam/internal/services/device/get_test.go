package device

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"exam/internal/domain"
	"exam/internal/repository"
	repoDevice "exam/internal/repository/device"
	repoEntity "exam/internal/repository/device/entity"
	"exam/internal/repository/echographic"
	"exam/internal/repository/image"
	"exam/internal/repository/kt"
	"exam/internal/repository/mri"
	"exam/internal/repository/node"
	"exam/internal/repository/segment"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
)

type daoMock struct {
	daolib.DAO
	repo repoDevice.Repository
}

func (d *daoMock) NewDeviceQuery(ctx context.Context) repoDevice.Repository {
	return d.repo
}

// SS3
func (d *daoMock) NewFileRepo() repository.FileRepo {
	return nil
}

func (d *daoMock) NewMriQuery(ctx context.Context) mri.Repository {
	return nil
}

func (d *daoMock) NewKtQuery(ctx context.Context) kt.Repository {
	return nil
}

func (d *daoMock) NewImageQuery(ctx context.Context) image.Repository {
	return nil
}

func (d *daoMock) NewSegmentQuery(ctx context.Context) segment.Repository {
	return nil
}

func (d *daoMock) NewNodeQuery(ctx context.Context) node.Repository {
	return nil
}

func (d *daoMock) NewEchographicQuery(ctx context.Context) echographic.Repository {
	return nil
}

type deviceRepoMock struct {
	called  bool
	devices []repoEntity.Device
	err     error
}

func (r *deviceRepoMock) GetDeviceList() ([]repoEntity.Device, error) {
	r.called = true
	return r.devices, r.err
}

func (r *deviceRepoMock) CreateDevice(name string) (int, error) {
	return 0, nil
}

// --- Tests ---

func TestService_GetDeviceList_OK(t *testing.T) {
	t.Parallel()

	// Arrange
	repo := &deviceRepoMock{
		devices: []repoEntity.Device{
			{Id: 1, Name: "CT"},
			{Id: 2, Name: "MRI"},
		},
	}
	var dao repository.DAO = &daoMock{repo: repo}

	s := New(dao)

	// Act
	got, err := s.GetDeviceList(context.Background())

	// Assert
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !repo.called {
		t.Fatalf("expected GetDeviceList to be called")
	}

	want := []domain.Device{
		{Id: 1, Name: "CT"},
		{Id: 2, Name: "MRI"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected result.\nwant: %#v\ngot:  %#v", want, got)
	}
}

func TestService_GetDeviceList_RepoError(t *testing.T) {
	t.Parallel()

	// Arrange
	wantErr := errors.New("db error")
	repo := &deviceRepoMock{err: wantErr}
	var dao repository.DAO = &daoMock{repo: repo}

	s := New(dao)

	// Act
	got, err := s.GetDeviceList(context.Background())

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}
	if got != nil {
		t.Fatalf("expected nil slice on error, got %#v", got)
	}
	if !repo.called {
		t.Fatalf("expected GetDeviceList to be called")
	}
}

func TestService_GetDeviceList_Empty(t *testing.T) {
	t.Parallel()

	// Arrange
	repo := &deviceRepoMock{devices: []repoEntity.Device{}}
	var dao repository.DAO = &daoMock{repo: repo}

	s := New(dao)

	// Act
	got, err := s.GetDeviceList(context.Background())

	// Assert
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got == nil {
		// If your mapper returns nil for empty input and you accept that,
		// change this assertion accordingly.
		t.Fatalf("expected empty slice, got nil")
	}
	if len(got) != 0 {
		t.Fatalf("expected 0 devices, got %d", len(got))
	}
	if !repo.called {
		t.Fatalf("expected GetDeviceList to be called")
	}
}
