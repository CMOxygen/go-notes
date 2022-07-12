package encryption

import "crypto/sha512"

func EncryptSHA(data []byte) (r []byte, err error) {
	h := sha512.New()

	_, err = h.Write(data)
	if err != nil {
		return nil, err
	}
	r = h.Sum(nil)

	return r, nil
}
