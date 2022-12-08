package test

import (
	"lbf.com/go-dp/iterator_dp"
	"testing"
)

func TestIteratorDp(t *testing.T) {
	hdlMenu := iterator_dp.NewHdlMenu()

	hdlMenu.AddItem(iterator_dp.NewHdlMenuItem("清蒸石斑鱼", 1, 59.9)).
		AddItem(iterator_dp.NewHdlMenuItem("椰子鸡", 1, 96.9)).
		AddItem(iterator_dp.NewHdlMenuItem("王老吉", 6, 5.0)).
		AddItem(iterator_dp.NewHdlMenuItem("清蒸鱼", 1, 19.9)).
		AddItem(iterator_dp.NewHdlMenuItem("金镶玉（西红柿炒蛋）", 1, 26.9)).
		AddItem(iterator_dp.NewHdlMenuItem("皇家香米", 6, 6.9))

	iterator := hdlMenu.Iterator()
	t.Log(iterator.HasNext())
	for iterator.HasNext() {
		err, item := iterator.Next()
		if err != nil {
			return
		}
		t.Log(item)
	}

}
