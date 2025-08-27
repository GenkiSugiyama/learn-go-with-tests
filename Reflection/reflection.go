package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numOfValues := 0
	var getField func(int) reflect.Value

	// フィールドの型をfield.Kind()で取得する
	switch val.Kind() {
	// 文字列だったら、field.String()で値を取得してfnを呼ぶ
	case reflect.String:
		fn(val.String())
	// 取得データが構造体なら再起的にwalkを呼ぶ
	case reflect.Struct:
		numOfValues = val.NumField()
		getField = val.Field
	// 取得データがスライスなら、各要素に対して再起的にwalkを呼ぶ
	case reflect.Slice:
		numOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	// reflect.ValueOf()で指定されたデータを取得する
	val := reflect.ValueOf(x)

	// 取得したデータがポインタ型なら、Elem()で中身を取得する
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
