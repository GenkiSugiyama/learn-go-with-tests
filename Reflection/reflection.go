package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	// フィールドの型をfield.Kind()で取得する
	switch val.Kind() {
	// 文字列だったら、field.String()で値を取得してfnを呼ぶ
	case reflect.String:
		fn(val.String())
	// 取得データが構造体なら再起的にwalkを呼ぶ
	case reflect.Struct:
		// reflect.Value.Field(i) でi番目のフィールドの値を取得する
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	// 取得データがスライスなら、各要素に対して再起的にwalkを呼ぶ
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		// MapKeys() でmapに設定されているキーのスライスを取得する
		for _, key := range val.MapKeys() {
			// MapIndex() で指定したキーに対応する値を取得する
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		// valが関数の場合、Call(引数)で関数を呼び出して結果を取得する
		// 受け取った結果の数分ループする
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
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
