package testexport

// PublicFunc is an exported function
func PublicFunc() string {
	return "public"
}

// privateFunc is an unexported function
func privateFunc() string {
	return "private"
}

// PublicType is an exported type
type PublicType struct {
	// PublicField is exported
	PublicField string
	// privateField is unexported
	privateField int
}

// PublicMethod is an exported method
func (p *PublicType) PublicMethod() {}

// privateMethod is an unexported method
func (p *PublicType) privateMethod() {}

// NewPublicType is an exported constructor
func NewPublicType() *PublicType {
	return &PublicType{}
}

// newPrivateConstructor is an unexported constructor
func newPrivateConstructor() *PublicType {
	return &PublicType{}
}

// privateType is an unexported type
type privateType struct {
	field string
}

// PublicVar is an exported variable
var PublicVar = "public"

// privateVar is an unexported variable
var privateVar = "private"

// PublicConst is an exported constant
const PublicConst = "public"

// privateConst is an unexported constant
const privateConst = "private"

