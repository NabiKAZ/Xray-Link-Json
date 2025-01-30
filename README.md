# DESCRIPTION
This program takes input from the command line and checks whether the input is a JSON or a link.
If the input is a JSON, it converts it to Share links, and if the input is a link, it converts it to JSON.
This is done using functions from the libXray library, and the result is encoded in Base64.
 
# INSTALL
`go mod tidy`

# USAGE
Example:\
`go run main.go "vless://123456789@example.com:443?security=tls&sni=sni.example.com&type=ws&host=host.example.com&path=%2F#sample-vless"`

Output:
```
{"success":true,"data":{"transport":null,"log":null,"routing":null,"dns":null,"inbounds":null,"outbounds":[{"protocol":"vless","sendThrough":"sample-vless","tag":"","settings":{"vnext":[{"address":"example.com","port":443,"users":[{"id":"123456789","encryption":"none"}]}]},"streamSettings":{"address":null,"port":0,"network":"ws","security":"tls","tlsSettings":{"allowInsecure":false,"certificates":null,"serverName":"sni.example.com","alpn":null,"enableSessionResumption":false,"disableSystemRoot":false,"minVersion":"","maxVersion":"","cipherSuites":"","fingerprint":"","rejectUnknownSni":false,"pinnedPeerCertificateChainSha256":null,"pinnedPeerCertificatePublicKeySha256":null,"curvePreferences":null,"masterKeyLog":"","serverNameToVerify":""},"realitySettings":null,"rawSettings":null,"tcpSettings":null,"xhttpSettings":null,"splithttpSettings":null,"kcpSettings":null,"grpcSettings":null,"wsSettings":{"host":"host.example.com","path":"/","headers":null,"acceptProxyProtocol":false,"heartbeatPeriod":0},"httpupgradeSettings":null,"sockopt":null},"proxySettings":null,"mux":null}],"policy":null,"api":null,"metrics":null,"stats":null,"reverse":null,"fakeDns":null,"observatory":null,"burstObservatory":null}}
```

Example:\
`go run main.go "dm1lc3M6Ly9leGFtcGxlLmNvbTo0NDM="`

Output:
```
{"success":true,"data":{"transport":null,"log":null,"routing":null,"dns":null,"inbounds":null,"outbounds":[{"protocol":"vmess","sendThrough":"","tag":"","settings":{"vnext":[{"address":"example.com","port":443,"users":[{"id":"","security":"","experiments":""}]}]},"streamSettings":null,"proxySettings":null,"mux":null}],"policy":null,"api":null,"metrics":null,"stats":null,"reverse":null,"fakeDns":null,"observatory":null,"burstObservatory":null}}
```

Example:\
`go run main.go "{\"outbounds\":[{\"protocol\":\"vless\",\"settings\":{\"vnext\":[{\"address\":\"example.com\",\"port\":443}]}}]}"`

Output:
```
{"success":true,"data":"vless://example.com:443"}
```

# AUTHOR
Programmer: NabiKAZ ([x.com/NabiKAZ](x.com/NabiKAZ))

# DONATION
If this project was useful for you and you are willing, you can make me happy by giving a Star at the top of this GitHub page. \
Also this is my wallet address for Donate:

USDT (TRC20): `TEHjxGqu5Y2ExKBWzArBJEmrtzz3mgV5Hb` \
TON: `UQAzK0qhttfz1kte3auTXGqVeRul0SyFaCZORFyV1WmYlZQj`
