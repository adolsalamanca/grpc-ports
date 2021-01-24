package server_test

import (
	"context"
	"github.com/adolsalamanca/ports/server/infrastructure/api"
	"github.com/adolsalamanca/ports/server/domain/entity"
	"github.com/adolsalamanca/ports/server/domain/repository/mocks"
	"github.com/adolsalamanca/ports/server/infrastructure/common"
	"github.com/adolsalamanca/ports/server/infrastructure/persistence"
	"github.com/adolsalamanca/ports/server/interface/server"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	FakeError = common.Error("Fake error storing item")
)

var _ = Describe("Server handler test suite", func() {

	var ctrl *gomock.Controller
	var portRepository *mocks.MockPortRepository
	var s *server.Server
	var ctx context.Context
	var params api.EmptyParams

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		portRepository = mocks.NewMockPortRepository(ctrl)
		s = server.NewServer(nil, portRepository)
		params = api.EmptyParams{}
	})

	Context("error cases", func() {
		BeforeEach(func() {
			portRepository.EXPECT().GetPorts().Return(nil, persistence.EmptyRepositoryErr)
		})

		It("should return error if GetAllPorts was called", func() {

			r, err := s.GetAllPorts(ctx, &params)

			Expect(r).To(BeNil())
			Expect(err).To(BeEquivalentTo(server.NotFoundError))
		})

		It("should return error if StorePorts failed", func() {
			ports := api.MultiplePortsInput{
				Ports: []*api.Port{
					{
						Name:        "PortX",
						Unlocs:      []string{"KeyX"},
					},
				},
			}
			portRepository.EXPECT().StorePorts(gomock.Any()).Return(FakeError)
			r, err := s.StorePorts(ctx, &ports)

			Expect(r).To(BeNil())
			Expect(err).To(BeEquivalentTo(server.InternalServerError))
		})

	})

	Context("happy path cases", func() {
		var ports []entity.PortInfo
		var province, timezone, code string

		BeforeEach(func() {
			province = "Province"
			timezone = "Timezone"
			code = "Code"
			ports = []entity.PortInfo{
				{
					Name:     "PortY",
					Unlocs: []string{"KeyY"},
					Province: &province,
					Timezone: &timezone,
					Code: &code,
				},
			}
			portRepository.EXPECT().GetPorts().Return(ports, nil)
		})

		It("should return all the elements if GetAllPorts was properly executed", func() {
			r, err := s.GetAllPorts(ctx, &params)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(r.Ports)).To(BeEquivalentTo(1))
		})

		It("should return success if StorePorts was properly executed", func() {
			ports := api.MultiplePortsInput{
				Ports: []*api.Port{
					{
						Name:        "PortZ",
						Unlocs:      []string{"KeyZ"},
					},
				},
			}

			portRepository.EXPECT().StorePorts(gomock.Any()).Return(nil)
			r, err := s.StorePorts(ctx, &ports)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(r.Message).To(BeEquivalentTo(server.SuccessCreateMessage))
		})

	})

})
