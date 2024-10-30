module github.com/mattn/todo-webapp

go 1.22.0

toolchain go1.23.1

require (
	github.com/labstack/echo/v4 v4.12.0
	github.com/mattn/go-oci8 v0.1.1
	github.com/uptrace/bun v1.2.6-0.20241030103142-1bb5597f1a32
	github.com/uptrace/bun/dialect/oracledialect v0.0.0-20241030103142-1bb5597f1a32
	github.com/uptrace/bun/dialect/pgdialect v1.2.5
	github.com/uptrace/bun/extra/bundebug v1.2.5
)

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/puzpuzpuz/xsync/v3 v3.4.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)

//replace github.com/uptrace/bun => ../../uptrace/bun
