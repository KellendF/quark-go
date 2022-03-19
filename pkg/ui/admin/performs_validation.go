package admin

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/derekstavis/go-qs"
	"github.com/gofiber/fiber/v2"
)

// 创建请求的验证器
func (p *Resource) ValidatorForCreation(c *fiber.Ctx, resourceInstance interface{}, data interface{}) interface{} {
	ruleData := p.RulesForCreation(c, resourceInstance)

	//  return Validator::make($data ? $data : $request->all(), $ruleData['rules'], $ruleData['messages'])
	// 		 ->after(function ($validator) use ($request) {
	// 			 static::afterValidation($request, $validator);
	// 			 static::afterCreationValidation($request, $validator);
	// 		 });

	return ruleData
}

// 创建请求的验证规则
func (p *Resource) RulesForCreation(c *fiber.Ctx, resourceInstance interface{}) map[string]interface{} {

	fields := resourceInstance.(interface {
		CreationFieldsWithoutWhen(*fiber.Ctx, interface{}) interface{}
	}).CreationFieldsWithoutWhen(c, resourceInstance)

	rules := []interface{}{}
	ruleMessages := []interface{}{}

	for _, v := range fields.([]interface{}) {
		getResult := p.getRulesForCreation(c, v)
		rules = append(rules, getResult["rules"])
		ruleMessages = append(ruleMessages, getResult["messages"])

		when := reflect.
			ValueOf(v).
			Elem().
			FieldByName("When").Interface()

		if when != nil {
			whenItems := reflect.
				ValueOf(when).
				Elem().
				FieldByName("Items").Interface()

			if whenItems != nil {
				for _, vi := range whenItems.([]map[string]interface{}) {
					if p.needValidateWhenRules(c, vi) {
						body := vi["body"]
						if body != nil {
							// 如果为数组
							getBody, ok := body.([]interface{})
							if ok {
								for _, bv := range getBody {
									whenItemResult := p.getRulesForCreation(c, bv)
									rules = append(rules, whenItemResult["rules"])
									ruleMessages = append(ruleMessages, whenItemResult["messages"])
								}
							} else {
								whenItemResult := p.getRulesForCreation(c, getBody)
								rules = append(rules, whenItemResult["rules"])
								ruleMessages = append(ruleMessages, whenItemResult["messages"])
							}
						}
					}
				}
			}
		}

	}

	return map[string]interface{}{
		"rules":    rules,
		"messages": ruleMessages,
	}
}

// 判断是否需要验证When组件里的规则
func (p *Resource) needValidateWhenRules(c *fiber.Ctx, when map[string]interface{}) bool {
	conditionName := when["condition_name"].(string)
	conditionOption := when["condition_option"]
	conditionOperator := when["condition_operator"].(string)
	result := false

	data, error := qs.Unmarshal(c.OriginalURL())
	if error != nil {
		return false
	}

	value, ok := data[conditionName]
	if !ok {
		return false
	}

	valueString, isString := value.(string)
	if isString {
		if valueString == "" {
			return false
		}
	}

	switch conditionOperator {
	case "=":
		result = (value.(string) == conditionOption.(string))
	case ">":
		result = (value.(string) > conditionOption.(string))
	case "<":
		result = (value.(string) < conditionOption.(string))
	case "<=":
		result = (value.(string) <= conditionOption.(string))
	case ">=":
		result = (value.(string) >= conditionOption.(string))
	case "has":
		_, isArray := value.([]string)
		if isArray {
			getJson, err := json.Marshal(value)
			if err != nil {
				result = strings.Contains(string(getJson), conditionOption.(string))
			}
		} else {
			result = strings.Contains(value.(string), conditionOption.(string))
		}
	case "in":
		conditionOptionArray, isArray := conditionOption.([]string)
		if isArray {
			for _, v := range conditionOptionArray {
				if v == value.(string) {
					result = true
				}
			}
		}
	default:
		result = (value.(string) == conditionOption)
	}

	return result
}

// 获取创建请求资源规则
func (p *Resource) getRulesForCreation(c *fiber.Ctx, field interface{}) map[string]interface{} {
	getRules := map[string]interface{}{}
	getRuleMessages := map[string]interface{}{}

	name := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Name").String()

	rules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Rules").Interface()

	ruleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("RuleMessages").Interface()

	creationRules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("CreationRules").Interface()

	creationRuleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("CreationRuleMessages").Interface()

	items := []interface{}{}

	for _, v := range p.formatRules(c, rules.([]string)) {
		items = append(items, v)
	}

	for key, v := range ruleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	for _, v := range p.formatRules(c, creationRules.([]string)) {
		items = append(items, v)
	}

	for key, v := range creationRuleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	if len(items) > 0 {
		getRules[name] = items
	}

	return map[string]interface{}{
		"rules":    getRules,
		"messages": getRuleMessages,
	}
}

// 更新请求的验证器
func (p *Resource) ValidatorForUpdate(c *fiber.Ctx, resourceInstance interface{}, data interface{}) interface{} {
	ruleData := p.RulesForUpdate(c, resourceInstance)

	//  return Validator::make($data ? $data : $request->all(), $ruleData['rules'], $ruleData['messages'])
	// 		 ->after(function ($validator) use ($request) {
	// 			 static::afterValidation($request, $validator);
	// 			 static::afterImportValidation($request, $validator);
	// 		 });
	return ruleData
}

