#!/bin/bash -x


__dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
. ${__dir}/vars/variables.sh

echo ""
echo " ::: STARTING ETH PROVIDERS :::" $ETH_RPC_HTTP

# SINGLE PROXY
MOCK_PORT=2001
go run ./testutil/e2e/proxy/. $ETH_RPC_HTTP  -p $MOCK_PORT -cache=true -id eth -strict=false -no-save=false &

# Multi Port Proxy
# sh ./scripts/mock_proxy_eth.sh &
# sleep 2

# echo " ::: DONE ETHMOCK PROXY ::: "

echo " ::: RUNNING ETH PROVIDERS :::"
# SINGLE MOCK PROXY
lavad server 127.0.0.1 2221 http://0.0.0.0:$MOCK_PORT ETH1 jsonrpc --from servicer1 --geolocation 1 --log_level debug &
lavad server 127.0.0.1 2222 http://0.0.0.0:$MOCK_PORT ETH1 jsonrpc --from servicer2 --geolocation 1 --log_level debug &
lavad server 127.0.0.1 2223 http://0.0.0.0:$MOCK_PORT ETH1 jsonrpc --from servicer3 --geolocation 1 --log_level debug &
lavad server 127.0.0.1 2224 http://0.0.0.0:$MOCK_PORT ETH1 jsonrpc --from servicer4 --geolocation 1 --log_level debug &
lavad server 127.0.0.1 2225 http://0.0.0.0:$MOCK_PORT ETH1 jsonrpc --from servicer5 --geolocation 1 --log_level debug &
lavad portal_server 127.0.0.1 3333 ETH1 jsonrpc --from user1 --geolocation 1 --log_level debug

# Multi Port Proxy
# lavad server 127.0.0.1 2221 http://0.0.0.0:2001/$ETH_URL_PATH ETH1 jsonrpc --from servicer1 &
# lavad server 127.0.0.1 2222 http://0.0.0.0:2002/$ETH_URL_PATH ETH1 jsonrpc --from servicer2 &
# lavad server 127.0.0.1 2223 http://0.0.0.0:2003/$ETH_URL_PATH ETH1 jsonrpc --from servicer3 &
# lavad server 127.0.0.1 2224 http://0.0.0.0:2004/$ETH_URL_PATH ETH1 jsonrpc --from servicer4 &
# lavad server 127.0.0.1 2225 http://0.0.0.0:2005/$ETH_URL_PATH ETH1 jsonrpc --from servicer5 

# NO MOCK PROXY
# lavad server 127.0.0.1 2221 $ETH_RPC_WS ETH1 jsonrpc --from servicer1 &
# lavad server 127.0.0.1 2222 $ETH_RPC_WS ETH1 jsonrpc --from servicer2 &
# lavad server 127.0.0.1 2223 $ETH_RPC_WS ETH1 jsonrpc --from servicer3 &
# lavad server 127.0.0.1 2224 $ETH_RPC_WS ETH1 jsonrpc --from servicer4 &
# lavad server 127.0.0.1 2225 $ETH_RPC_WS ETH1 jsonrpc --from servicer5 

echo " ::: ETH PROVIDERS DONE! :::"