package xivnet_test

import (
	"encoding/json"

	"github.com/ff14wed/xivnet"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Types", func() {
	Describe("Header", func() {
		It("marshals to JSON correctly", func() {
			bytes, err := json.Marshal(&expectedZlibFrame.Header)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(jsonZlibFrameHeader))
		})
		It("unmarshals from JSON correctly", func() {
			var b xivnet.Header
			err := json.Unmarshal([]byte(jsonZlibFrameHeader), &b)
			Expect(err).ToNot(HaveOccurred())
			Expect(b).To(Equal(expectedZlibFrame.Header))
		})
		It("errors when the input is invalid hexadecimal", func() {
			var b xivnet.Header
			err := json.Unmarshal([]byte(`"XX XX"`), &b)
			Expect(err).To(HaveOccurred())
		})
	})
	Describe("BlockHeader", func() {
		It("marshals to JSON correctly", func() {
			bytes, err := json.Marshal(&expectedZlibFrame.Blocks[0].Header)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(jsonZlibBlock0Header))
		})
		It("unmarshals from JSON correctly", func() {
			var b xivnet.BlockHeader
			err := json.Unmarshal([]byte(jsonZlibBlock0Header), &b)
			Expect(err).ToNot(HaveOccurred())
			Expect(b).To(Equal(expectedZlibFrame.Blocks[0].Header))
		})
		It("errors when the input is invalid hexadecimal", func() {
			var b xivnet.BlockHeader
			err := json.Unmarshal([]byte(`"XX XX"`), &b)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GenericGenericBlockData", func() {
		It("marshals to JSON correctly", func() {
			bytes, err := json.Marshal(&expectedZlibFrame.Blocks[0].Data)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(jsonZlibBlock0Data))
		})
		It("marshals the empty block data correctly", func() {
			var b xivnet.GenericBlockData
			bytes, err := json.Marshal(&b)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(`""`))
		})
		It("marshals the single block data correctly", func() {
			var b xivnet.GenericBlockData
			bytes, err := json.Marshal(&b)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(bytes)).To(Equal(`""`))
		})
		It("unmarshals from JSON correctly", func() {
			var b xivnet.GenericBlockData
			err := json.Unmarshal([]byte(jsonZlibBlock0Data), &b)
			Expect(err).ToNot(HaveOccurred())
			expectedBlockData := expectedZlibFrame.Blocks[0].Data.(*xivnet.GenericBlockData)
			Expect(b).To(Equal(*expectedBlockData))
		})
		It("unmarshals the empty block data correctly", func() {
			var b xivnet.GenericBlockData
			err := json.Unmarshal([]byte(`""`), &b)
			Expect(err).ToNot(HaveOccurred())
			Expect(b).To(BeEmpty())
		})
		It("errors when the input is empty", func() {
			var b xivnet.GenericBlockData
			err := json.Unmarshal([]byte{}, &b)
			Expect(err).To(HaveOccurred())
			Expect(b).To(BeEmpty())
		})
		It("errors when the input is invalid JSON", func() {
			var b xivnet.GenericBlockData
			err := json.Unmarshal([]byte(`"00 00`), &b)
			Expect(err).To(HaveOccurred())
			Expect(b).To(BeEmpty())
		})
		It("errors when the input is invalid hexadecimal", func() {
			var b xivnet.GenericBlockData
			err := json.Unmarshal([]byte(`"XX XX"`), &b)
			Expect(err).To(HaveOccurred())
			Expect(b).To(BeEmpty())
		})
	})

	Describe("Frame", func() {
		It("marshals to JSON correctly", func() {
			bytes, err := json.Marshal(&expectedZlibFrame)
			Expect(err).ToNot(HaveOccurred())
			Expect(bytes).To(ContainSubstring(`"Header":` + jsonZlibFrameHeader))
			Expect(bytes).To(ContainSubstring(`"Length":148`))
			Expect(bytes).To(ContainSubstring(`"Header":` + jsonZlibBlock0Header))
			Expect(bytes).To(ContainSubstring(`"Data":` + jsonZlibBlock0Data))
		})
	})
})