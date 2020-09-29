package request

import "gin-vue-admin/model"

type CvsSearch struct {
	model.Cvs
	PageInfo
}
