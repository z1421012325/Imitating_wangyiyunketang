package util


import q "github.com/skip2/go-qrcode"

func CreateQRcode(qrcode string)(bytes []byte,err error)  {

	bytes,err = q.Encode(qrcode,100,256)
	if err != nil {
		return
	}

	return
}