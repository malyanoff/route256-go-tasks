/*
Есть JSON, состоящий из массива объектов, в каждом объекте которого есть поле data. В поле data лежит объект, в котором гарантируется наличие строкового ключа key`.
Длина массива 0 ≤ x ≤ 105.
Требуется посчитать количество уникальных значений среди всех key во всех data.


Input data format
На вход подаётся число 0 < N ≤ 105, далее идут N строк, содержащие в себе JSON файл. Гарантируется, что длина каждой строки l < 105 символов.

Output data format
На выход требуется отправить единственное число – количество уникальных ключей во всех data.
*/

package solution

import (
	"encoding/json"
)



type Data struct {
	Key string `json:"key"`
}

type JsonObj struct {
	Data Data
}

func CountUniqueKeys(input string) int {

	var data []JsonObj
	json.Unmarshal([]byte(input),&data)

	//collect all keys from data.key values
	keys := make(map[string]bool)
	items:= []string{}

	for _, obj := range data {
		items = append(items,obj.Data.Key)
	}

	resultArr := []string{}
	//remove duplicates from keys
	for _, item := range items{
		if _, value := keys[item]; !value {
			keys[item] = true
			resultArr = append(resultArr, item)
			
		}
	}


	return len(resultArr)
}

