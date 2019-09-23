package odoo

const (
	// MethodCreate is used to create an object in Odoo
	MethodCreate = "create"
	// MethodSearchRead will perform a search and immediately invoke the read method
	MethodSearchRead = "search_read"
	// MethodSearch executes the search method
	MethodSearch = "search"
	// MethodRead executes the read method
	MethodRead = "read"
	// MethodUpdate executes an update method on an Odoo object
	MethodUpdate = "write"
	// MethodDelete is used to delete an object from Odoo
	MethodDelete = "unlink"
	// MethodSearchCount will perform a search on provided conditions and count the number of results
	MethodSearchCount = "search_count"
	// MethodFieldsGet will return all applicable fields in the queried object
	MethodFieldsGet = "fields_get"
)

const (
	// Equal is the equality operator
	Equal = "="
	// NotEqual is the equality operator that is used to check for unequality
	NotEqual = "!="
	// LessThanOrEqualTo is used to check if numeric values are less than or equal to a value
	LessThanOrEqualTo = "<="
	// LessThan is used to check if a numeric value is less than a value
	LessThan = "<"
	// GreaterThan is used to check if a numeric value is greater than a value
	GreaterThan = ">"
	// GreaterThanOrEqualTo is used to check if a numeric value is greater than or equal to a value
	GreaterThanOrEqualTo = ">="
	// Like is used to check if a string value is like another string value and is case sensitive
	Like = "like"
	// CaseInsensitiveLike is used to check if a string value is like another string value and is not case sensitive
	CaseInsensitiveLike = "ilike"
	// NotLike will evaluate if a string is not like a given string and is case sensitive
	NotLike = "not like"
	// NotCaseInsensitiveLike will evaluate if a string is not like a given string and is not case sensitive
	NotCaseInsensitiveLike = "not ilike"
	// UnsetOrEqualTo is used for whatever this is used for... If you know Odoo then this probably makes sense
	UnsetOrEqualTo = "=?"
	// In checks if a value exists in a list of values
	In = "in"
	// NotIn checks if a values does not exist in a list of values
	NotIn = "not in"
	// ChildOf checks if the object is a child of another object
	ChildOf = "child of"
	// ParentOf checks if the object is a parent of a child object
	ParentOf = "parent of"
	// And concatinates conditions and ensures the condition is also met
	And = "&"
	// Or adds an additional condition
	Or = "|"
	// Not will check against the opposing result of a condition
	Not = "!"
)

const (
	// ModelSaleOrder refers to the sale order model
	ModelSaleOrder = "sale.order"
)
