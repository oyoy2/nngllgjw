package utils

import (
	"strings"
)

func ExcludeGPA(name string) bool {
	excludedCourses := []string{
		"职业发展与就业指导（一）",
		"职业发展与就业指导（二）",
		"职业发展与就业指导（三）",
		"职业发展与就业指导（四）",
		"职业发展与就业指导（五）",
		"就业指导与创业基础（一）",
		"就业指导与创业基础（二）",
		"就业指导与创业基础（三）",
		"创业基础（一）",
		"创业基础（二）",
		"创业基础（三）",
		"大学生心理健康（一）",
		"大学生心理健康（二）",
		"形势与政策（一）",
		"形势与政策（二）",
		"形势与政策（三）",
		"形势与政策（四）",
		"形势与政策（五）",
		"形势与政策（六）",
		"形势与政策（七）",
		"形势与政策（八）",
		"形势与政策",
		"体育（一）",
		"体育（二）",
		"体育（三）",
		"体育（四）",
		"大学生安全教育（十）",
		"大学生安全教育（一）",
		"大学生安全教育（二）",
		"大学生安全教育（三）",
		"大学生安全教育（四）",
		"大学生安全教育（五）",
		"大学生安全教育（六）",
		"大学生安全教育（七）",
		"大学生安全教育（八）",
		"大学生安全教育（九）",
		"职业发展与就业指导（一）",
		"职业发展与就业指导（二）",
		"职业发展与就业指导（三）",
		"职业发展与就业指导（四）",
		"职业发展与就业指导（五）",
		"就业指导与创业基础（一）",
		"就业指导与创业基础（二）",
		"就业指导与创业基础（三）",
		"创业基础（一）",
		"创业基础（二）",
		"大学生心理学（一）",
		"大学生心理学（二）",
		"大学生心理健康教育（一）",
		"大学生心理健康教育（二）",
		"形势与政策（一）",
		"形势与政策（二）",
		"形势与政策（三）",
		"形势与政策（四）",
		"形势与政策（五）",
		"形势与政策（六）",
		"形势与政策",
		"体育（一）",
		"体育（二）",
		"体育（三）",
		"体育（四）",
		"大学生安全教育（一）",
		"大学生安全教育（二）",
		"大学生安全教育（三）",
		"大学生安全教育（四）",
		"大学生安全教育（五）",
		"大学生安全教育（六）",
	}

	for _, course := range excludedCourses {
		if strings.Contains(course, name) {
			return false
		}
	}
	return true
}
