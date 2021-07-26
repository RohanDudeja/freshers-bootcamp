package Models
type Student struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	LastName string `json:"last_name"`
	DOB uint `json:"DOB"`
	Address string `json:"address"`
	Subject   string `json:"subject"`
	Marks   string `json:"marks"`
}
func (b *Student) TableName() string {
	return "student"
}