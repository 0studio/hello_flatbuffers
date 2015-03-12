package main

// https://groups.google.com/forum/#!searchin/flatbuffers/go/flatbuffers/a51RIPFfRLo/LyDB1T8B460J
import (
	"fmt"
	"github.com/0studio/hello_flatbuffers/demo"
	"github.com/0studio/hello_flatbuffers/demo2"
	flatbuffers "github.com/google/flatbuffers/go"
	"io/ioutil"
	"os"
)

func testdemo1() {
	builder := flatbuffers.NewBuilder(0)

	demo.MonsterStartInventoryVector(builder, 5)
	for i := 4; i >= 0; i-- {
		builder.PrependByte(byte(i))
	}
	inv := builder.EndVector(5)
	// 对 inventory vector 的构建要在 monsterStart 开始之前
	// 即 一个 start 必然对应一个end
	// 不可以嵌套
	// 不可以  MonsterStart MonsterStartInventoryVector  EndVector, MonsterEnd 的顺序

	name := builder.CreateString("monsterName")
	demo.MonsterStart(builder)
	demo.MonsterAddHp(builder, 100)
	demo.MonsterAddName(builder, name)
	// builder2 := flatbuffers.NewBuilder(0)
	demo.MonsterAddInventory(builder, inv)

	offset := demo.MonsterEnd(builder)
	builder.Finish(offset)
	byteData := builder.Bytes
	monster := demo.GetRootAsMonster(byteData, builder.Head())
	if monster.Hp() != 100 {
		fmt.Println("monster.hp!=100 ,that is wrong")
	} else {
		fmt.Println("monster.hp=", monster.Hp())
		fmt.Println("monster.name=", monster.Name())
		fmt.Println("monster.InventoryLength() should==5", monster.InventoryLength())
		fmt.Println("monster.Inventory[0] should==0", monster.Inventory(0))
		fmt.Println("monster.Inventory[1] should==1", monster.Inventory(1))
	}
}
func testdemo2() {
	fmt.Println("------------------------------------------------------------------------")
	builder := flatbuffers.NewBuilder(0)

	// create user 1
	demo2.UserStart(builder)
	demo2.UserAddName(builder, builder.CreateString("user1"))
	demo2.UserAddUiid(builder, builder.CreateString("user1_uuid"))
	user1 := demo2.UserEnd(builder)

	// create user 2
	demo2.UserStart(builder)
	demo2.UserAddName(builder, builder.CreateString("user2222"))
	demo2.UserAddUiid(builder, builder.CreateString("user2222_uuid"))
	user2 := demo2.UserEnd(builder)

	// create vector
	demo2.RoomInfoStartUsersVector(builder, 2)
	builder.PrependUOffsetT(user2) // prepend 所以要倒序 ， 先prepend user2 ,后preend user1
	builder.PrependUOffsetT(user1)
	vec := builder.EndVector(2)

	/// write roomOffset info
	demo2.RoomInfoStart(builder)
	demo2.RoomInfoAddId(builder, 1000)
	demo2.RoomInfoAddUsers(builder, vec)
	roomOffset := demo2.RoomInfoEnd(builder)
	builder.Finish(roomOffset)

	byteData := builder.Bytes
	roominfo := demo2.GetRootAsRoomInfo(byteData, builder.Head())
	fmt.Println("roominfo.Id should =10000", roominfo.Id())
	var userObje demo2.User
	roominfo.Users(&userObje, 0)
	fmt.Println("roominfo.User[0].Name should user1 :", userObje.Name())
	fmt.Println("roominfo.User[0].Uiid should user1_uuid:", userObje.Uiid())
	ioutil.WriteFile("/tmp/roominfo.go.data", builder.Bytes[builder.Head():], os.FileMode(0644))
}

func main() {
	testdemo1()
	testdemo2()

}
