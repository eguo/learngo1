package error
type Value interface { String() string }
func New(s string) Value

