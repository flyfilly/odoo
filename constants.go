package odoo

const (
	MethodCreate      = "create"
	MethodSearchRead  = "search_read"
	MethodSearch      = "search"
	MethodRead        = "read"
	MethodUpdate      = "write"
	MethodDelete      = "unlink"
	MethodSearchCount = "search_count"
	MethodFieldsGet   = "fields_get"

	Equal                  = "="
	NotEqual               = "!="
	LessThanOrEqualTo      = "<="
	LessThan               = "<"
	GreaterThan            = ">"
	GreaterThanOrEqualTo   = ">="
	Like                   = "like"
	CaseInsensitiveLike    = "ilike"
	NotLike                = "not like"
	NotCaseInsensitiveLike = "not ilike"
	UnsetOrEqualTo         = "=?"
	In                     = "in"
	NotIn                  = "not in"
	ChildOf                = "child of"
	ParentOf               = "parent of"
	And                    = "&"
	Or                     = "|"
	Not                    = "!"

	SaleOrder = "sale.order"
)
