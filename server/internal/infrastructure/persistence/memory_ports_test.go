package persistence_test

import (
	"github.com/adolsalamanca/grpc-ports/server/internal/domain/entity"
	"github.com/adolsalamanca/grpc-ports/server/internal/infrastructure/persistence"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ItemMemoryRepository test", func() {

	var p entity.PortInfo
	var r *persistence.PortsMemoryRepository

	BeforeEach(func() {
		r = persistence.NewPortsMemoryRepository()

		p = entity.PortInfo{
			Unlocs: []string{"KeyForPortX"},
			Name:   "PortX",
		}
	})

	When("having an empty repository", func() {

		It("should return EmptyRepository error if get items is called", func() {
			_, err := r.GetPorts()
			Expect(err).To(BeEquivalentTo(persistence.EmptyRepositoryErr))
		})

		It("should not return error if store is called", func() {
			err := r.StorePorts([]entity.PortInfo{
				p,
			})
			Expect(err).ShouldNot(HaveOccurred())

			items, err := r.GetPorts()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(items).ToNot(BeNil())
		})

	})

	When("having a non empty repository", func() {

		var p2, p3 entity.PortInfo

		BeforeEach(func() {
			p2 = entity.PortInfo{
				Unlocs: []string{"KeyForPortY"},
				Name:   "PortY",
			}
			p3 = entity.PortInfo{
				Unlocs: []string{"KeyForPortZ"},
				Name:   "PortZ"}

			err := r.StorePorts([]entity.PortInfo{p2, p3})
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should return all ports if get items is called", func() {
			i, err := r.GetPorts()

			Expect(len(i)).To(BeEquivalentTo(2))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should not return error if store is called", func() {
			err := r.StorePorts([]entity.PortInfo{p})
			Expect(err).ShouldNot(HaveOccurred())

			items, err := r.GetPorts()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(items).ToNot(BeNil())
		})

	})

})
