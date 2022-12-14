package zookeeper

import (
	"github.com/go-zookeeper/zk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("ZooKeeper", func() {
	When("A ZooKeeper server is online at localhost", func() {
		It(`Should be able to establish connections
		    and create ephemeral nodes with all 
		    permissions from clients`, func() {
			local := "127.0.0.1"
			conn, _, err := zk.Connect([]string{local}, time.Second)
			Expect(err).To(BeNil())

			zkPath := "/zkPath"
			zkFlags := int32(zk.FlagEphemeral)
			data := "zkPathData"
			registeredPath, err := CreateZNode(conn, zkPath, zkFlags, data)
			Expect(err).To(BeNil())
			Expect(registeredPath).To(Equal(zkPath))

			exists := CheckZNode(conn, zkPath)
			Expect(exists).To(Equal(true))

			retrievedData, _ := GetZNode(conn, zkPath)
			Expect(retrievedData).To(Equal(data))
		})

		It(`Should be able to establish connections
			and create persistent nodes with all
			permissions from clients`, func() {
			local := "127.0.0.1"
			conn, _, err := zk.Connect([]string{local}, time.Second)
			Expect(err).To(BeNil())
		
			zkPath := "/zkPersistent"
			zkFlags := int32(0)
			data := "zkPersistentData"
			registeredPath, err := CreateZNode(conn, zkPath, zkFlags, data)
			Expect(err).To(BeNil())
			Expect(registeredPath).To(Equal(zkPath))
		})

		It(`Should be able to establish connections
			and retrieve data from persistent nodes`, func() {
			local := "127.0.0.1"
			conn, _, err := zk.Connect([]string{local}, time.Second)
			Expect(err).To(BeNil())

			zkPath := "/zkPersistent"
			data := "zkPersistentData"
			exists := CheckZNode(conn, zkPath)
			Expect(exists).To(Equal(true))

			retrievedData, _ := GetZNode(conn, zkPath)
			Expect(retrievedData).To(Equal(data))
		})
	})
})