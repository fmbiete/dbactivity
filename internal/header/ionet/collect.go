package ionet

func (o *IONET) Collect() {
	o.IO.Collect()
	o.NET.Collect()
}
