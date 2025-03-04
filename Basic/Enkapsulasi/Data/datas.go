package data

type Data struct {
	nama   string
	alamat string
}

/* Ini constructor */
func NewData() Data {

	return Data{
		nama:   "Deni",
		alamat: "Jakarta",
	}
}

/*
Nama setter Getter nya harus huruf besar
Supaya bisa di akses public
*/
func (d *Data) SetName(nama string) {

	d.nama = nama
}

func (d *Data) SetAlamat(alamat string) {

	d.alamat = alamat
}

func (d *Data) GetName() string {

	return d.nama
}

func (d *Data) GetAlamat() string {
	return d.alamat
}
