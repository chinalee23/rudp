package rudp

import ()

const (
	_PACAKGE_LEN = 512
)

type stMessage struct {
	data []byte
	id   int
	tick int
}

type PackageBuffer struct {
	buffer [_PACAKGE_LEN]byte
	sz     int
}

type Rudp struct {
	curr_tick        int
	expired_interval int
	send_id          int
	recv_id_min      int
	recv_id_max      int

	send_queue    []*stMessage
	recv_queue    []*stMessage
	history_queue []*stMessage
	request_queue []int

	pkg_queue []*PackageBuffer
}

func New() *Rudp {
	return &Rudp{
		curr_tick:        0,
		expired_interval: 10,
		send_id:          0,
		recv_id_min:      0,
		recv_id_max:      0,
	}
}

func new_message(buffer []byte, sz int) *stMessage {
	return &stMessage{
		data: buffer[:sz],
		id:   0,
		tich: 0,
	}
}

func fill_header(buffer []byte, sz int, id int) {
	var offset int
	if sz < 128 {
		buffer[0] = byte(sz)
		offset = 1
	} else {
		buffer[0] = (sz & 0x7f00) >> 8
		buffer[1] = sz & 0xff
		offset = 2
	}

	buffer[offset] = (id & 0xff00) >> 8
	buffer[offset+1] = id & 0xff
}

func get_id(buffer []byte) int {
	return buffer[0]*256 + buffer[1]
}

func (p *Rudp) Send(data []byte, sz int) {

}

func (p *Rudp) Recv(buffer []byte) int {
	return 0
}

func (p *Rudp) Update(data []byte, sz int) {

}
