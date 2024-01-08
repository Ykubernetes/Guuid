package utils

type Guid struct {
	machine []byte
	time    uint32
	rand    uint32
	id      uint32
}

func NewGuid() *Guid {
	guid := &Guid{
		machine: MachineId,
		time:    timeStamp(),
		rand:    myrand(),
		id:      DeltaIdCounter(),
	}
	return guid
}

func (g *Guid) Serial() []byte {
	id := make([]byte, 0, 36)
	id = append(id, g.machine[0:4]...)
	id = append(id, uint32ToBin(g.time)...)
	id = append(id, uint32ToBin(g.id)...)
	id = append(id, uint32ToBin(g.rand)...)
	return id
}

func CreateBaseId() []byte {
	guid := NewGuid()
	return guid.Serial()
}

func CreateSimpleUUID() string {
	return EncodeToString(CreateBaseId())
}

func CreateUUID() string {
	gid := make([]byte, 0, 36)
	id := []byte(EncodeToString(CreateBaseId()))

	gid = append(gid, id[0:8]...)
	gid = append(gid, '-')
	gid = append(gid, id[8:12]...)
	gid = append(gid, '-')
	gid = append(gid, id[12:16]...)
	gid = append(gid, '-')
	gid = append(gid, id[16:20]...)
	gid = append(gid, '-')
	gid = append(gid, id[20:]...)
	return string(gid)
}

var Data []string

func CreateMultiUUID(num int, isSimple bool) ([]string, error) {
	if num < 1 {
		num = 1
	}
	if num > 2000 {
		num = 2000
	}
	uuidChan := make(chan string, num)
	go func(isSimple bool, num int) {
		index := 0
		for {
			go func(isSimple bool) {
				if isSimple {
					uuidChan <- CreateSimpleUUID()
				} else {
					uuidChan <- CreateUUID()
				}
			}(isSimple)

			if index >= num {
				break
			}
			index++
		}
	}(isSimple, num)

	for {
		Data = append(Data, <-uuidChan)
		if len(Data) == num {
			break
		}
	}
	return Data, nil
}
