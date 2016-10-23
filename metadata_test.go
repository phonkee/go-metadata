package metadata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetadata(t *testing.T) {

	Convey("Test Default Metadata", t, func() {
		md := New()
		So(len(md.(*metadata).actions), ShouldEqual, 0)
	})

	Convey("Test Add Action", t, func() {
		md := New()
		So(len(md.(*metadata).actions), ShouldEqual, 0)

		action := md.Action("GET")
		So(len(md.(*metadata).actions), ShouldEqual, 1)

		action2 := md.Action("GET")
		So(len(md.(*metadata).actions), ShouldEqual, 1)

		So(action, ShouldEqual, action2)
	})

	Convey("Test Name get/set", t, func() {
		md := New()

		So(md.GetName(), ShouldEqual, "")
		name := "some name"

		md.Name(name)
		So(md.GetName(), ShouldEqual, name)
	})

	Convey("Test Description get/set", t, func() {
		md := New()

		So(md.GetDescription(), ShouldEqual, "")
		description := "description"

		md.Description(description)
		So(md.GetDescription(), ShouldEqual, description)
	})

	Convey("Test Action aliases", t, func() {
		md := New()

		actioncreate := md.ActionCreate()
		actioncreate2 := md.ActionCreate()
		So(actioncreate, ShouldEqual, actioncreate2)

		actionupdate := md.ActionUpdate()
		actionupdate2 := md.ActionUpdate()
		So(actionupdate, ShouldEqual, actionupdate2)

		actionretrieve := md.ActionRetrieve()
		actionretrieve2 := md.ActionRetrieve()
		So(actionretrieve, ShouldEqual, actionretrieve2)

		actiondelete := md.ActionDelete()
		actiondelete2 := md.ActionDelete()
		So(actiondelete, ShouldEqual, actiondelete2)

	})

	Convey("Test GetData/MarshalJSON", t, func() {

		name := "mdname"
		description := "mddescription"
		md := New().Name(name).Description(description)
		md.ActionCreate()

		data := md.GetData()
		So(data["name"], ShouldEqual, name)
		So(data["description"], ShouldEqual, description)

		_, err := md.MarshalJSON()
		So(err, ShouldBeNil)

	})

}
