package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	// reflect.Value型のNumField()メソッドでフィールド数を取得する（※　ポインタ型の値に対してNumField()は使えない）
	for i := 0; i < val.NumField(); i++ {
		// 各フィールドの値を取得
		field := val.Field(i)

		// フィールドの型をfield.Kind()で取得する
		switch field.Kind() {
		// 文字列だったら、field.String()で値を取得してfnを呼ぶ
		case reflect.String:
			fn(field.String())
		// 構造体だったら、field.Interface()で値を取得して再帰的にwalkを呼ぶ
		case reflect.Struct:
			walk(field.Interface(), fn)
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
