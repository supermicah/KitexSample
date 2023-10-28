module kitexsample

go 1.18

require (
    kitexserver v0.1.0
    kitexclient v0.1.0
)

replace (
	kitexserver => ./kitexserver
	kitexclient => ./kitexclient
)