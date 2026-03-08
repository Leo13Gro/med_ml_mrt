//go:build e2e

package patient_test

import (
	"time"

	pb "med/internal/generated/grpc/service"
	"med/tests/e2e/flow"

	"github.com/AlekSi/pointer"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestUpdatePatient_Success() {
	data, err := flow.New(suite.deps, flow.CreatePatient).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	active := gofakeit.Bool()
	malignancy := gofakeit.Bool()
	lastExamDate := gofakeit.Date().Format(time.RFC3339)

	updateResp, err := suite.deps.Adapter.UpdatePatient(
		suite.T().Context(),
		&pb.UpdatePatientIn{
			Id:           data.Patient.Id.String(),
			Active:       pointer.To(active),
			Malignancy:   pointer.To(malignancy),
			LastExamDate: pointer.ToString(lastExamDate),
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Patient.Id.String(), updateResp.Patient.Id)
	require.Equal(suite.T(), active, updateResp.Patient.Active)
	require.Equal(suite.T(), malignancy, updateResp.Patient.Malignancy)
	require.Equal(suite.T(), lastExamDate, *updateResp.Patient.LastExamDate)
}
