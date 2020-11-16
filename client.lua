
local socket = require("socket")
local ToServerPacket = require("packet/ToServerPacket")
local Constants = require("packet/Constants")
local PacketType = require("packet/PacketType")

local function tohex(str)
    return (str:gsub('.', function (c)
        return string.format('%02X', string.byte(c))
    end))
end

print("Sending data")
local udp = socket.udp()
udp:setpeername("remote.rudin.io", 30000)
udp:settimeout(100)

udp:send(
	Constants.protocol_id ..
	string.char(0x00, 0x00) .. -- peer_id
	string.char(0x00) .. -- channel
	PacketType.reliable ..
	string.char(0xff, 0xdc) .. -- seq nr
	string.char(0x01) .. -- subtype
	string.char(0x00, 0x00)
)

local data = udp:receive()
if data then
    print("Received: ", tohex(data))
end
