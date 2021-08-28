package _interface_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/adolsalamanca/grpc-ports/client/internal/domain/entity"
	"github.com/adolsalamanca/grpc-ports/client/internal/domain/repository/mocks"
	_interface "github.com/adolsalamanca/grpc-ports/client/internal/interface"
	"github.com/adolsalamanca/grpc-ports/client/pkg"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate mockgen -source ../domain/repository/port.go -destination ../domain/repository/mocks/port_repository.go -package=mocks

const (
	PortsBody = `{"AEAJM":{"name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"},"AEAUH":{"name":"Abu Dhabi","coordinates":[54.37,24.47],"city":"Abu Dhabi","province":"Abu Z¸aby [Abu Dhabi]","country":"United Arab Emirates","alias":[],"regions":[],"timezone":"Asia/Dubai","unlocs":["AEAUH"],"code":"52001"}}`
	FakeError = pkg.Error("Fake error indicating something did not ")
)

var _ = Describe("Handler test suite", func() {

	var portHandler *_interface.PortHandler
	var ctrl *gomock.Controller
	var portRepository *mocks.MockPortRepository

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		portRepository = mocks.NewMockPortRepository(ctrl)
		portHandler = _interface.NewPortHandler(portRepository)
	})

	AfterEach(func() {
		defer ctrl.Finish()
	})

	Context("get ports", func() {

		It("should retrieve not found err if ports were not found", func() {
			portRepository.EXPECT().GetAllPorts().Return(nil, FakeError)

			rec, c := newEchoContext(httptest.NewRequest(http.MethodGet, "/", nil))
			portHandler.GetPorts(c)

			Expect(rec.Code).To(BeEquivalentTo(http.StatusNotFound))
		})

		It("should return all ports if there are any", func() {

			ports := []entity.PortInfo{
				{
					Name:    "PortA",
					City:    "City",
					Country: "Countr",
				},
			}
			portRepository.EXPECT().GetAllPorts().Return(ports, nil)

			rec, c := newEchoContext(httptest.NewRequest(http.MethodGet, "/", nil))
			portHandler.GetPorts(c)

			Expect(rec.Code).To(BeEquivalentTo(http.StatusOK))
		})

	})

	Context("store ports", func() {

		It("should return internal server error if there was an error storing ports", func() {
			portRepository.EXPECT().StorePorts(gomock.Any()).Return(FakeError)

			rec, c := newEchoContext(httptest.NewRequest(http.MethodPost, "/", strings.NewReader(PortsBody)))
			portHandler.StorePorts(c)

			Expect(rec.Code).To(BeEquivalentTo(http.StatusInternalServerError))
		})

		It("should respond properly after storing ports", func() {
			portRepository.EXPECT().StorePorts(gomock.Any()).Return(nil)

			rec, c := newEchoContext(httptest.NewRequest(http.MethodPost, "/", strings.NewReader(PortsBody)))
			portHandler.StorePorts(c)

			Expect(rec.Code).To(BeEquivalentTo(http.StatusOK))
		})

	})
})

func newEchoContext(req *http.Request) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}
