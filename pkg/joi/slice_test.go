package joi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/softbrewery/gojoi/pkg/joi"
	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Slice", func() {

	Describe("Slice", func() {

		It("Should create a new schema", func() {
			s := Slice()
			Expect(s).ToNot(BeNil())
		})

		It("Should pass if data type is slice", func() {
			s := Slice()
			Expect(Validate([]string{}, s)).To(BeNil())
		})

		It("Should pass if data type is *slice", func() {
			s := Slice()
			data := []string{}
			Expect(Validate(&data, s)).To(BeNil())
		})

		It("Should fail if Any property fails", func() {
			s := Slice().Required()

			Expect(Validate(nil, s)).To(Equal(ErrAnyRequired))
		})

		It("Should fail if data type is string", func() {
			s := Slice()
			Expect(Validate("hello", s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is int", func() {
			s := Slice()
			Expect(Validate(100, s)).To(Equal(ErrAnyType))
		})

		It("Should fail if data type is bool", func() {
			s := Slice()
			Expect(Validate(true, s)).To(Equal(ErrAnyType))
		})
	})

	Describe("Kind", func() {

		It("Should return slice", func() {
			s := Slice()
			Expect(s.Kind()).To(Equal("slice"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Slice()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Min", func() {

		data := []string{"hello", "world"}

		It("Error should be nil if slice is bigger than", func() {
			s := Slice().Min(1)
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be nil if slice is bigger than *slice", func() {
			s := Slice().Min(1)
			Expect(Validate(&data, s)).To(BeNil())
		})

		It("Error should be nil if slice is equal", func() {
			s := Slice().Min(2)
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be nil if slice is equal *", func() {
			s := Slice().Min(2)
			Expect(Validate(&data, s)).To(BeNil())
		})

		It("Error should be not nil if slice is smaller than", func() {
			s := Slice().Min(3)
			Expect(Validate(data, s)).To(Equal(ErrSliceMin))
		})

		It("Error should be not nil if slice is smaller than *", func() {
			s := Slice().Min(3)
			Expect(Validate(&data, s)).To(Equal(ErrSliceMin))
		})
	})

	Describe("Max", func() {

		data := []string{"hello", "world"}

		It("Error should be nil if slice is smaller than", func() {
			s := Slice().Max(3)
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be nil if slice is equal", func() {
			s := Slice().Max(2)
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be not nil if slice is bigger than", func() {
			s := Slice().Max(1)
			Expect(Validate(data, s)).To(Equal(ErrSliceMax))
		})
	})

	Describe("Length", func() {

		data := []string{"hello", "world"}

		It("Error should be not nil if value is smaller than", func() {
			s := Slice().Length(1)
			Expect(Validate(data, s)).To(Equal(ErrSliceLength))
		})

		It("Error should be nil if value is equal", func() {
			s := Slice().Length(2)
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be not nil if value is bigger than", func() {
			s := Slice().Length(3)
			Expect(Validate(data, s)).To(Equal(ErrSliceLength))
		})
	})

	Describe("Items", func() {

		data := []string{"hello", "world"}

		It("Error should be nil if items are matching Any", func() {
			s := Slice().Items(
				joi.Any(),
			)
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be nil if items are matching String", func() {
			s := Slice().Items(
				joi.String(),
			)
			Expect(Validate(data, s)).To(BeNil())
		})

		It("Error should be nil if items are matching *String", func() {
			pData := []*string{}
			str1 := "hello"
			str2 := "world"
			pData = append(pData, &str1, &str2)
			s := Slice().Items(
				joi.String(),
			)
			Expect(Validate(&pData, s)).To(BeNil())
		})

		It("Error should be not nil if items are not matching", func() {
			s := Slice().Items(
				joi.Slice(),
			)
			Expect(Validate(data, s)).To(Equal(ErrAnyType))
		})
	})
})
