package metadata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAction(t *testing.T) {

	Convey("Test Default Action", t, func() {
		a := NewAction()

		So(len(a.GetFields()), ShouldEqual, 0)
	})

	Convey("Test New Field", t, func() {
		a := NewAction()

		name := "fieldname"

		So(a.HasField(name), ShouldBeFalse)

		a.Field(name)
		So(a.HasField(name), ShouldBeTrue)

		So(func() { a.Field() }, ShouldPanic)

		result := a.Field("response", "user")
		So(a.Field("response").Field("user"), ShouldEqual, result)

	})

	Convey("Test Has Field", t, func() {
		a := NewAction()

		So(func() { a.HasField() }, ShouldPanic)

		a.Field("response").Field("user")

		So(a.HasField("response", "user"), ShouldBeTrue)

	})

	Convey("Test Description", t, func() {
		a := NewAction()

		So(a.GetDescription(), ShouldEqual, "")
		description := "fieldname"
		So(a.Description(description).GetDescription(), ShouldEqual, description)

	})

	Convey("Test Action From", t, func() {
		a := NewAction()

		So(func() { a.From([]string{}) }, ShouldPanic)

		type TestStruct struct {
			A string `json:"a"`
			B string `json:"b"`
		}

		a.From(TestStruct{})

		So(a.HasField("a"), ShouldBeTrue)

		a.From(&TestStruct{})

		So(a.HasField("a"), ShouldBeTrue)

	})

	Convey("Test GetData", t, func() {
		a := NewAction()

		So(a.GetData(), ShouldNotBeNil)

		name := "fieldname"
		a.Field(name)

		So(len(a.GetData()), ShouldEqual, 1)

	})

}
