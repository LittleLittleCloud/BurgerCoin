package burgercoin

type meta struct {
	TimeStamp int32
	DataPad   [1024]byte
}

type certification struct {
	Pub [1024]byte

	// signiture of last block
	SigLast [256]byte
	// signiture of this block
	SigThis [256]byte
}
type BurgerCoin struct {
	Burble        [256]byte
	UserID        [128]byte
	Meta          meta
	Certification certification
}

type BaseCoin interface {
	// TODO
	// CRUD
}
