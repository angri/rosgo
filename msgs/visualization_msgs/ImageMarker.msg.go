// Code generated by ros-gen-go.
// source: ImageMarker.msg
// DO NOT EDIT!
package visualization_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/msgs/geometry_msgs"
	"github.com/ppg/rosgo/msgs/std_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgImageMarker struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgImageMarker) Text() string {
	return t.text
}

func (t *_MsgImageMarker) Name() string {
	return t.name
}

func (t *_MsgImageMarker) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgImageMarker) NewMessage() ros.Message {
	m := new(ImageMarker)

	return m
}

var (
	MsgImageMarker = &_MsgImageMarker{
		`uint8 CIRCLE=0
uint8 LINE_STRIP=1
uint8 LINE_LIST=2
uint8 POLYGON=3
uint8 POINTS=4

uint8 ADD=0
uint8 REMOVE=1

Header header
string ns		# namespace, used with id to form a unique id
int32 id          	# unique id within the namespace
int32 type        	# CIRCLE/LINE_STRIP/etc.
int32 action      	# ADD/REMOVE
geometry_msgs/Point position # 2D, in pixel-coords
float32 scale	 	# the diameter for a circle, etc.
std_msgs/ColorRGBA outline_color
uint8 filled		# whether to fill in the shape with color
std_msgs/ColorRGBA fill_color # color [0.0-1.0]
duration lifetime       # How long the object should last before being automatically deleted.  0 means forever


geometry_msgs/Point[] points # used for LINE_STRIP/LINE_LIST/POINTS/etc., 2D in pixel coords
std_msgs/ColorRGBA[] outline_colors # a color for each line, point, etc.`,
		"visualization_msgs/ImageMarker",
		"32f5635ee965c48f5d3f565629b204e2",
	}
)

type ImageMarker struct {
	Header        std_msgs.Header
	Ns            string
	ID            int32
	Type          int32
	Action        int32
	Position      geometry_msgs.Point
	Scale         float32
	OutlineColor  std_msgs.ColorRGBA
	Filled        uint8
	FillColor     std_msgs.ColorRGBA
	Lifetime      ros.Duration
	Points        []geometry_msgs.Point
	OutlineColors []std_msgs.ColorRGBA
}

func (m *ImageMarker) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Ns); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.ID); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.Type); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.Action); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Point", &m.Position); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Scale); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "std_msgs/ColorRGBA", &m.OutlineColor); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Filled); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "std_msgs/ColorRGBA", &m.FillColor); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "duration", &m.Lifetime); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Points)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Points {
		if err = ros.SerializeMessageField(w, "geometry_msgs/Point", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.OutlineColors)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.OutlineColors {
		if err = ros.SerializeMessageField(w, "std_msgs/ColorRGBA", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *ImageMarker) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Ns
	if err = ros.DeserializeMessageField(r, "string", &m.Ns); err != nil {
		return err
	}

	// ID
	if err = ros.DeserializeMessageField(r, "int32", &m.ID); err != nil {
		return err
	}

	// Type
	if err = ros.DeserializeMessageField(r, "int32", &m.Type); err != nil {
		return err
	}

	// Action
	if err = ros.DeserializeMessageField(r, "int32", &m.Action); err != nil {
		return err
	}

	// Position
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Point", &m.Position); err != nil {
		return err
	}

	// Scale
	if err = ros.DeserializeMessageField(r, "float32", &m.Scale); err != nil {
		return err
	}

	// OutlineColor
	if err = ros.DeserializeMessageField(r, "std_msgs/ColorRGBA", &m.OutlineColor); err != nil {
		return err
	}

	// Filled
	if err = ros.DeserializeMessageField(r, "uint8", &m.Filled); err != nil {
		return err
	}

	// FillColor
	if err = ros.DeserializeMessageField(r, "std_msgs/ColorRGBA", &m.FillColor); err != nil {
		return err
	}

	// Lifetime
	if err = ros.DeserializeMessageField(r, "duration", &m.Lifetime); err != nil {
		return err
	}

	// Points
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Points: %s", err)
		}
		m.Points = make([]geometry_msgs.Point, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "geometry_msgs/Point", &m.Points[i]); err != nil {
				return err
			}
		}
	}

	// OutlineColors
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for OutlineColors: %s", err)
		}
		m.OutlineColors = make([]std_msgs.ColorRGBA, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "std_msgs/ColorRGBA", &m.OutlineColors[i]); err != nil {
				return err
			}
		}
	}

	return
}
