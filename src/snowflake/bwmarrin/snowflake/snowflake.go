package snowflake

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	nodeBits        = 10
	stepBits        = 12
	nodeMax         = -1 ^ (-1 << nodeBits) // 2^10 -1  1023  机器位
	stepMask  int64 = -1 ^ (-1 << stepBits) // 2^12 -1 4095  序号位
	timeShift uint8 = nodeBits + stepBits
	nodeShift uint8 = stepBits
)

const encodeBase58Map = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var decodeBase58Map [256]byte

type ID int64

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ original []byte }

func (j JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(j.original))
}

func init() {
	for i := 0; i < len(encodeBase58Map); i++ {
		decodeBase58Map[i] = 0xFF
	}

	for i := 0; i < len(encodeBase58Map); i++ {
		decodeBase58Map[encodeBase58Map[i]] = byte(i)
	}
}

var ErrInvalidBase58 = errors.New("invalid base58")

// Epoch is set to the twitter snowflake epoch of 2006-03-21:20:50:14 GMT
// You may customize this to set a different epoch for your application.
var Epoch int64 = 1288834974657

type Node struct {
	mu   sync.Mutex
	time int64
	node int64
	step int64
}

func NewNode(node int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}

	return &Node{
		time: 0,
		node: node,
		step: 0,
	}, nil
}

func (n *Node) Generate() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixNano() / 1000000

	if n.time == now {
		n.step = (n.step + 1) & stepMask
		if n.step == 0 {
			for now <= n.time {
				now = time.Now().UnixNano() / 100000
			}
		}
	} else {
		n.step = 0
	}

	n.time = now

	r := ID((now-Epoch)<<timeShift |
		(n.node << nodeShift) |
		(n.step),
	)

	return r
}

// Int64 returns an int64 of the snowflake ID
func (f ID) Int64() int64 {
	return int64(f)
}

func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10) // 格式化为10进制　字符串输出
}

func (f ID) Base2() string {
	return strconv.FormatInt(int64(f), 2) // 转换为二进制 字符串输出
}

// Base36 returns a base36 string of the snowflake ID
func (f ID) Base36() string {
	return strconv.FormatInt(int64(f), 36)
}

// Base58 returns a base58 string of the snowflake ID
func (f ID) Base58() string {

	if f < 58 {
		return string(encodeBase58Map[f])
	}

	b := make([]byte, 0, 11)
	for f >= 58 {
		b = append(b, encodeBase58Map[f%58])
		f /= 58
	}
	b = append(b, encodeBase58Map[f])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(b []byte) (ID, error) {

	var id int64

	for i := range b {
		if decodeBase58Map[b[i]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		id = id*58 + int64(decodeBase58Map[b[i]])
	}

	return ID(id), nil
}

func (f ID) Base64() string {
	return base64.StdEncoding.EncodeToString(f.Bytes())
}

func (f ID) Bytes() []byte {
	return []byte(f.String())
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (f ID) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(f))
	return b
}

// Time returns an int64 unix timestamp of the snowflake ID time
func (f ID) Time() int64 {
	return (int64(f) >> 22) + Epoch
}

// Node returns an int64 of the snowflake ID node number
func (f ID) Node() int64 {
	return int64(f) & 0x00000000003FF000 >> nodeShift
}

// Step returns an int64 of the snowflake step (or sequence) number
func (f ID) Step() int64 {
	return int64(f) & 0x0000000000000FFF
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (f ID) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(f), 10)
	buff = append(buff, '"')
	return buff, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (f *ID) UnmarshalJSON(b []byte) error {
	if len(b) < 3 || b[0] != '"' || b[len(b)-1] != '"' {
		return JSONSyntaxError{b}
	}

	i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}

	*f = ID(i)
	return nil
}
