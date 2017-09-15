package mongodbhelper

import (
    "encoding/json"
    "bytes"
    "fmt"
)

//实现mongo json 查询

const KEY = "key"
const dataWraper = "value"
const NS = "chaincodeid"

var VALID_OPERATORS = []string{
    "$eq","$gt","$gte","$in","$lt","$lte","$ne","$nin","$and","$not","$nor","$or","$exists",
    "$type","$mod","$regex","$text","$where","$geoIntersects","$geoWithin","$near","$nearSphere",
    "$all","$elemMatch","$size","$bitsAllClear","$bitsAnyClear","$bitsAnySet","$comment","$","$elemMatch",
    "$meta","$slice",
}

//当到该函数时,应保证query已经是一个json字符串了
func GetQueryBson(namespace, query string) (interface{}, error) {
    //创建查询jsonmap
    jsonQueryMap := make(map[string]interface{})
    
    //unmarshal the selected json into the generic map
    decoder := json.NewDecoder(bytes.NewBuffer([]byte(query)))
    decoder.UseNumber()
    err := decoder.Decode(&jsonQueryMap)
    if err != nil {
        return "", err
    }
    processAndWrapQuery(jsonQueryMap)
    jsonQueryMap[NS] = namespace
    return jsonQueryMap, nil
}

//对field和value分别进行处理
func processAndWrapQuery(jsonQueryMap map[string]interface{}) {
    for jsonKey, jsonValue := range jsonQueryMap {
        wraperField := processField(jsonKey)
        delete(jsonQueryMap, jsonKey)
        wraperValue := processValue(jsonValue)
        jsonQueryMap[wraperField] = wraperValue
    }
}

//处理field
func processField(field string) string {
    if isValidOperator(field) {
        return field
    } else {
        return fmt.Sprintf("%v.%v", dataWraper, field)
    }
}

//处理value
func processValue(value interface{}) interface{} {
    switch value.(type) {
    case string:
        return value

    case json.Number:
        return value
    
    case []interface{}:
        arrayValue := value.([]interface{})
        resultValue := make([]interface{}, 0)
        for _, itemValue := range arrayValue {
            resultValue = append(resultValue, processValue(itemValue.(interface{})))
        }
        return resultValue
    
    case interface{}:
        processAndWrapQuery(value.(map[string]interface{}))
        return value
        
    default:
        return value
    }
}

func isValidOperator(field string) bool {
    return arrayContains(VALID_OPERATORS, field)
}

func arrayContains(source []string, item string) bool {
    for _, s := range source {
        if s == item {
            return true
        }
    }
    return false
}
