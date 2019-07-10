package test

import (
	"reflect"
	"testing"
	burgercoin "../struct"
)

func TestBurgerCoinStruct(t *testing.T) {
	// TODO
	// test the structure of BurgerCoin
	bc := new(burgercoin.BurgerCoin)
	if len(bc.Burble) != 256 {
		t.Errorf("Burble's size should be 256")
	}

	if len(bc.UserID) != 128 {
		t.Errorf("UserID's size should be 128")
	}

	if reflect.ValueOf(bc.Meta.TimeStamp).String() != "<int32 Value>" {
		t.Errorf("Meta TimeStamp type should be int32")
	}
	if len(bc.Meta.DataPad) != 1024 {
		t.Errorf("Meta DataPad length should be 1024")
	}
	if len(bc.Certification.Pub) != 1024{
		t.Errorf("Certification Pub length should be 1024")
	}
	if len(bc.Certification.SigLast) != 256{
		t.Errorf("Certification SigLast length should be 256")
	}
	if len(bc.Certification.SigThis) != 256{
		t.Errorf("Certification SigThis length should be 256")
	}
}

func TestBurble(t *testing.T) {

}
