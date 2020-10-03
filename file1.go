package error
typ Value interface { String() string }
func New(s string) Value

