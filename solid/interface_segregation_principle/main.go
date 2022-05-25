package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

func (m MultiFunctionPrinter) Fax(d Document) {
	//TODO implement me
	panic("implement me")
}

func (m MultiFunctionPrinter) Scan(d Document) {
	//TODO implement me
	panic("implement me")
}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

func (o OldFashionedPrinter) Fax(d Document) {
	//TODO implement me
	panic("operation not supported")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	//TODO implement me
	panic("operation not supported")
}

// better approach: split into several interfaces

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// printer only

type MyPrinter struct {
	// ...
}

func (m MyPrinter) Print(d Document) {
	// ...
}

// combine interfaces

type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	//
}

func (p Photocopier) Print(d Document) {
	//
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// interface combination + decorator

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
