package main

type InputStruct struct {
	Number1 *int `json:"number1" binding:"required"`
	Number2 *int `json:"number2" binding:"required"`
}

type ResultStruct struct {
	Result int `json:"result"`
}
