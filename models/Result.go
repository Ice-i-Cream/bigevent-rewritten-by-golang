package models

// Result 结构体表示操作的结果，使用泛型T来表示数据的类型
type Result[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// Success 创建一个表示操作成功的 Result 对象，使用泛型T来表示数据的类型
func Success[T any](data T) Result[T] {
	return Result[T]{
		Code:    0,
		Message: "操作成功",
		Data:    data,
	}
}

func SuccessNoData() Result[interface{}] {
	return Result[interface{}]{
		Code:    0,
		Message: "操作成功",
		Data:    nil,
	}
}

func Error(message string) Result[interface{}] {
	return Result[interface{}]{
		Code:    1,
		Message: message,
		Data:    nil,
	}
}
