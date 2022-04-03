package hasher

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/vlpolak/swtgo/pkg/hasher/msg"
	"log"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {

	tests := []struct {
		password string
		checked  bool
	}{
		{
			password: "122234234",
			checked:  true,
		},
		{
			password: "FGghjhjh223123231",
			checked:  true,
		},
		{
			password: "jljkl12123hhihihi",
			checked:  true,
		},
		{
			password: "qeqwe23423434nmjjkjk",
			checked:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			hp, err := HashPassword(tt.password)
			assert.Nil(t, err)
			assert.NotNil(t, hp)
			ch := CheckPasswordHash(tt.password, hp)
			assert.True(t, ch)
		})
	}

	msg1 := &msg.Msg{
		Key:   "122234234",
		Value: true,
	}

	data, err := proto.Marshal(msg1)
	if err != nil {
		log.Fatal("marshaling error: ", err)
		return
	}

	msg2 := new(msg.Msg)
	err = proto.Unmarshal(data, msg2)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Printf("msg2: %s", msg2)

	log.Println("Done")
}
