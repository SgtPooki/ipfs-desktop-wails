#!/usr/bin/env bash

# see https://github.com/ipfs/kubo/blob/839b0848aec623ca68975efe7b628b06054074bf/core/corehttp/webui.go#L4
# or
# https://github.com/ipfs/ipfs-webui/releases
# WEBUI_CID=bafybeihfkeactw26tghz7m3puzh4zqlukvft2f7atfdc7t2qmqn2vszhc4 #v2.17.3
WEBUI_CID=bafybeibozpulxtpv5nhfa2ue3dcjx23ndh3gwr5vwllk7ptoyfwnfjjr4q #v2.15.1

main() {
  ipfs get "/ipfs/$WEBUI_CID" -o frontend
}

main
