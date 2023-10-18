package main

import (
	"encoding/json"
	"github.com/dave/jennifer/jen"
)

func GenerateStructFromJSON(jsonStr string) (*jen.File, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}

	file := jen.NewFile("main")

	err = generateStruct(file, "User", data, 1)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func generateStruct(file *jen.File, name string, data interface{}, level int) error {
	switch v := data.(type) {
	case map[string]interface{}:
		if level > 5 {
			return nil
		}
		file.Type().Id(name).StructFunc(func(g *jen.Group) {
			for key, value := range v {
				fieldName := jen.Id(key).Add(generateFieldType(value, level+1))
				g.Add(fieldName)
			}
		})
	case []interface{}:
		if len(v) == 0 {
			file.Type().Id(name).Interface()
		} else {
			return generateStruct(file, name, v[0], level)
		}
	default:
		file.Type().Id(name).Add(generateFieldType(data, level))
	}
	return nil
}

func generateFieldType(data interface{}, level int) *jen.Statement {
	switch v := data.(type) {
	case map[string]interface{}:
		if level > 5 {
			return jen.Empty()
		}
		return jen.StructFunc(func(g *jen.Group) {
			for key, value := range v {
				fieldName := jen.Id(key).Add(generateFieldType(value, level+1))
				g.Add(fieldName)
			}
		})
	case []interface{}:
		if len(v) == 0 {
			return jen.Interface()
		} else {
			return jen.Index().Add(generateFieldType(v[0], level))
		}
	case float64:
		return jen.Int()
	case bool:
		return jen.Bool()
	case string:
		return jen.String()
	default:
		return jen.Empty()
	}
}

func main() {
	jsonStr := `{
		"field1": 20,
		"field2": "Hello World",
		"field3": true,
		"field4": {
			"nested1": 3.14,
			"nested2": ["a", "b", "c"],
			"nested3": {
				"nested1": 3.14,
				"nested2": ["a", "b", "c"],
				"nested3": "Shyntas"
			}
		},
		"field5": [1, 2, 3]
	}`

	//r := mux.NewRouter()
	//r.HandleFunc("/create_struct", handleRequestA).Methods("POST")
	//log.Println(http.ListenAndServe("127.0.0.1:8080", r))

	file, err := GenerateStructFromJSON(jsonStr)
	if err != nil {
		panic(err)
	}

	err = file.Save("./Lecture06_28.09.2023/generated_struct.go")
	if err != nil {
		panic(err)
	}
}

//func handleRequestA(w http.ResponseWriter, r *http.Request){
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		panic(err)
//	}
//	file, err := GenerateStructFromJSON(string(body))
//	if err != nil {
//		panic(err)
//	}
//
//	err = file.Save("./Lecture06_28.09.2023/generated_struct.go")
//	if err != nil {
//		panic(err)
//	}
//}
