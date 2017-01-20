# nfq
Minimal packet processign in userspace, using subgraph pure Go netlink library

The `master` branch includes a version of the subgraph package modified to
include additional debug output as well as correct message alignment.

The `subgraph` branch includes the original subgraph/go-nfnetlink version.

## Set up IPTables
`# sudo iptables -A OUTPUT -p tcp -j NFQUEUE --queue-num 1 --queue-bypass`

## Run program
`# sudo ./nfq`

In another terminal, try initiating some outbound connection (e.g., `curl www.google.com`)
Processed packets would be printed.

## Delete the rule
`# sudo iptables -L OUTPUT --line-numbers`

followed by

`# sudo iptables -D OUTPUT <line number>`