// 更新请求的验证规则
func (p *Resource) RulesForUpdate(c *fiber.Ctx, resourceInstance interface{}) map[string]interface{} {

	fields := resourceInstance.(interface {
		UpdateFieldsWithoutWhen(*fiber.Ctx, interface{}) interface{}
	}).UpdateFieldsWithoutWhen(c, resourceInstance)

	rules := []interface{}{}
	ruleMessages := []interface{}{}

	for _, v := range fields.([]interface{}) {
		getResult := p.getRulesForUpdate(c, v)
		rules = append(rules, getResult["rules"])
		ruleMessages = append(ruleMessages, getResult["messages"])

		when := reflect.
			ValueOf(v).
			Elem().
			FieldByName("When").Interface()

		if when != nil {
			whenItems := reflect.
				ValueOf(when).
				Elem().
				FieldByName("Items").Interface()

			if whenItems != nil {
				for _, vi := range whenItems.([]map[string]interface{}) {
					if p.needValidateWhenRules(c, vi) {
						body := vi["body"]

						if body != nil {

							// 如果为数组
							getBody, ok := body.([]interface{})
							if ok {
								for _, bv := range getBody {
									whenItemResult := p.getRulesForUpdate(c, bv)
									rules = append(rules, whenItemResult["rules"])
									ruleMessages = append(ruleMessages, whenItemResult["messages"])
								}
							} else {
								whenItemResult := p.getRulesForUpdate(c, getBody)
								rules = append(rules, whenItemResult["rules"])
								ruleMessages = append(ruleMessages, whenItemResult["messages"])
							}
						}
					}
				}
			}
		}

	}

	return map[string]interface{}{
		"rules":    rules,
		"messages": ruleMessages,
	}
}

// 获取更新请求资源规则
func (p *Resource) getRulesForUpdate(c *fiber.Ctx, field interface{}) map[string]interface{} {

	getRules := map[string]interface{}{}
	getRuleMessages := map[string]interface{}{}

	name := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Name").String()

	rules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Rules").Interface()

	ruleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("RuleMessages").Interface()

	updateRules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("UpdateRules").Interface()

	updateRuleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("UpdateRuleMessages").Interface()

	items := []interface{}{}

	for _, v := range p.formatRules(c, rules.([]string)) {
		items = append(items, v)
	}

	for key, v := range ruleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	for _, v := range p.formatRules(c, updateRules.([]string)) {
		items = append(items, v)
	}

	for key, v := range updateRuleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	if len(items) > 0 {
		getRules[name] = items
	}

	return map[string]interface{}{
		"rules":    getRules,
		"messages": getRuleMessages,
	}
}

// 导入请求的验证器
func (p *Resource) ValidatorForImport(c *fiber.Ctx, resourceInstance interface{}, data interface{}) interface{} {
	ruleData := p.RulesForImport(c, resourceInstance)

	//  return Validator::make($data ? $data : $request->all(), $ruleData['rules'], $ruleData['messages'])
	// 		 ->after(function ($validator) use ($request) {
	// 			 static::afterValidation($request, $validator);
	// 			 static::afterImportValidation($request, $validator);
	// 		 });

	return ruleData
}

// 创建请求的验证规则
func (p *Resource) RulesForImport(c *fiber.Ctx, resourceInstance interface{}) map[string]interface{} {

	fields := resourceInstance.(interface {
		ImportFieldsWithoutWhen(*fiber.Ctx, interface{}) interface{}
	}).ImportFieldsWithoutWhen(c, resourceInstance)

	rules := []interface{}{}
	ruleMessages := []interface{}{}

	for _, v := range fields.([]map[string]interface{}) {
		getResult := p.getRulesForCreation(c, v)
		rules = append(rules, getResult["rules"])
		ruleMessages = append(ruleMessages, getResult["messages"])

		when := reflect.
			ValueOf(v).
			Elem().
			FieldByName("When").Interface()

		if when != nil {
			whenItems := reflect.
				ValueOf(when).
				Elem().
				FieldByName("Items").Interface()

			if whenItems != nil {
				for _, vi := range whenItems.([]map[string]interface{}) {
					if p.needValidateWhenRules(c, vi) {
						body := vi["body"]

						if body != nil {

							// 如果为数组
							getBody, ok := body.([]interface{})
							if ok {
								for _, bv := range getBody {
									whenItemResult := p.getRulesForCreation(c, bv)
									rules = append(rules, whenItemResult["rules"])
									ruleMessages = append(ruleMessages, whenItemResult["messages"])
								}
							} else {
								whenItemResult := p.getRulesForCreation(c, getBody)
								rules = append(rules, whenItemResult["rules"])
								ruleMessages = append(ruleMessages, whenItemResult["messages"])
							}
						}
					}
				}
			}
		}

	}

	return map[string]interface{}{
		"rules":    rules,
		"messages": ruleMessages,
	}
}

// 格式化规则
func (p *Resource) formatRules(c *fiber.Ctx, rules []string) []string {
	id := c.Query("id")

	if id == "" {
		return rules
	}

	for key, v := range rules {
		rules[key] = strings.Replace(v, "{id}", id, -1)
	}

	return rules
}

// 验证完成后回调
func (p *Resource) afterValidation(c *fiber.Ctx, validator interface{}) {
	//
}

// 创建请求验证完成后回调
func (p *Resource) afterCreationValidation(c *fiber.Ctx, validator interface{}) {
	//
}

// 更新请求验证完成后回调
func (p *Resource) afterUpdateValidation(c *fiber.Ctx, validator interface{}) {
	//
}

// 创建请求验证完成后回调
func (p *Resource) afterImportValidation(c *fiber.Ctx, validator interface{}) {
	//
}
