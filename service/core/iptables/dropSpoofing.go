package iptables

type dropSpoofing struct{ iptablesSetter }

var DropSpoofing dropSpoofing

func (r *dropSpoofing) GetSetupCommands() SetupCommands {
	commands := `
iptables -w 2 -N DROP_SPOOFING
iptables -w 2 -A DROP_SPOOFING -p udp --sport 53 -m string --algo bm --hex-string "|00047f|" --from 60 --to 180 -j DROP
iptables -w 2 -A DROP_SPOOFING -p udp --sport 53 -m string --algo bm --hex-string "|000400000000|" --from 60 --to 180 -j DROP
iptables -w 2 -I INPUT -j DROP_SPOOFING 
iptables -w 2 -I FORWARD -j DROP_SPOOFING
`
	return SetupCommands(commands)
}

func (r *dropSpoofing) GetCleanCommands() CleanCommands {
	commands := `
iptables -w 2 -D INPUT -j DROP_SPOOFING 
iptables -w 2 -D FORWARD -j DROP_SPOOFING 
iptables -w 2 -F DROP_SPOOFING
iptables -w 2 -X DROP_SPOOFING
`
	return CleanCommands(commands)
}
