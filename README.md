# nfq
Minimal packet processign in userspace, using subgraph pure Go netlink library

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
