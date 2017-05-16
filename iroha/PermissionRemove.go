// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PermissionRemove struct {
	_tab flatbuffers.Table
}

func GetRootAsPermissionRemove(buf []byte, offset flatbuffers.UOffsetT) *PermissionRemove {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PermissionRemove{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PermissionRemove) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PermissionRemove) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PermissionRemove) TargetAccount() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PermissionRemove) PermissionType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PermissionRemove) MutatePermissionType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *PermissionRemove) Permission(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func PermissionRemoveStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PermissionRemoveAddTargetAccount(builder *flatbuffers.Builder, targetAccount flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(targetAccount), 0)
}
func PermissionRemoveAddPermissionType(builder *flatbuffers.Builder, permissionType byte) {
	builder.PrependByteSlot(1, permissionType, 0)
}
func PermissionRemoveAddPermission(builder *flatbuffers.Builder, permission flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(permission), 0)
}
func PermissionRemoveEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
