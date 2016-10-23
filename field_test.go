package metadata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestField(t *testing.T) {

	Convey("Test Default Field", t, func() {
		f := newField()
		So(f, ShouldNotBeNil)
	})

	Convey("Test Label", t, func() {
		label := "testlabel"
		f := newField().Label(label)
		So(f.GetLabel(), ShouldEqual, label)
	})

	Convey("Test Description", t, func() {
		descrption := "testht"
		f := newField().Description(descrption)
		So(f.GetDescription(), ShouldEqual, descrption)
	})

	Convey("Test Add Field", t, func() {
		f := newField()
		name := "testfield"
		sub := f.Field(name)
		sub2 := f.Field(name)
		So(sub, ShouldEqual, sub2)

		other := "other"
		otherfield := newField()
		f.AddField(other, otherfield)

		So(f.Field(other), ShouldEqual, otherfield)

		So(func() { f.Field() }, ShouldPanic)

		resultuser := f.Field("result", "user")

		So(f.Field("result").Field("user"), ShouldEqual, resultuser)
	})

	Convey("Test Has Field", t, func() {
		f := newField()
		name := "testfield"

		So(func() { f.HasField() }, ShouldPanic)

		So(f.HasField(name), ShouldBeFalse)
		f.Field(name)
		So(f.HasField(name), ShouldBeTrue)

		f.Field("one", "two", "three")
		So(f.HasField("one", "two", "three"), ShouldBeTrue)
	})

	Convey("Test Fields", t, func() {
		f := newField()

		So(len(f.GetFields()), ShouldEqual, 0)

		name := "some field"
		f.Field(name)

		So(len(f.GetFields()), ShouldEqual, 1)

		f.Field(name)
		So(len(f.GetFields()), ShouldEqual, 1)

		f.RemoveField(name)
		So(len(f.GetFields()), ShouldEqual, 0)
	})

	Convey("Test GetData", t, func() {

		label := "lllabel"
		description := "dddesc"

		f := newField().Label(label).Description(description)

		data := f.GetData()

		So(data["label"], ShouldEqual, label)
		So(data["fields"], ShouldBeNil)

		f.Field("subfield")

		data = f.GetData()

		So(len(data["fields"].(map[string]Field)), ShouldEqual, 1)

		So(data["choices"], ShouldBeNil)

		f.Choices().Add("value", "display")
		data = f.GetData()

		So(data["choices"], ShouldHaveSameTypeAs, newChoices())

	})

	Convey("Test RemoveField", t, func() {
		f := newField()
		name := "testfield"
		So(f.HasField(name), ShouldBeFalse)
		f.Field(name)
		So(f.HasField(name), ShouldBeTrue)
		f.RemoveField(name)
		So(f.HasField(name), ShouldBeFalse)
	})

	Convey("Test Required", t, func() {
		f := newField()
		So(f.IsRequired(), ShouldBeFalse)

		f.Required(true)
		So(f.IsRequired(), ShouldBeTrue)
	})

	Convey("Test Type", t, func() {
		f := newField()
		So(f.GetType(), ShouldEqual, "")

		typ := "sometype"

		f.Type(typ)

		So(f.GetType(), ShouldEqual, typ)

	})

	Convey("Test Field.From", t, func() {
		f := newField()

		type TestStruct struct {
			First string `json:"first"`
		}

		f.From(TestStruct{})

		So(f.HasField("first"), ShouldBeTrue)
		So(f.IsRequired(), ShouldBeTrue)

		f.From(&TestStruct{})
		So(f.HasField("first"), ShouldBeTrue)
		So(f.IsRequired(), ShouldBeFalse)
		So(f.GetType(), ShouldEqual, FIELD_STRUCT)
	})

	Convey("Test MarshalJSON", t, func() {
		f := newField()
		_, err := f.MarshalJSON()
		So(err, ShouldBeNil)
	})

	Convey("Test Choices", t, func() {
		f := newField()
		So(f.Choices(), ShouldHaveSameTypeAs, newChoices())
	})
}
