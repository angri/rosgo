// Code generated by ros-gen-go.
// source: MapMetaData.msg
// DO NOT EDIT!
package nav_msgs

import (
	"io"

	"github.com/ppg/rosgo/msgs/geometry_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgMapMetaData struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMapMetaData) Text() string {
	return t.text
}

func (t *_MsgMapMetaData) Name() string {
	return t.name
}

func (t *_MsgMapMetaData) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMapMetaData) NewMessage() ros.Message {
	m := new(MapMetaData)

	return m
}

var (
	MsgMapMetaData = &_MsgMapMetaData{
		`# This hold basic information about the characterists of the OccupancyGrid

# The time at which the map was loaded
time map_load_time
# The map resolution [m/cell]
float32 resolution
# Map width [cells]
uint32 width
# Map height [cells]
uint32 height
# The origin of the map [m, m, rad].  This is the real-world pose of the
# cell (0,0) in the map.
geometry_msgs/Pose origin`,
		"nav_msgs/MapMetaData",
		"84f12218a664df5896e1e5f18e75d1b0",
	}
)

type MapMetaData struct {
	MapLoadTime ros.Time
	Resolution  float32
	Width       uint32
	Height      uint32
	Origin      geometry_msgs.Pose
}

func (m *MapMetaData) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "time", &m.MapLoadTime); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Resolution); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Width); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Height); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Pose", &m.Origin); err != nil {
		return err
	}

	return
}

func (m *MapMetaData) Deserialize(r io.Reader) (err error) {
	// MapLoadTime
	if err = ros.DeserializeMessageField(r, "time", &m.MapLoadTime); err != nil {
		return err
	}

	// Resolution
	if err = ros.DeserializeMessageField(r, "float32", &m.Resolution); err != nil {
		return err
	}

	// Width
	if err = ros.DeserializeMessageField(r, "uint32", &m.Width); err != nil {
		return err
	}

	// Height
	if err = ros.DeserializeMessageField(r, "uint32", &m.Height); err != nil {
		return err
	}

	// Origin
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Pose", &m.Origin); err != nil {
		return err
	}

	return
}
