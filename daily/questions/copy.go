package questions

import "log"

var array [1024]byte

func CopySlice() {
	str := `Kasme=md5
	AUTH=offical@hebeiyidong.3gpp.net
	imsi=6651234545135
	RAND=658707243`
	copy(array[:], []byte(str))
	log.Println(array)
	log.Println(str)
	str = `imsi=6651234545135
	Kasme=md5`
	copy(array[:], []byte(str))
	log.Println(array, []byte(str))
}
