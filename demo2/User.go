// automatically generated, do not modify

package demo2

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type User struct {
	_tab flatbuffers.Table
}

func (rcv *User) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *User) Uiid() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *User) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func UserStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func UserAddUiid(builder *flatbuffers.Builder, uiid flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(uiid), 0) }
func UserAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0) }
func UserEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
