
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package res;


import (
    "KryptonGo/pkg/luban"
)

import "errors"

type Reward struct {
    Id int32
    Name string
    Count int32
}

const TypeId_Reward = -1850459313

func (*Reward) GetTypeId() int32 {
    return -1850459313
}

func NewReward(_buf *luban.ByteBuf) (_v *Reward, err error) {
    _v = &Reward{}
    { if _v.Id, err = _buf.ReadInt(); err != nil { err = errors.New("error"); return } }
    { if _v.Name, err = _buf.ReadString(); err != nil { err = errors.New("error"); return } }
    { if _v.Count, err = _buf.ReadInt(); err != nil { err = errors.New("error"); return } }
    return
}

