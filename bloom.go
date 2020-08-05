package bloom

import "github.com/garyburd/redigo/redis"

type Bloom struct {
	Conn      redis.Conn
	Key       string
	HashFuncs []F //保存hash函数
}

func NewBloom(con redis.Conn) *Bloom {
	return &Bloom{Conn: con, Key: "bloom", HashFuncs: NewFunc()}
}
func (b *Bloom) Add(str string) error {
	var err error
	for _, f := range b.HashFuncs {
		offset := f(str)
		_, err := b.Conn.Do("setbit", b.Key, offset, 1)
		if err != nil {
			return err
		}
	}
	return err
}
func (b *Bloom) Exist(str string) bool {
	var a int64 = 1
	for _, f := range b.HashFuncs {
		offset := f(str)
		bitValue, _ := b.Conn.Do("getbit", b.Key, offset)
		if bitValue != a {
			return false
		}
	}
	return true
}
