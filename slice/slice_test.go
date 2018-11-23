package main

import (
	"reflect"
	"strconv"
	"testing"
)

// スライスの初期化の確認
func TestSliceLen(t *testing.T) {
	slice := make([]int, 3)
	expectLen := 3
	actualLen := len(slice)

	if expectLen != actualLen {
		t.Errorf("%d != %d", expectLen, actualLen)
	}

	expectValues := []int{0, 0, 0}
	actualValue := slice

	if !reflect.DeepEqual(expectValues, actualValue) {
		t.Errorf("%d != %d", expectValues, expectValues)
	}
}

// 第三引数が容量になる確認
func TestSliceCap(t *testing.T) {
	slice := make([]int, 10, 20)
	expectCap := 20
	actualCap := cap(slice)

	if expectCap != actualCap {
		t.Errorf("%d != %d", expectCap, actualCap)
	}
}

// スライスの代入の確認
func TestSliceValue(t *testing.T) {
	slice := make([]int, 10)
	expect1 := 0
	actual1 := slice[3]

	if expect1 != actual1 {
		t.Errorf("%d != %d", expect1, actual1)
	}

	slice[3] = 4
	expect2 := 4
	actual2 := slice[3]

	if expect2 != actual2 {
		t.Errorf("%d != %d", expect2, actual2)
	}
}

// make 以外にスライスを作るリテラルの確認
func TestSliceLiteral(t *testing.T) {
	expect := make([]int, 3)
	actual := []int{0, 0, 0}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// 配列から一部を簡易スライスとして生成する
func TestSimpleSlice(t *testing.T) {
	base := [5]int{1, 2, 3, 4, 5}

	expect1 := []int{2, 3}
	actual1 := base[1:3]

	if !reflect.DeepEqual(expect1, actual1) {
		t.Errorf("%v != %v", expect1, actual1)
	}

	expect2 := []int{1, 2, 3}
	actual2 := base[:3]

	if !reflect.DeepEqual(expect2, actual2) {
		t.Errorf("%v != %v", expect2, actual2)
	}

	expect3 := []int{4, 5}
	actual3 := base[3:]

	if !reflect.DeepEqual(expect3, actual3) {
		t.Errorf("%v != %v", expect2, actual2)
	}
}

// 文字列の場合の簡易スライス生成確認
func TestStringSlice(t *testing.T) {
	expect := "BC"
	actual := "ABCDE"[1:3]

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

// マルチバイト
func TestMultiByteSlice(t *testing.T) {
	expectLen := 3
	actualLen := len("あ")

	if expectLen != actualLen {
		t.Errorf("%d != %d", expectLen, actualLen)
	}

	expect := "いう"
	actual := "あいうえお"[3:9]

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

// appendの動作確認
func TestAppend(t *testing.T) {
	base := []int{1, 2, 3}
	expect := []int{1, 2, 3, 4}
	actual := append(base, 4)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// スライス同士のマージは第二引数の末尾が特殊
func TestSliceMerge(t *testing.T) {
	base1 := []int{1}
	base2 := []int{2}

	expect := []int{1, 2}
	actual := append(base1, base2...)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// マルチバイトのまーじ
func TestByteSliceMerge(t *testing.T) {
	var b []byte
	expect := []byte{227, 129, 130, 227, 129, 132, 227, 129, 134, 227, 129, 136, 227, 129, 138}
	actual := append(b, "あいうえお"...)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// copy は要素を上書きする
func TestCopy(t *testing.T) {
	base1 := []int{1, 2, 3, 4, 5}
	base2 := []int{10, 11}
	copy(base1, base2)

	expect := []int{10, 11, 3, 4, 5}
	actual := base1

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// copy は第一引数を元として容量を超えた分は丸められる
func TestCopyRounded(t *testing.T) {
	base1 := []int{1, 2}
	base2 := []int{10, 11, 12, 13, 14}
	copy(base1, base2)

	expect := []int{10, 11}
	actual := base1

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// range なら ループ中要素数が変わっても無限ループにならない
func TestRangeFor(t *testing.T) {
	expect := []string{"A", "B", "C", "A_0", "B_1", "C_2"}
	actual := []string{"A", "B", "C"}
	for i, v := range actual {
		actual = append(actual, v+"_"+strconv.Itoa(i))
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// 配列だと値渡しになって変化ない
func TestPassByValueArray(t *testing.T) {
	pow := func(a [3]int) {
		for i, v := range a {
			a[i] = v * v
		}
		return
	}
	expect := [3]int{1, 2, 3}
	actual := [3]int{1, 2, 3}
	pow(actual)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// スライスの場合参照渡しになる
func TestReferenceSlice(t *testing.T) {
	pow := func(a []int) {
		for i, v := range a {
			a[i] = v * v
		}
		return
	}
	expect := []int{1, 4, 9}
	actual := []int{1, 2, 3}
	pow(actual)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}
}

// スライスさん宣言だけの場合の初期値は nil
func TestInitialValueNilSlice(t *testing.T) {
	var actual []int
	if nil != actual {
		t.Errorf("%v", actual)
	}
}
