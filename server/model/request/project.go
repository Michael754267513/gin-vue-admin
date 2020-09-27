package request

import "gin-vue-admin/model"

type ProjectSearch struct {
	model.Project
	PageInfo
}
