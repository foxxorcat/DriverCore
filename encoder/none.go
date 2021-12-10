package encoder

const NONE = "none"

type None struct{}

func (n *None) Name() string {
	return NONE
}

func (n *None) Encoded(in []byte) (out []byte, err error) {
	return in, nil
}

func (n *None) Decode(in []byte) (out []byte, err error) {
	return in, nil
}
