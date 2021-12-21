module tcpgo

go 1.17

require (
	tcpgo.com/tcpserver v0.0.0

)

replace (
	tcpgo.com/tcpserver v0.0.0 => ./tcp
)
