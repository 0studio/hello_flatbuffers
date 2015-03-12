// automatically generated, do not modify

package demo2

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type RoomInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsRoomInfo(buf []byte, offset flatbuffers.UOffsetT) *RoomInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RoomInfo{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *RoomInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RoomInfo) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RoomInfo) Users(obj *User, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(User)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RoomInfo) UsersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func RoomInfoStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func RoomInfoAddId(builder *flatbuffers.Builder, id uint32) { builder.PrependUint32Slot(0, id, 0) }
func RoomInfoAddUsers(builder *flatbuffers.Builder, users flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(users), 0) }
func RoomInfoStartUsersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func RoomInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
