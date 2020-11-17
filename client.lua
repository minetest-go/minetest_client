
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
local udp = socket.udp()
udp:setpeername(server, tonumber(port))
udp:settimeout(100)

local function tx(def)
  print("TX: " .. DumpPacket(def))
  local data = ClientServerPacket.create(def)
  print(">> Sending:  " .. tohex(data) .. " len: " .. #data)
  udp:send(data)
end

tx({
  peer_id = 0,
  channel = 0,
  type = "reliable",
  sequence_nr = 65500,
  subtype = "original",
  payload = string.char(0x00, 0x00)
})

local peer_id = 1

while true do
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

      tx({
        peer_id = peer_id,
        channel = 1,
        type = "original",
        command = "INIT",
        payload = {
          client_max = 28,
          supp_compr_modes = 0,
          min_net_proto_version = 37,
          max_net_proto_version = 39,
          player_name = "blah"
        }
      })

      -- TODO: send "INIT" after a delay
    end

    if packet.command == "HELLO" then
      tx({
        peer_id = peer_id,
        channel = 1,
        type = "original",
        command = "FIRST_SRP",
        payload = {
          salt = "", -- 16 bytes
          verification_key = "", -- 256 bytes
          is_empty = 0
        }
      })

    end

  end
end
